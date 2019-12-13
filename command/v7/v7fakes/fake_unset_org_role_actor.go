// Code generated by counterfeiter. DO NOT EDIT.
package v7fakes

import (
	"sync"

	"code.cloudfoundry.org/cli/actor/v7action"
	"code.cloudfoundry.org/cli/api/cloudcontroller/ccv3/constant"
	v7 "code.cloudfoundry.org/cli/command/v7"
)

type FakeUnsetOrgRoleActor struct {
	DeleteOrgRoleStub        func(constant.RoleType, string, string, string, bool) (v7action.Warnings, error)
	deleteOrgRoleMutex       sync.RWMutex
	deleteOrgRoleArgsForCall []struct {
		arg1 constant.RoleType
		arg2 string
		arg3 string
		arg4 string
		arg5 bool
	}
	deleteOrgRoleReturns struct {
		result1 v7action.Warnings
		result2 error
	}
	deleteOrgRoleReturnsOnCall map[int]struct {
		result1 v7action.Warnings
		result2 error
	}
	GetOrganizationByNameStub        func(string) (v7action.Organization, v7action.Warnings, error)
	getOrganizationByNameMutex       sync.RWMutex
	getOrganizationByNameArgsForCall []struct {
		arg1 string
	}
	getOrganizationByNameReturns struct {
		result1 v7action.Organization
		result2 v7action.Warnings
		result3 error
	}
	getOrganizationByNameReturnsOnCall map[int]struct {
		result1 v7action.Organization
		result2 v7action.Warnings
		result3 error
	}
	GetUserStub        func(string, string) (v7action.User, error)
	getUserMutex       sync.RWMutex
	getUserArgsForCall []struct {
		arg1 string
		arg2 string
	}
	getUserReturns struct {
		result1 v7action.User
		result2 error
	}
	getUserReturnsOnCall map[int]struct {
		result1 v7action.User
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeUnsetOrgRoleActor) DeleteOrgRole(arg1 constant.RoleType, arg2 string, arg3 string, arg4 string, arg5 bool) (v7action.Warnings, error) {
	fake.deleteOrgRoleMutex.Lock()
	ret, specificReturn := fake.deleteOrgRoleReturnsOnCall[len(fake.deleteOrgRoleArgsForCall)]
	fake.deleteOrgRoleArgsForCall = append(fake.deleteOrgRoleArgsForCall, struct {
		arg1 constant.RoleType
		arg2 string
		arg3 string
		arg4 string
		arg5 bool
	}{arg1, arg2, arg3, arg4, arg5})
	fake.recordInvocation("DeleteOrgRole", []interface{}{arg1, arg2, arg3, arg4, arg5})
	fake.deleteOrgRoleMutex.Unlock()
	if fake.DeleteOrgRoleStub != nil {
		return fake.DeleteOrgRoleStub(arg1, arg2, arg3, arg4, arg5)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.deleteOrgRoleReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeUnsetOrgRoleActor) DeleteOrgRoleCallCount() int {
	fake.deleteOrgRoleMutex.RLock()
	defer fake.deleteOrgRoleMutex.RUnlock()
	return len(fake.deleteOrgRoleArgsForCall)
}

func (fake *FakeUnsetOrgRoleActor) DeleteOrgRoleCalls(stub func(constant.RoleType, string, string, string, bool) (v7action.Warnings, error)) {
	fake.deleteOrgRoleMutex.Lock()
	defer fake.deleteOrgRoleMutex.Unlock()
	fake.DeleteOrgRoleStub = stub
}

func (fake *FakeUnsetOrgRoleActor) DeleteOrgRoleArgsForCall(i int) (constant.RoleType, string, string, string, bool) {
	fake.deleteOrgRoleMutex.RLock()
	defer fake.deleteOrgRoleMutex.RUnlock()
	argsForCall := fake.deleteOrgRoleArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4, argsForCall.arg5
}

func (fake *FakeUnsetOrgRoleActor) DeleteOrgRoleReturns(result1 v7action.Warnings, result2 error) {
	fake.deleteOrgRoleMutex.Lock()
	defer fake.deleteOrgRoleMutex.Unlock()
	fake.DeleteOrgRoleStub = nil
	fake.deleteOrgRoleReturns = struct {
		result1 v7action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeUnsetOrgRoleActor) DeleteOrgRoleReturnsOnCall(i int, result1 v7action.Warnings, result2 error) {
	fake.deleteOrgRoleMutex.Lock()
	defer fake.deleteOrgRoleMutex.Unlock()
	fake.DeleteOrgRoleStub = nil
	if fake.deleteOrgRoleReturnsOnCall == nil {
		fake.deleteOrgRoleReturnsOnCall = make(map[int]struct {
			result1 v7action.Warnings
			result2 error
		})
	}
	fake.deleteOrgRoleReturnsOnCall[i] = struct {
		result1 v7action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeUnsetOrgRoleActor) GetOrganizationByName(arg1 string) (v7action.Organization, v7action.Warnings, error) {
	fake.getOrganizationByNameMutex.Lock()
	ret, specificReturn := fake.getOrganizationByNameReturnsOnCall[len(fake.getOrganizationByNameArgsForCall)]
	fake.getOrganizationByNameArgsForCall = append(fake.getOrganizationByNameArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("GetOrganizationByName", []interface{}{arg1})
	fake.getOrganizationByNameMutex.Unlock()
	if fake.GetOrganizationByNameStub != nil {
		return fake.GetOrganizationByNameStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.getOrganizationByNameReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeUnsetOrgRoleActor) GetOrganizationByNameCallCount() int {
	fake.getOrganizationByNameMutex.RLock()
	defer fake.getOrganizationByNameMutex.RUnlock()
	return len(fake.getOrganizationByNameArgsForCall)
}

func (fake *FakeUnsetOrgRoleActor) GetOrganizationByNameCalls(stub func(string) (v7action.Organization, v7action.Warnings, error)) {
	fake.getOrganizationByNameMutex.Lock()
	defer fake.getOrganizationByNameMutex.Unlock()
	fake.GetOrganizationByNameStub = stub
}

func (fake *FakeUnsetOrgRoleActor) GetOrganizationByNameArgsForCall(i int) string {
	fake.getOrganizationByNameMutex.RLock()
	defer fake.getOrganizationByNameMutex.RUnlock()
	argsForCall := fake.getOrganizationByNameArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeUnsetOrgRoleActor) GetOrganizationByNameReturns(result1 v7action.Organization, result2 v7action.Warnings, result3 error) {
	fake.getOrganizationByNameMutex.Lock()
	defer fake.getOrganizationByNameMutex.Unlock()
	fake.GetOrganizationByNameStub = nil
	fake.getOrganizationByNameReturns = struct {
		result1 v7action.Organization
		result2 v7action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeUnsetOrgRoleActor) GetOrganizationByNameReturnsOnCall(i int, result1 v7action.Organization, result2 v7action.Warnings, result3 error) {
	fake.getOrganizationByNameMutex.Lock()
	defer fake.getOrganizationByNameMutex.Unlock()
	fake.GetOrganizationByNameStub = nil
	if fake.getOrganizationByNameReturnsOnCall == nil {
		fake.getOrganizationByNameReturnsOnCall = make(map[int]struct {
			result1 v7action.Organization
			result2 v7action.Warnings
			result3 error
		})
	}
	fake.getOrganizationByNameReturnsOnCall[i] = struct {
		result1 v7action.Organization
		result2 v7action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeUnsetOrgRoleActor) GetUser(arg1 string, arg2 string) (v7action.User, error) {
	fake.getUserMutex.Lock()
	ret, specificReturn := fake.getUserReturnsOnCall[len(fake.getUserArgsForCall)]
	fake.getUserArgsForCall = append(fake.getUserArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("GetUser", []interface{}{arg1, arg2})
	fake.getUserMutex.Unlock()
	if fake.GetUserStub != nil {
		return fake.GetUserStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getUserReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeUnsetOrgRoleActor) GetUserCallCount() int {
	fake.getUserMutex.RLock()
	defer fake.getUserMutex.RUnlock()
	return len(fake.getUserArgsForCall)
}

func (fake *FakeUnsetOrgRoleActor) GetUserCalls(stub func(string, string) (v7action.User, error)) {
	fake.getUserMutex.Lock()
	defer fake.getUserMutex.Unlock()
	fake.GetUserStub = stub
}

func (fake *FakeUnsetOrgRoleActor) GetUserArgsForCall(i int) (string, string) {
	fake.getUserMutex.RLock()
	defer fake.getUserMutex.RUnlock()
	argsForCall := fake.getUserArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeUnsetOrgRoleActor) GetUserReturns(result1 v7action.User, result2 error) {
	fake.getUserMutex.Lock()
	defer fake.getUserMutex.Unlock()
	fake.GetUserStub = nil
	fake.getUserReturns = struct {
		result1 v7action.User
		result2 error
	}{result1, result2}
}

func (fake *FakeUnsetOrgRoleActor) GetUserReturnsOnCall(i int, result1 v7action.User, result2 error) {
	fake.getUserMutex.Lock()
	defer fake.getUserMutex.Unlock()
	fake.GetUserStub = nil
	if fake.getUserReturnsOnCall == nil {
		fake.getUserReturnsOnCall = make(map[int]struct {
			result1 v7action.User
			result2 error
		})
	}
	fake.getUserReturnsOnCall[i] = struct {
		result1 v7action.User
		result2 error
	}{result1, result2}
}

func (fake *FakeUnsetOrgRoleActor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.deleteOrgRoleMutex.RLock()
	defer fake.deleteOrgRoleMutex.RUnlock()
	fake.getOrganizationByNameMutex.RLock()
	defer fake.getOrganizationByNameMutex.RUnlock()
	fake.getUserMutex.RLock()
	defer fake.getUserMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeUnsetOrgRoleActor) recordInvocation(key string, args []interface{}) {
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

var _ v7.UnsetOrgRoleActor = new(FakeUnsetOrgRoleActor)