/*
Copyright 2018 the Velero contributors.

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

package velero

import (
	"context"
	"sort"
	"testing"
	"time"

	. "github.com/onsi/gomega"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/util/clock"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"

	velerov1api "github.com/vmware-tanzu/velero/pkg/apis/velero/v1"
	"github.com/vmware-tanzu/velero/pkg/builder"
	"github.com/vmware-tanzu/velero/pkg/buildinfo"
	"github.com/vmware-tanzu/velero/pkg/generated/clientset/versioned/scheme"
	"github.com/vmware-tanzu/velero/pkg/plugin/framework"
)

func statusRequestBuilder(resourceVersion string) *builder.ServerStatusRequestBuilder {
	return builder.ForServerStatusRequest(velerov1api.DefaultNamespace, "sr-1", resourceVersion)
}

func TestProcess(t *testing.T) {
	// now will be used to set the fake clock's time; capture
	// it here so it can be referenced in the test case defs.
	now, err := time.Parse(time.RFC1123, time.RFC1123)
	require.NoError(t, err)
	now = now.Local()

	buildinfo.Version = "test-version-val"

	tests := []struct {
		name            string
		req             *velerov1api.ServerStatusRequest
		reqPluginLister *fakePluginLister
		expected        *velerov1api.ServerStatusRequest
		expectedErrMsg  string
	}{
		{
			name: "server status request with empty phase gets processed",
			req:  statusRequestBuilder("0").Result(),
			reqPluginLister: &fakePluginLister{
				plugins: []framework.PluginIdentifier{
					{
						Name: "custom.io/myown",
						Kind: "VolumeSnapshotter",
					},
				},
			},
			expected: statusRequestBuilder("1").
				ServerVersion(buildinfo.Version).
				Phase(velerov1api.ServerStatusRequestPhaseProcessed).
				ProcessedTimestamp(now).
				Plugins([]velerov1api.PluginInfo{
					{
						Name: "custom.io/myown",
						Kind: "VolumeSnapshotter",
					},
				}).
				Result(),
		},
		{
			name: "server status request with phase=New gets processed",
			req: statusRequestBuilder("0").
				Phase(velerov1api.ServerStatusRequestPhaseNew).
				Result(),
			reqPluginLister: &fakePluginLister{
				plugins: []framework.PluginIdentifier{
					{
						Name: "velero.io/aws",
						Kind: "ObjectStore",
					},
					{
						Name: "custom.io/myown",
						Kind: "VolumeSnapshotter",
					},
				},
			},
			expected: statusRequestBuilder("1").
				ServerVersion(buildinfo.Version).
				Phase(velerov1api.ServerStatusRequestPhaseProcessed).
				ProcessedTimestamp(now).
				Plugins([]velerov1api.PluginInfo{
					{
						Name: "velero.io/aws",
						Kind: "ObjectStore",
					},
					{
						Name: "custom.io/myown",
						Kind: "VolumeSnapshotter",
					},
				}).
				Result(),
		},
		{
			name: "server status request with phase=Processed gets deleted if expired",
			req: statusRequestBuilder("0").
				Phase(velerov1api.ServerStatusRequestPhaseProcessed).
				ProcessedTimestamp(now.Add(-61 * time.Second)).
				Result(),
			reqPluginLister: &fakePluginLister{
				plugins: []framework.PluginIdentifier{
					{
						Name: "custom.io/myown",
						Kind: "VolumeSnapshotter",
					},
				},
			},
			expected: nil,
		},
		{
			name: "server status request with phase=Processed does not get deleted if not expired",
			req: statusRequestBuilder("0").
				Phase(velerov1api.ServerStatusRequestPhaseProcessed).
				ProcessedTimestamp(now.Add(-59 * time.Second)).
				Result(),
			expected: statusRequestBuilder("0").
				Phase(velerov1api.ServerStatusRequestPhaseProcessed).
				ProcessedTimestamp(now.Add(-59 * time.Second)).
				Result(),
		},
		{
			name: "server status request with invalid phase returns an error",
			req: statusRequestBuilder("0").
				Phase(velerov1api.ServerStatusRequestPhase("an-invalid-phase")).
				Result(),
			expected: statusRequestBuilder("0").
				Phase(velerov1api.ServerStatusRequestPhase("an-invalid-phase")).
				Result(),
			expectedErrMsg: "unexpected ServerStatusRequest phase \"an-invalid-phase\"",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			g := NewWithT(t)

			kbClient := fake.NewFakeClientWithScheme(scheme.Scheme, tc.req)

			err := Process(tc.req, kbClient, tc.reqPluginLister, clock.NewFakeClock(now), logrus.StandardLogger())
			if tc.expectedErrMsg == "" {
				assert.Nil(t, err)
			} else {
				assert.EqualError(t, err, tc.expectedErrMsg)
			}

			key := client.ObjectKey{Name: tc.req.Name, Namespace: tc.req.Namespace}
			instance := &velerov1api.ServerStatusRequest{}
			err = kbClient.Get(context.Background(), key, instance)

			if tc.expected == nil {
				g.Expect(&velerov1api.ServerStatusRequest{}).To(BeEquivalentTo((&velerov1api.ServerStatusRequest{})))
				g.Expect(apierrors.IsNotFound(err)).To(BeTrue())
			} else {
				sortPluginsByKindAndName(tc.expected.Status.Plugins)
				sortPluginsByKindAndName(instance.Status.Plugins)
				g.Expect(tc.expected.Status.Plugins).To(BeEquivalentTo((instance.Status.Plugins)))
				g.Expect(tc.expected).To(BeEquivalentTo((instance)))
				g.Expect(err).To(BeNil())
			}
		})
	}
}

func sortPluginsByKindAndName(plugins []velerov1api.PluginInfo) {
	sort.Slice(plugins, func(i, j int) bool {
		if plugins[i].Kind != plugins[j].Kind {
			return plugins[i].Kind < plugins[j].Kind
		}
		return plugins[i].Name < plugins[j].Name
	})
}

type fakePluginLister struct {
	plugins []framework.PluginIdentifier
}

func (l *fakePluginLister) List(kind framework.PluginKind) []framework.PluginIdentifier {
	var plugins []framework.PluginIdentifier
	for _, plugin := range l.plugins {
		if plugin.Kind == kind {
			plugins = append(plugins, plugin)
		}
	}

	return plugins
}
