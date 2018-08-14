// Code generated by counterfeiter. DO NOT EDIT.
package counterfeiterexamplefakes

import (
	"sync"

	counterfeiter_example "github.com/huangping40/fabric_learning/counterfeiter-example"
)

type FakePeerLedgerProvider struct {
	CreateStub        func(genesisBlock *counterfeiter_example.Block) (counterfeiter_example.PeerLedger, error)
	createMutex       sync.RWMutex
	createArgsForCall []struct {
		genesisBlock *counterfeiter_example.Block
	}
	createReturns struct {
		result1 counterfeiter_example.PeerLedger
		result2 error
	}
	createReturnsOnCall map[int]struct {
		result1 counterfeiter_example.PeerLedger
		result2 error
	}
	OpenStub        func(ledgerID string) (counterfeiter_example.PeerLedger, error)
	openMutex       sync.RWMutex
	openArgsForCall []struct {
		ledgerID string
	}
	openReturns struct {
		result1 counterfeiter_example.PeerLedger
		result2 error
	}
	openReturnsOnCall map[int]struct {
		result1 counterfeiter_example.PeerLedger
		result2 error
	}
	ExistsStub        func(ledgerID string) (bool, error)
	existsMutex       sync.RWMutex
	existsArgsForCall []struct {
		ledgerID string
	}
	existsReturns struct {
		result1 bool
		result2 error
	}
	existsReturnsOnCall map[int]struct {
		result1 bool
		result2 error
	}
	ListStub        func() ([]string, error)
	listMutex       sync.RWMutex
	listArgsForCall []struct{}
	listReturns     struct {
		result1 []string
		result2 error
	}
	listReturnsOnCall map[int]struct {
		result1 []string
		result2 error
	}
	CloseStub        func()
	closeMutex       sync.RWMutex
	closeArgsForCall []struct{}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakePeerLedgerProvider) Create(genesisBlock *counterfeiter_example.Block) (counterfeiter_example.PeerLedger, error) {
	fake.createMutex.Lock()
	ret, specificReturn := fake.createReturnsOnCall[len(fake.createArgsForCall)]
	fake.createArgsForCall = append(fake.createArgsForCall, struct {
		genesisBlock *counterfeiter_example.Block
	}{genesisBlock})
	fake.recordInvocation("Create", []interface{}{genesisBlock})
	fake.createMutex.Unlock()
	if fake.CreateStub != nil {
		return fake.CreateStub(genesisBlock)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.createReturns.result1, fake.createReturns.result2
}

func (fake *FakePeerLedgerProvider) CreateCallCount() int {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return len(fake.createArgsForCall)
}

func (fake *FakePeerLedgerProvider) CreateArgsForCall(i int) *counterfeiter_example.Block {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return fake.createArgsForCall[i].genesisBlock
}

func (fake *FakePeerLedgerProvider) CreateReturns(result1 counterfeiter_example.PeerLedger, result2 error) {
	fake.CreateStub = nil
	fake.createReturns = struct {
		result1 counterfeiter_example.PeerLedger
		result2 error
	}{result1, result2}
}

func (fake *FakePeerLedgerProvider) CreateReturnsOnCall(i int, result1 counterfeiter_example.PeerLedger, result2 error) {
	fake.CreateStub = nil
	if fake.createReturnsOnCall == nil {
		fake.createReturnsOnCall = make(map[int]struct {
			result1 counterfeiter_example.PeerLedger
			result2 error
		})
	}
	fake.createReturnsOnCall[i] = struct {
		result1 counterfeiter_example.PeerLedger
		result2 error
	}{result1, result2}
}

func (fake *FakePeerLedgerProvider) Open(ledgerID string) (counterfeiter_example.PeerLedger, error) {
	fake.openMutex.Lock()
	ret, specificReturn := fake.openReturnsOnCall[len(fake.openArgsForCall)]
	fake.openArgsForCall = append(fake.openArgsForCall, struct {
		ledgerID string
	}{ledgerID})
	fake.recordInvocation("Open", []interface{}{ledgerID})
	fake.openMutex.Unlock()
	if fake.OpenStub != nil {
		return fake.OpenStub(ledgerID)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.openReturns.result1, fake.openReturns.result2
}

func (fake *FakePeerLedgerProvider) OpenCallCount() int {
	fake.openMutex.RLock()
	defer fake.openMutex.RUnlock()
	return len(fake.openArgsForCall)
}

func (fake *FakePeerLedgerProvider) OpenArgsForCall(i int) string {
	fake.openMutex.RLock()
	defer fake.openMutex.RUnlock()
	return fake.openArgsForCall[i].ledgerID
}

func (fake *FakePeerLedgerProvider) OpenReturns(result1 counterfeiter_example.PeerLedger, result2 error) {
	fake.OpenStub = nil
	fake.openReturns = struct {
		result1 counterfeiter_example.PeerLedger
		result2 error
	}{result1, result2}
}

func (fake *FakePeerLedgerProvider) OpenReturnsOnCall(i int, result1 counterfeiter_example.PeerLedger, result2 error) {
	fake.OpenStub = nil
	if fake.openReturnsOnCall == nil {
		fake.openReturnsOnCall = make(map[int]struct {
			result1 counterfeiter_example.PeerLedger
			result2 error
		})
	}
	fake.openReturnsOnCall[i] = struct {
		result1 counterfeiter_example.PeerLedger
		result2 error
	}{result1, result2}
}

func (fake *FakePeerLedgerProvider) Exists(ledgerID string) (bool, error) {
	fake.existsMutex.Lock()
	ret, specificReturn := fake.existsReturnsOnCall[len(fake.existsArgsForCall)]
	fake.existsArgsForCall = append(fake.existsArgsForCall, struct {
		ledgerID string
	}{ledgerID})
	fake.recordInvocation("Exists", []interface{}{ledgerID})
	fake.existsMutex.Unlock()
	if fake.ExistsStub != nil {
		return fake.ExistsStub(ledgerID)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.existsReturns.result1, fake.existsReturns.result2
}

func (fake *FakePeerLedgerProvider) ExistsCallCount() int {
	fake.existsMutex.RLock()
	defer fake.existsMutex.RUnlock()
	return len(fake.existsArgsForCall)
}

func (fake *FakePeerLedgerProvider) ExistsArgsForCall(i int) string {
	fake.existsMutex.RLock()
	defer fake.existsMutex.RUnlock()
	return fake.existsArgsForCall[i].ledgerID
}

func (fake *FakePeerLedgerProvider) ExistsReturns(result1 bool, result2 error) {
	fake.ExistsStub = nil
	fake.existsReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakePeerLedgerProvider) ExistsReturnsOnCall(i int, result1 bool, result2 error) {
	fake.ExistsStub = nil
	if fake.existsReturnsOnCall == nil {
		fake.existsReturnsOnCall = make(map[int]struct {
			result1 bool
			result2 error
		})
	}
	fake.existsReturnsOnCall[i] = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakePeerLedgerProvider) List() ([]string, error) {
	fake.listMutex.Lock()
	ret, specificReturn := fake.listReturnsOnCall[len(fake.listArgsForCall)]
	fake.listArgsForCall = append(fake.listArgsForCall, struct{}{})
	fake.recordInvocation("List", []interface{}{})
	fake.listMutex.Unlock()
	if fake.ListStub != nil {
		return fake.ListStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.listReturns.result1, fake.listReturns.result2
}

func (fake *FakePeerLedgerProvider) ListCallCount() int {
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	return len(fake.listArgsForCall)
}

func (fake *FakePeerLedgerProvider) ListReturns(result1 []string, result2 error) {
	fake.ListStub = nil
	fake.listReturns = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakePeerLedgerProvider) ListReturnsOnCall(i int, result1 []string, result2 error) {
	fake.ListStub = nil
	if fake.listReturnsOnCall == nil {
		fake.listReturnsOnCall = make(map[int]struct {
			result1 []string
			result2 error
		})
	}
	fake.listReturnsOnCall[i] = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakePeerLedgerProvider) Close() {
	fake.closeMutex.Lock()
	fake.closeArgsForCall = append(fake.closeArgsForCall, struct{}{})
	fake.recordInvocation("Close", []interface{}{})
	fake.closeMutex.Unlock()
	if fake.CloseStub != nil {
		fake.CloseStub()
	}
}

func (fake *FakePeerLedgerProvider) CloseCallCount() int {
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	return len(fake.closeArgsForCall)
}

func (fake *FakePeerLedgerProvider) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	fake.openMutex.RLock()
	defer fake.openMutex.RUnlock()
	fake.existsMutex.RLock()
	defer fake.existsMutex.RUnlock()
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakePeerLedgerProvider) recordInvocation(key string, args []interface{}) {
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

var _ counterfeiter_example.PeerLedgerProvider = new(FakePeerLedgerProvider)