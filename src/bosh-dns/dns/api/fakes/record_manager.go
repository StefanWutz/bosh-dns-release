// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"bosh-dns/dns/api"
	"bosh-dns/dns/server/record"
	"sync"
)

type FakeRecordManager struct {
	AllRecordsStub        func() *[]record.Record
	allRecordsMutex       sync.RWMutex
	allRecordsArgsForCall []struct {
	}
	allRecordsReturns struct {
		result1 *[]record.Record
	}
	allRecordsReturnsOnCall map[int]struct {
		result1 *[]record.Record
	}
	ExpandAliasesStub        func(string) []string
	expandAliasesMutex       sync.RWMutex
	expandAliasesArgsForCall []struct {
		arg1 string
	}
	expandAliasesReturns struct {
		result1 []string
	}
	expandAliasesReturnsOnCall map[int]struct {
		result1 []string
	}
	ResolveRecordsStub        func([]string, bool) ([]record.Record, error)
	resolveRecordsMutex       sync.RWMutex
	resolveRecordsArgsForCall []struct {
		arg1 []string
		arg2 bool
	}
	resolveRecordsReturns struct {
		result1 []record.Record
		result2 error
	}
	resolveRecordsReturnsOnCall map[int]struct {
		result1 []record.Record
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeRecordManager) AllRecords() *[]record.Record {
	fake.allRecordsMutex.Lock()
	ret, specificReturn := fake.allRecordsReturnsOnCall[len(fake.allRecordsArgsForCall)]
	fake.allRecordsArgsForCall = append(fake.allRecordsArgsForCall, struct {
	}{})
	fake.recordInvocation("AllRecords", []interface{}{})
	fake.allRecordsMutex.Unlock()
	if fake.AllRecordsStub != nil {
		return fake.AllRecordsStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.allRecordsReturns
	return fakeReturns.result1
}

func (fake *FakeRecordManager) AllRecordsCallCount() int {
	fake.allRecordsMutex.RLock()
	defer fake.allRecordsMutex.RUnlock()
	return len(fake.allRecordsArgsForCall)
}

func (fake *FakeRecordManager) AllRecordsCalls(stub func() *[]record.Record) {
	fake.allRecordsMutex.Lock()
	defer fake.allRecordsMutex.Unlock()
	fake.AllRecordsStub = stub
}

func (fake *FakeRecordManager) AllRecordsReturns(result1 *[]record.Record) {
	fake.allRecordsMutex.Lock()
	defer fake.allRecordsMutex.Unlock()
	fake.AllRecordsStub = nil
	fake.allRecordsReturns = struct {
		result1 *[]record.Record
	}{result1}
}

func (fake *FakeRecordManager) AllRecordsReturnsOnCall(i int, result1 *[]record.Record) {
	fake.allRecordsMutex.Lock()
	defer fake.allRecordsMutex.Unlock()
	fake.AllRecordsStub = nil
	if fake.allRecordsReturnsOnCall == nil {
		fake.allRecordsReturnsOnCall = make(map[int]struct {
			result1 *[]record.Record
		})
	}
	fake.allRecordsReturnsOnCall[i] = struct {
		result1 *[]record.Record
	}{result1}
}

func (fake *FakeRecordManager) ExpandAliases(arg1 string) []string {
	fake.expandAliasesMutex.Lock()
	ret, specificReturn := fake.expandAliasesReturnsOnCall[len(fake.expandAliasesArgsForCall)]
	fake.expandAliasesArgsForCall = append(fake.expandAliasesArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("ExpandAliases", []interface{}{arg1})
	fake.expandAliasesMutex.Unlock()
	if fake.ExpandAliasesStub != nil {
		return fake.ExpandAliasesStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.expandAliasesReturns
	return fakeReturns.result1
}

func (fake *FakeRecordManager) ExpandAliasesCallCount() int {
	fake.expandAliasesMutex.RLock()
	defer fake.expandAliasesMutex.RUnlock()
	return len(fake.expandAliasesArgsForCall)
}

func (fake *FakeRecordManager) ExpandAliasesCalls(stub func(string) []string) {
	fake.expandAliasesMutex.Lock()
	defer fake.expandAliasesMutex.Unlock()
	fake.ExpandAliasesStub = stub
}

func (fake *FakeRecordManager) ExpandAliasesArgsForCall(i int) string {
	fake.expandAliasesMutex.RLock()
	defer fake.expandAliasesMutex.RUnlock()
	argsForCall := fake.expandAliasesArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeRecordManager) ExpandAliasesReturns(result1 []string) {
	fake.expandAliasesMutex.Lock()
	defer fake.expandAliasesMutex.Unlock()
	fake.ExpandAliasesStub = nil
	fake.expandAliasesReturns = struct {
		result1 []string
	}{result1}
}

func (fake *FakeRecordManager) ExpandAliasesReturnsOnCall(i int, result1 []string) {
	fake.expandAliasesMutex.Lock()
	defer fake.expandAliasesMutex.Unlock()
	fake.ExpandAliasesStub = nil
	if fake.expandAliasesReturnsOnCall == nil {
		fake.expandAliasesReturnsOnCall = make(map[int]struct {
			result1 []string
		})
	}
	fake.expandAliasesReturnsOnCall[i] = struct {
		result1 []string
	}{result1}
}

func (fake *FakeRecordManager) ResolveRecords(arg1 []string, arg2 bool) ([]record.Record, error) {
	var arg1Copy []string
	if arg1 != nil {
		arg1Copy = make([]string, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.resolveRecordsMutex.Lock()
	ret, specificReturn := fake.resolveRecordsReturnsOnCall[len(fake.resolveRecordsArgsForCall)]
	fake.resolveRecordsArgsForCall = append(fake.resolveRecordsArgsForCall, struct {
		arg1 []string
		arg2 bool
	}{arg1Copy, arg2})
	fake.recordInvocation("ResolveRecords", []interface{}{arg1Copy, arg2})
	fake.resolveRecordsMutex.Unlock()
	if fake.ResolveRecordsStub != nil {
		return fake.ResolveRecordsStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.resolveRecordsReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeRecordManager) ResolveRecordsCallCount() int {
	fake.resolveRecordsMutex.RLock()
	defer fake.resolveRecordsMutex.RUnlock()
	return len(fake.resolveRecordsArgsForCall)
}

func (fake *FakeRecordManager) ResolveRecordsCalls(stub func([]string, bool) ([]record.Record, error)) {
	fake.resolveRecordsMutex.Lock()
	defer fake.resolveRecordsMutex.Unlock()
	fake.ResolveRecordsStub = stub
}

func (fake *FakeRecordManager) ResolveRecordsArgsForCall(i int) ([]string, bool) {
	fake.resolveRecordsMutex.RLock()
	defer fake.resolveRecordsMutex.RUnlock()
	argsForCall := fake.resolveRecordsArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeRecordManager) ResolveRecordsReturns(result1 []record.Record, result2 error) {
	fake.resolveRecordsMutex.Lock()
	defer fake.resolveRecordsMutex.Unlock()
	fake.ResolveRecordsStub = nil
	fake.resolveRecordsReturns = struct {
		result1 []record.Record
		result2 error
	}{result1, result2}
}

func (fake *FakeRecordManager) ResolveRecordsReturnsOnCall(i int, result1 []record.Record, result2 error) {
	fake.resolveRecordsMutex.Lock()
	defer fake.resolveRecordsMutex.Unlock()
	fake.ResolveRecordsStub = nil
	if fake.resolveRecordsReturnsOnCall == nil {
		fake.resolveRecordsReturnsOnCall = make(map[int]struct {
			result1 []record.Record
			result2 error
		})
	}
	fake.resolveRecordsReturnsOnCall[i] = struct {
		result1 []record.Record
		result2 error
	}{result1, result2}
}

func (fake *FakeRecordManager) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.allRecordsMutex.RLock()
	defer fake.allRecordsMutex.RUnlock()
	fake.expandAliasesMutex.RLock()
	defer fake.expandAliasesMutex.RUnlock()
	fake.resolveRecordsMutex.RLock()
	defer fake.resolveRecordsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeRecordManager) recordInvocation(key string, args []interface{}) {
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

var _ api.RecordManager = new(FakeRecordManager)
