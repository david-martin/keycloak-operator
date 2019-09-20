package test

import (
	"github.com/go-logr/logr"
	"github.com/keycloak/keycloak-operator/pkg/common"
	"k8s.io/apimachinery/pkg/runtime"
)

type MockActionRunner struct {
	logger           logr.Logger
	ResourcesCreated int
	ResourcesUpdated int
}

func NewMockActionRunner() common.ActionRunner {
	return &MockActionRunner{
		ResourcesCreated: 0,
		ResourcesUpdated: 0,
	}
}

func (i *MockActionRunner) Create(obj runtime.Object) error {
	i.ResourcesCreated = i.ResourcesCreated + 1
	return nil
}

func (i *MockActionRunner) Update(obj runtime.Object) error {
	i.ResourcesUpdated = i.ResourcesUpdated + 1
	return nil
}

func (i *MockActionRunner) RunAll(desiredState common.DesiredClusterState) error {
	for _, action := range desiredState {
		_, err := action.Run(i)
		if err != nil {
			return err
		}
	}

	return nil
}