/*
Copyright the Velero contributors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package e2e

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"time"

	"github.com/pkg/errors"

	veleroexec "github.com/vmware-tanzu/velero/pkg/util/exec"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	kibishiiNamespace = "kibishii-workload"
	jumpPadPod        = "jump-pad"
)

func installKibishiiWorkload(client testClient, cloudPlatform string) error {
	fmt.Println("Creating the kibishii namespace and workload")
	fiveMinTimeout, cancel := context.WithTimeout(client.ctx, 5*time.Minute)
	defer cancel()

	if err := createNamespace(fiveMinTimeout, client, kibishiiNamespace); err != nil {
		return errors.Wrapf(err, "Failed to create the namespace kibishii for the Kibishii workload")
	}

	// We use kustomize to generate YAML for Kibishii from the checked-in yaml directories
	kibishiiInstallCmd := exec.CommandContext(client.ctx, "kubectl", "apply", "-n", kibishiiNamespace, "-k",
		"github.com/vmware-tanzu-experiments/distributed-data-generator/kubernetes/yaml/"+cloudPlatform)

	_, _, err := veleroexec.RunCommand(kibishiiInstallCmd)
	if err != nil {
		return errors.Wrap(err, "failed to install kibishii")
	}

	kibishiiSetWaitCmd := exec.CommandContext(client.ctx, "kubectl", "rollout", "status", "statefulset.apps/kibishii-deployment",
		"-n", kibishiiNamespace, "-w", "--timeout=30m")
	kibishiiSetWaitCmd.Stdout = os.Stdout
	kibishiiSetWaitCmd.Stderr = os.Stderr
	_, _, err = veleroexec.RunCommand(kibishiiSetWaitCmd)

	if err != nil {
		return err
	}

	fmt.Printf("Waiting for kibishii jump-pad pod to be ready\n")
	jumpPadWaitCmd := exec.CommandContext(client.ctx, "kubectl", "wait", "--for=condition=ready", "-n", kibishiiNamespace, "pod/jump-pad")
	jumpPadWaitCmd.Stdout = os.Stdout
	jumpPadWaitCmd.Stderr = os.Stderr
	_, _, err = veleroexec.RunCommand(jumpPadWaitCmd)
	if err != nil {
		return errors.Wrapf(err, "Failed to wait for ready status of pod %s/%s", kibishiiNamespace, jumpPadPod)
	}

	return err
}

func terminateKibishiiWorkload(client testClient) error {
	fiveMinTimeout, cancel := context.WithTimeout(client.ctx, 5*time.Minute)
	defer cancel()

	oneHourTimeout, stop := context.WithTimeout(client.ctx, time.Minute*15) // TODO(carlisai) set this back to 60 min
	defer stop()

	interval := 5 * time.Second
	timeout := 10 * time.Minute

	// delete the ns
	if err := client.clientGo.CoreV1().Namespaces().Delete(oneHourTimeout, kibishiiNamespace, metav1.DeleteOptions{}); err != nil {
		return errors.Wrap(err, "Failed to delete the kibishii namespace")
	}

	// wait for ns delete
	err := waitForNamespaceDeletion(fiveMinTimeout, interval, timeout, client, kibishiiNamespace)
	if err != nil {
		return errors.Wrap(err, "Failed to wait for the kibishii namespace to terminate")
	}

	return nil
}

// runKibishiiTests:
// - inserts data into the kibishii workload
// - creates a Velero backup of the kibishii workload in the given namespace
// - deletes the kibishii workload by deleting its namespace to simulate a disaster
// - creates a Velero restore of the created backup in the given names-ace
// - verifies the data restored is what's expected
// Assumes the kibishii workload has been created and Velero has been installed.
func runKibishiiTests(client testClient, testNamespace testNamespace, providerName, veleroCLI, backupName,
	restoreName, backupLocation string, useVolumeSnapshots bool) error {
	fmt.Println("Starting the backup and restore of the kibishii workload")
	fiveMinTimeout, cancel := context.WithTimeout(client.ctx, 5*time.Minute)
	defer cancel()

	oneHourTimeout, stop := context.WithTimeout(client.ctx, time.Minute*60)
	defer stop()

	// wait for kibishii pod startup
	// TODO - Fix kibishii so we can check that it is ready to go
	fmt.Printf("Waiting for kibishii pods to be ready\n")
	if err := waitForKibishiiPods(oneHourTimeout, client); err != nil {
		return errors.Wrapf(err, "failed to wait for ready status of kibishii pods in %s", kibishiiNamespace)
	}

	// if err := generateData(oneHourTimeout, 2, 10, 10, 1024, 1024, 0, 2); err != nil {
	if err := generateData(oneHourTimeout, 1, 1, 1, 1024, 1024, 0, 2); err != nil {
		return errors.Wrap(err, "failed to generate data")
	}

	if err := veleroBackupNamespace(oneHourTimeout, testNamespace, veleroCLI, backupName, kibishiiNamespace, backupLocation, useVolumeSnapshots); err != nil {
		veleroBackupLocationStatus(fiveMinTimeout, testNamespace, veleroCLI, backupLocation)
		veleroBackupLogs(fiveMinTimeout, testNamespace, veleroCLI, backupName)

		err = fmt.Errorf("failed to backup the kibishii namespace with error %s", errors.WithStack(err))
		return err
	}

	if providerName == "vsphere" && useVolumeSnapshots {
		// Wait for uploads started by the Velero Plug-in for vSphere to complete
		// TODO - remove after upload progress monitoring is implemented
		fmt.Println("Waiting for vSphere uploads to complete")
		if err := waitForVSphereUploadCompletion(oneHourTimeout, time.Hour, kibishiiNamespace); err != nil {
			return errors.Wrapf(err, "error waiting for uploads to complete")
		}
	}

	fmt.Println("Simulating a disaster by removing the kibishii namespace")
	if err := terminateKibishiiWorkload(client); err != nil {
		return errors.Wrap(err, "failed to simulate a disaster")
	}

	if err := veleroRestoreNamespace(oneHourTimeout, testNamespace, veleroCLI, restoreName, backupName); err != nil {
		veleroBackupLocationStatus(fiveMinTimeout, testNamespace, veleroCLI, backupLocation)
		veleroRestoreLogs(fiveMinTimeout, testNamespace, veleroCLI, restoreName)

		err = fmt.Errorf("restore %s failed from backup %s with error %s", restoreName, backupName, errors.WithStack(err))
		return err
	}

	// wait for kibishii pod startup
	// TODO - Fix kibishii so we can check that it is ready to go
	fmt.Printf("Waiting for kibishii pods to be ready\n")
	if err := waitForKibishiiPods(oneHourTimeout, client); err != nil {
		return errors.Wrap(err, "failed to wait for ready status of kibishii pods in the kibishii namespace")
	}

	// TODO - check that namespace exists
	fmt.Printf("running kibishii verify\n")
	if err := verifyData(oneHourTimeout, 2, 10, 10, 1024, 1024, 0, 2); err != nil {
		return errors.Wrap(err, "failed to verify data generated by kibishii")
	}

	fmt.Println("Backup and restore of the kibishii workload completed successfully")
	return nil
}

func waitForKibishiiPods(ctx context.Context, client testClient) error {
	return waitForPods(ctx, client, kibishiiNamespace, []string{"jump-pad", "etcd0", "etcd1", "etcd2", "kibishii-deployment-0", "kibishii-deployment-1"})
}

func generateData(ctx context.Context, levels int, filesPerLevel int, dirsPerLevel int, fileSize int,
	blockSize int, passNum int, expectedNodes int) error {
	kibishiiGenerateCmd := exec.CommandContext(ctx, "kubectl", "exec", "-n", kibishiiNamespace, "jump-pad", "--",
		"/usr/local/bin/generate.sh", strconv.Itoa(levels), strconv.Itoa(filesPerLevel), strconv.Itoa(dirsPerLevel), strconv.Itoa(fileSize),
		strconv.Itoa(blockSize), strconv.Itoa(passNum), strconv.Itoa(expectedNodes))
	fmt.Printf("kibishiiGenerateCmd cmd =%v\n", kibishiiGenerateCmd)

	_, _, err := veleroexec.RunCommand(kibishiiGenerateCmd)
	if err != nil {
		return errors.Wrapf(err, "failed to generate data")
	}

	return nil
}

func verifyData(ctx context.Context, levels int, filesPerLevel int, dirsPerLevel int, fileSize int,
	blockSize int, passNum int, expectedNodes int) error {
	kibishiiVerifyCmd := exec.CommandContext(ctx, "kubectl", "exec", "-n", kibishiiNamespace, "jump-pad", "--",
		"/usr/local/bin/verify.sh", strconv.Itoa(levels), strconv.Itoa(filesPerLevel), strconv.Itoa(dirsPerLevel), strconv.Itoa(fileSize),
		strconv.Itoa(blockSize), strconv.Itoa(passNum), strconv.Itoa(expectedNodes))
	fmt.Printf("kibishiiVerifyCmd cmd =%v\n", kibishiiVerifyCmd)

	_, _, err := veleroexec.RunCommand(kibishiiVerifyCmd)
	if err != nil {
		return errors.Wrapf(err, "failed to verify data")
	}

	return nil
}
