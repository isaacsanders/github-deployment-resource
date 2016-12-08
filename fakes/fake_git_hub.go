// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/ahume/github-deployment-resource"
	"github.com/ahume/go-github/github"
)

type FakeGitHub struct {
	ListDeploymentsStub        func() ([]*github.Deployment, error)
	listDeploymentsMutex       sync.RWMutex
	listDeploymentsArgsForCall []struct{}
	listDeploymentsReturns     struct {
		result1 []*github.Deployment
		result2 error
	}
	ListDeploymentStatusesStub        func(ID int) ([]*github.DeploymentStatus, error)
	listDeploymentStatusesMutex       sync.RWMutex
	listDeploymentStatusesArgsForCall []struct {
		ID int
	}
	listDeploymentStatusesReturns struct {
		result1 []*github.DeploymentStatus
		result2 error
	}
	GetDeploymentStub        func(ID int) (*github.Deployment, error)
	getDeploymentMutex       sync.RWMutex
	getDeploymentArgsForCall []struct {
		ID int
	}
	getDeploymentReturns struct {
		result1 *github.Deployment
		result2 error
	}
	CreateDeploymentStub        func(request *github.DeploymentRequest) (*github.Deployment, error)
	createDeploymentMutex       sync.RWMutex
	createDeploymentArgsForCall []struct {
		request *github.DeploymentRequest
	}
	createDeploymentReturns struct {
		result1 *github.Deployment
		result2 error
	}
	CreateDeploymentStatusStub        func(ID int, request *github.DeploymentStatusRequest) (*github.DeploymentStatus, error)
	createDeploymentStatusMutex       sync.RWMutex
	createDeploymentStatusArgsForCall []struct {
		ID      int
		request *github.DeploymentStatusRequest
	}
	createDeploymentStatusReturns struct {
		result1 *github.DeploymentStatus
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeGitHub) ListDeployments() ([]*github.Deployment, error) {
	fake.listDeploymentsMutex.Lock()
	fake.listDeploymentsArgsForCall = append(fake.listDeploymentsArgsForCall, struct{}{})
	fake.recordInvocation("ListDeployments", []interface{}{})
	fake.listDeploymentsMutex.Unlock()
	if fake.ListDeploymentsStub != nil {
		return fake.ListDeploymentsStub()
	} else {
		return fake.listDeploymentsReturns.result1, fake.listDeploymentsReturns.result2
	}
}

func (fake *FakeGitHub) ListDeploymentsCallCount() int {
	fake.listDeploymentsMutex.RLock()
	defer fake.listDeploymentsMutex.RUnlock()
	return len(fake.listDeploymentsArgsForCall)
}

func (fake *FakeGitHub) ListDeploymentsReturns(result1 []*github.Deployment, result2 error) {
	fake.ListDeploymentsStub = nil
	fake.listDeploymentsReturns = struct {
		result1 []*github.Deployment
		result2 error
	}{result1, result2}
}

func (fake *FakeGitHub) ListDeploymentStatuses(ID int) ([]*github.DeploymentStatus, error) {
	fake.listDeploymentStatusesMutex.Lock()
	fake.listDeploymentStatusesArgsForCall = append(fake.listDeploymentStatusesArgsForCall, struct {
		ID int
	}{ID})
	fake.recordInvocation("ListDeploymentStatuses", []interface{}{ID})
	fake.listDeploymentStatusesMutex.Unlock()
	if fake.ListDeploymentStatusesStub != nil {
		return fake.ListDeploymentStatusesStub(ID)
	} else {
		return fake.listDeploymentStatusesReturns.result1, fake.listDeploymentStatusesReturns.result2
	}
}

func (fake *FakeGitHub) ListDeploymentStatusesCallCount() int {
	fake.listDeploymentStatusesMutex.RLock()
	defer fake.listDeploymentStatusesMutex.RUnlock()
	return len(fake.listDeploymentStatusesArgsForCall)
}

func (fake *FakeGitHub) ListDeploymentStatusesArgsForCall(i int) int {
	fake.listDeploymentStatusesMutex.RLock()
	defer fake.listDeploymentStatusesMutex.RUnlock()
	return fake.listDeploymentStatusesArgsForCall[i].ID
}

func (fake *FakeGitHub) ListDeploymentStatusesReturns(result1 []*github.DeploymentStatus, result2 error) {
	fake.ListDeploymentStatusesStub = nil
	fake.listDeploymentStatusesReturns = struct {
		result1 []*github.DeploymentStatus
		result2 error
	}{result1, result2}
}

func (fake *FakeGitHub) GetDeployment(ID int) (*github.Deployment, error) {
	fake.getDeploymentMutex.Lock()
	fake.getDeploymentArgsForCall = append(fake.getDeploymentArgsForCall, struct {
		ID int
	}{ID})
	fake.recordInvocation("GetDeployment", []interface{}{ID})
	fake.getDeploymentMutex.Unlock()
	if fake.GetDeploymentStub != nil {
		return fake.GetDeploymentStub(ID)
	} else {
		return fake.getDeploymentReturns.result1, fake.getDeploymentReturns.result2
	}
}

func (fake *FakeGitHub) GetDeploymentCallCount() int {
	fake.getDeploymentMutex.RLock()
	defer fake.getDeploymentMutex.RUnlock()
	return len(fake.getDeploymentArgsForCall)
}

func (fake *FakeGitHub) GetDeploymentArgsForCall(i int) int {
	fake.getDeploymentMutex.RLock()
	defer fake.getDeploymentMutex.RUnlock()
	return fake.getDeploymentArgsForCall[i].ID
}

func (fake *FakeGitHub) GetDeploymentReturns(result1 *github.Deployment, result2 error) {
	fake.GetDeploymentStub = nil
	fake.getDeploymentReturns = struct {
		result1 *github.Deployment
		result2 error
	}{result1, result2}
}

func (fake *FakeGitHub) CreateDeployment(request *github.DeploymentRequest) (*github.Deployment, error) {
	fake.createDeploymentMutex.Lock()
	fake.createDeploymentArgsForCall = append(fake.createDeploymentArgsForCall, struct {
		request *github.DeploymentRequest
	}{request})
	fake.recordInvocation("CreateDeployment", []interface{}{request})
	fake.createDeploymentMutex.Unlock()
	if fake.CreateDeploymentStub != nil {
		return fake.CreateDeploymentStub(request)
	} else {
		return fake.createDeploymentReturns.result1, fake.createDeploymentReturns.result2
	}
}

func (fake *FakeGitHub) CreateDeploymentCallCount() int {
	fake.createDeploymentMutex.RLock()
	defer fake.createDeploymentMutex.RUnlock()
	return len(fake.createDeploymentArgsForCall)
}

func (fake *FakeGitHub) CreateDeploymentArgsForCall(i int) *github.DeploymentRequest {
	fake.createDeploymentMutex.RLock()
	defer fake.createDeploymentMutex.RUnlock()
	return fake.createDeploymentArgsForCall[i].request
}

func (fake *FakeGitHub) CreateDeploymentReturns(result1 *github.Deployment, result2 error) {
	fake.CreateDeploymentStub = nil
	fake.createDeploymentReturns = struct {
		result1 *github.Deployment
		result2 error
	}{result1, result2}
}

func (fake *FakeGitHub) CreateDeploymentStatus(ID int, request *github.DeploymentStatusRequest) (*github.DeploymentStatus, error) {
	fake.createDeploymentStatusMutex.Lock()
	fake.createDeploymentStatusArgsForCall = append(fake.createDeploymentStatusArgsForCall, struct {
		ID      int
		request *github.DeploymentStatusRequest
	}{ID, request})
	fake.recordInvocation("CreateDeploymentStatus", []interface{}{ID, request})
	fake.createDeploymentStatusMutex.Unlock()
	if fake.CreateDeploymentStatusStub != nil {
		return fake.CreateDeploymentStatusStub(ID, request)
	} else {
		return fake.createDeploymentStatusReturns.result1, fake.createDeploymentStatusReturns.result2
	}
}

func (fake *FakeGitHub) CreateDeploymentStatusCallCount() int {
	fake.createDeploymentStatusMutex.RLock()
	defer fake.createDeploymentStatusMutex.RUnlock()
	return len(fake.createDeploymentStatusArgsForCall)
}

func (fake *FakeGitHub) CreateDeploymentStatusArgsForCall(i int) (int, *github.DeploymentStatusRequest) {
	fake.createDeploymentStatusMutex.RLock()
	defer fake.createDeploymentStatusMutex.RUnlock()
	return fake.createDeploymentStatusArgsForCall[i].ID, fake.createDeploymentStatusArgsForCall[i].request
}

func (fake *FakeGitHub) CreateDeploymentStatusReturns(result1 *github.DeploymentStatus, result2 error) {
	fake.CreateDeploymentStatusStub = nil
	fake.createDeploymentStatusReturns = struct {
		result1 *github.DeploymentStatus
		result2 error
	}{result1, result2}
}

func (fake *FakeGitHub) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.listDeploymentsMutex.RLock()
	defer fake.listDeploymentsMutex.RUnlock()
	fake.listDeploymentStatusesMutex.RLock()
	defer fake.listDeploymentStatusesMutex.RUnlock()
	fake.getDeploymentMutex.RLock()
	defer fake.getDeploymentMutex.RUnlock()
	fake.createDeploymentMutex.RLock()
	defer fake.createDeploymentMutex.RUnlock()
	fake.createDeploymentStatusMutex.RLock()
	defer fake.createDeploymentStatusMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeGitHub) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ resource.GitHub = new(FakeGitHub)
