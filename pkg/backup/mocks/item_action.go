/*
Copyright 2018 the Heptio Ark contributors.

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
// Code generated by mockery v1.0.0. DO NOT EDIT.
package mocks

import (
	mock "github.com/stretchr/testify/mock"
	runtime "k8s.io/apimachinery/pkg/runtime"

	v1 "github.com/heptio/velero/pkg/apis/velero/v1"
	"github.com/heptio/velero/pkg/plugin/interface/actioninterface"
)

// BackupItemAction is an autogenerated mock type for the BackupItemAction type
type BackupItemAction struct {
	mock.Mock
}

// AppliesTo provides a mock function with given fields:
func (_m *BackupItemAction) AppliesTo() (actioninterface.ResourceSelector, error) {
	ret := _m.Called()

	var r0 actioninterface.ResourceSelector
	if rf, ok := ret.Get(0).(func() actioninterface.ResourceSelector); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(actioninterface.ResourceSelector)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Execute provides a mock function with given fields: item, _a1
func (_m *BackupItemAction) Execute(item runtime.Unstructured, _a1 *v1.Backup) (runtime.Unstructured, []actioninterface.ResourceIdentifier, error) {
	ret := _m.Called(item, _a1)

	var r0 runtime.Unstructured
	if rf, ok := ret.Get(0).(func(runtime.Unstructured, *v1.Backup) runtime.Unstructured); ok {
		r0 = rf(item, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(runtime.Unstructured)
		}
	}

	var r1 []actioninterface.ResourceIdentifier
	if rf, ok := ret.Get(1).(func(runtime.Unstructured, *v1.Backup) []actioninterface.ResourceIdentifier); ok {
		r1 = rf(item, _a1)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]actioninterface.ResourceIdentifier)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(runtime.Unstructured, *v1.Backup) error); ok {
		r2 = rf(item, _a1)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}
