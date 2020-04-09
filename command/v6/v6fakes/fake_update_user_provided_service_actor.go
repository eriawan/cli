// Code generated by counterfeiter. DO NOT EDIT.
package v6fakes

import (
	"sync"

	"code.cloudfoundry.org/cli/actor/v2action"
	v6 "code.cloudfoundry.org/cli/command/v6"
)

type FakeUpdateUserProvidedServiceActor struct {
	GetServiceInstanceByNameAndSpaceStub        func(string, string) (v2action.ServiceInstance, v2action.Warnings, error)
	getServiceInstanceByNameAndSpaceMutex       sync.RWMutex
	getServiceInstanceByNameAndSpaceArgsForCall []struct {
		arg1 string
		arg2 string
	}
	getServiceInstanceByNameAndSpaceReturns struct {
		result1 v2action.ServiceInstance
		result2 v2action.Warnings
		result3 error
	}
	getServiceInstanceByNameAndSpaceReturnsOnCall map[int]struct {
		result1 v2action.ServiceInstance
		result2 v2action.Warnings
		result3 error
	}
	UpdateUserProvidedServiceInstanceStub        func(string, v2action.UserProvidedServiceInstance) (v2action.Warnings, error)
	updateUserProvidedServiceInstanceMutex       sync.RWMutex
	updateUserProvidedServiceInstanceArgsForCall []struct {
		arg1 string
		arg2 v2action.UserProvidedServiceInstance
	}
	updateUserProvidedServiceInstanceReturns struct {
		result1 v2action.Warnings
		result2 error
	}
	updateUserProvidedServiceInstanceReturnsOnCall map[int]struct {
		result1 v2action.Warnings
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeUpdateUserProvidedServiceActor) GetServiceInstanceByNameAndSpace(arg1 string, arg2 string) (v2action.ServiceInstance, v2action.Warnings, error) {
	fake.getServiceInstanceByNameAndSpaceMutex.Lock()
	ret, specificReturn := fake.getServiceInstanceByNameAndSpaceReturnsOnCall[len(fake.getServiceInstanceByNameAndSpaceArgsForCall)]
	fake.getServiceInstanceByNameAndSpaceArgsForCall = append(fake.getServiceInstanceByNameAndSpaceArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("GetServiceInstanceByNameAndSpace", []interface{}{arg1, arg2})
	fake.getServiceInstanceByNameAndSpaceMutex.Unlock()
	if fake.GetServiceInstanceByNameAndSpaceStub != nil {
		return fake.GetServiceInstanceByNameAndSpaceStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.getServiceInstanceByNameAndSpaceReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeUpdateUserProvidedServiceActor) GetServiceInstanceByNameAndSpaceCallCount() int {
	fake.getServiceInstanceByNameAndSpaceMutex.RLock()
	defer fake.getServiceInstanceByNameAndSpaceMutex.RUnlock()
	return len(fake.getServiceInstanceByNameAndSpaceArgsForCall)
}

func (fake *FakeUpdateUserProvidedServiceActor) GetServiceInstanceByNameAndSpaceCalls(stub func(string, string) (v2action.ServiceInstance, v2action.Warnings, error)) {
	fake.getServiceInstanceByNameAndSpaceMutex.Lock()
	defer fake.getServiceInstanceByNameAndSpaceMutex.Unlock()
	fake.GetServiceInstanceByNameAndSpaceStub = stub
}

func (fake *FakeUpdateUserProvidedServiceActor) GetServiceInstanceByNameAndSpaceArgsForCall(i int) (string, string) {
	fake.getServiceInstanceByNameAndSpaceMutex.RLock()
	defer fake.getServiceInstanceByNameAndSpaceMutex.RUnlock()
	argsForCall := fake.getServiceInstanceByNameAndSpaceArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeUpdateUserProvidedServiceActor) GetServiceInstanceByNameAndSpaceReturns(result1 v2action.ServiceInstance, result2 v2action.Warnings, result3 error) {
	fake.getServiceInstanceByNameAndSpaceMutex.Lock()
	defer fake.getServiceInstanceByNameAndSpaceMutex.Unlock()
	fake.GetServiceInstanceByNameAndSpaceStub = nil
	fake.getServiceInstanceByNameAndSpaceReturns = struct {
		result1 v2action.ServiceInstance
		result2 v2action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeUpdateUserProvidedServiceActor) GetServiceInstanceByNameAndSpaceReturnsOnCall(i int, result1 v2action.ServiceInstance, result2 v2action.Warnings, result3 error) {
	fake.getServiceInstanceByNameAndSpaceMutex.Lock()
	defer fake.getServiceInstanceByNameAndSpaceMutex.Unlock()
	fake.GetServiceInstanceByNameAndSpaceStub = nil
	if fake.getServiceInstanceByNameAndSpaceReturnsOnCall == nil {
		fake.getServiceInstanceByNameAndSpaceReturnsOnCall = make(map[int]struct {
			result1 v2action.ServiceInstance
			result2 v2action.Warnings
			result3 error
		})
	}
	fake.getServiceInstanceByNameAndSpaceReturnsOnCall[i] = struct {
		result1 v2action.ServiceInstance
		result2 v2action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeUpdateUserProvidedServiceActor) UpdateUserProvidedServiceInstance(arg1 string, arg2 v2action.UserProvidedServiceInstance) (v2action.Warnings, error) {
	fake.updateUserProvidedServiceInstanceMutex.Lock()
	ret, specificReturn := fake.updateUserProvidedServiceInstanceReturnsOnCall[len(fake.updateUserProvidedServiceInstanceArgsForCall)]
	fake.updateUserProvidedServiceInstanceArgsForCall = append(fake.updateUserProvidedServiceInstanceArgsForCall, struct {
		arg1 string
		arg2 v2action.UserProvidedServiceInstance
	}{arg1, arg2})
	fake.recordInvocation("UpdateUserProvidedServiceInstance", []interface{}{arg1, arg2})
	fake.updateUserProvidedServiceInstanceMutex.Unlock()
	if fake.UpdateUserProvidedServiceInstanceStub != nil {
		return fake.UpdateUserProvidedServiceInstanceStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.updateUserProvidedServiceInstanceReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeUpdateUserProvidedServiceActor) UpdateUserProvidedServiceInstanceCallCount() int {
	fake.updateUserProvidedServiceInstanceMutex.RLock()
	defer fake.updateUserProvidedServiceInstanceMutex.RUnlock()
	return len(fake.updateUserProvidedServiceInstanceArgsForCall)
}

func (fake *FakeUpdateUserProvidedServiceActor) UpdateUserProvidedServiceInstanceCalls(stub func(string, v2action.UserProvidedServiceInstance) (v2action.Warnings, error)) {
	fake.updateUserProvidedServiceInstanceMutex.Lock()
	defer fake.updateUserProvidedServiceInstanceMutex.Unlock()
	fake.UpdateUserProvidedServiceInstanceStub = stub
}

func (fake *FakeUpdateUserProvidedServiceActor) UpdateUserProvidedServiceInstanceArgsForCall(i int) (string, v2action.UserProvidedServiceInstance) {
	fake.updateUserProvidedServiceInstanceMutex.RLock()
	defer fake.updateUserProvidedServiceInstanceMutex.RUnlock()
	argsForCall := fake.updateUserProvidedServiceInstanceArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeUpdateUserProvidedServiceActor) UpdateUserProvidedServiceInstanceReturns(result1 v2action.Warnings, result2 error) {
	fake.updateUserProvidedServiceInstanceMutex.Lock()
	defer fake.updateUserProvidedServiceInstanceMutex.Unlock()
	fake.UpdateUserProvidedServiceInstanceStub = nil
	fake.updateUserProvidedServiceInstanceReturns = struct {
		result1 v2action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeUpdateUserProvidedServiceActor) UpdateUserProvidedServiceInstanceReturnsOnCall(i int, result1 v2action.Warnings, result2 error) {
	fake.updateUserProvidedServiceInstanceMutex.Lock()
	defer fake.updateUserProvidedServiceInstanceMutex.Unlock()
	fake.UpdateUserProvidedServiceInstanceStub = nil
	if fake.updateUserProvidedServiceInstanceReturnsOnCall == nil {
		fake.updateUserProvidedServiceInstanceReturnsOnCall = make(map[int]struct {
			result1 v2action.Warnings
			result2 error
		})
	}
	fake.updateUserProvidedServiceInstanceReturnsOnCall[i] = struct {
		result1 v2action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeUpdateUserProvidedServiceActor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getServiceInstanceByNameAndSpaceMutex.RLock()
	defer fake.getServiceInstanceByNameAndSpaceMutex.RUnlock()
	fake.updateUserProvidedServiceInstanceMutex.RLock()
	defer fake.updateUserProvidedServiceInstanceMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeUpdateUserProvidedServiceActor) recordInvocation(key string, args []interface{}) {
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

var _ v6.UpdateUserProvidedServiceActor = new(FakeUpdateUserProvidedServiceActor)