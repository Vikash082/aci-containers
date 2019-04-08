// Code generated by counterfeiter. DO NOT EDIT.
package dbfakes

import (
	"sync"

	"code.cloudfoundry.org/lager"
	"code.cloudfoundry.org/locket/db"
	"code.cloudfoundry.org/locket/models"
)

type FakeLockDB struct {
	LockStub        func(logger lager.Logger, resource *models.Resource, ttl int64) (*db.Lock, error)
	lockMutex       sync.RWMutex
	lockArgsForCall []struct {
		logger   lager.Logger
		resource *models.Resource
		ttl      int64
	}
	lockReturns struct {
		result1 *db.Lock
		result2 error
	}
	lockReturnsOnCall map[int]struct {
		result1 *db.Lock
		result2 error
	}
	ReleaseStub        func(logger lager.Logger, resource *models.Resource) error
	releaseMutex       sync.RWMutex
	releaseArgsForCall []struct {
		logger   lager.Logger
		resource *models.Resource
	}
	releaseReturns struct {
		result1 error
	}
	releaseReturnsOnCall map[int]struct {
		result1 error
	}
	FetchStub        func(logger lager.Logger, key string) (*db.Lock, error)
	fetchMutex       sync.RWMutex
	fetchArgsForCall []struct {
		logger lager.Logger
		key    string
	}
	fetchReturns struct {
		result1 *db.Lock
		result2 error
	}
	fetchReturnsOnCall map[int]struct {
		result1 *db.Lock
		result2 error
	}
	FetchAndReleaseStub        func(logger lager.Logger, lock *db.Lock) (bool, error)
	fetchAndReleaseMutex       sync.RWMutex
	fetchAndReleaseArgsForCall []struct {
		logger lager.Logger
		lock   *db.Lock
	}
	fetchAndReleaseReturns struct {
		result1 bool
		result2 error
	}
	fetchAndReleaseReturnsOnCall map[int]struct {
		result1 bool
		result2 error
	}
	FetchAllStub        func(logger lager.Logger, lockType string) ([]*db.Lock, error)
	fetchAllMutex       sync.RWMutex
	fetchAllArgsForCall []struct {
		logger   lager.Logger
		lockType string
	}
	fetchAllReturns struct {
		result1 []*db.Lock
		result2 error
	}
	fetchAllReturnsOnCall map[int]struct {
		result1 []*db.Lock
		result2 error
	}
	CountStub        func(logger lager.Logger, lockType string) (int, error)
	countMutex       sync.RWMutex
	countArgsForCall []struct {
		logger   lager.Logger
		lockType string
	}
	countReturns struct {
		result1 int
		result2 error
	}
	countReturnsOnCall map[int]struct {
		result1 int
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeLockDB) Lock(logger lager.Logger, resource *models.Resource, ttl int64) (*db.Lock, error) {
	fake.lockMutex.Lock()
	ret, specificReturn := fake.lockReturnsOnCall[len(fake.lockArgsForCall)]
	fake.lockArgsForCall = append(fake.lockArgsForCall, struct {
		logger   lager.Logger
		resource *models.Resource
		ttl      int64
	}{logger, resource, ttl})
	fake.recordInvocation("Lock", []interface{}{logger, resource, ttl})
	fake.lockMutex.Unlock()
	if fake.LockStub != nil {
		return fake.LockStub(logger, resource, ttl)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.lockReturns.result1, fake.lockReturns.result2
}

func (fake *FakeLockDB) LockCallCount() int {
	fake.lockMutex.RLock()
	defer fake.lockMutex.RUnlock()
	return len(fake.lockArgsForCall)
}

func (fake *FakeLockDB) LockArgsForCall(i int) (lager.Logger, *models.Resource, int64) {
	fake.lockMutex.RLock()
	defer fake.lockMutex.RUnlock()
	return fake.lockArgsForCall[i].logger, fake.lockArgsForCall[i].resource, fake.lockArgsForCall[i].ttl
}

func (fake *FakeLockDB) LockReturns(result1 *db.Lock, result2 error) {
	fake.LockStub = nil
	fake.lockReturns = struct {
		result1 *db.Lock
		result2 error
	}{result1, result2}
}

func (fake *FakeLockDB) LockReturnsOnCall(i int, result1 *db.Lock, result2 error) {
	fake.LockStub = nil
	if fake.lockReturnsOnCall == nil {
		fake.lockReturnsOnCall = make(map[int]struct {
			result1 *db.Lock
			result2 error
		})
	}
	fake.lockReturnsOnCall[i] = struct {
		result1 *db.Lock
		result2 error
	}{result1, result2}
}

func (fake *FakeLockDB) Release(logger lager.Logger, resource *models.Resource) error {
	fake.releaseMutex.Lock()
	ret, specificReturn := fake.releaseReturnsOnCall[len(fake.releaseArgsForCall)]
	fake.releaseArgsForCall = append(fake.releaseArgsForCall, struct {
		logger   lager.Logger
		resource *models.Resource
	}{logger, resource})
	fake.recordInvocation("Release", []interface{}{logger, resource})
	fake.releaseMutex.Unlock()
	if fake.ReleaseStub != nil {
		return fake.ReleaseStub(logger, resource)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.releaseReturns.result1
}

func (fake *FakeLockDB) ReleaseCallCount() int {
	fake.releaseMutex.RLock()
	defer fake.releaseMutex.RUnlock()
	return len(fake.releaseArgsForCall)
}

func (fake *FakeLockDB) ReleaseArgsForCall(i int) (lager.Logger, *models.Resource) {
	fake.releaseMutex.RLock()
	defer fake.releaseMutex.RUnlock()
	return fake.releaseArgsForCall[i].logger, fake.releaseArgsForCall[i].resource
}

func (fake *FakeLockDB) ReleaseReturns(result1 error) {
	fake.ReleaseStub = nil
	fake.releaseReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeLockDB) ReleaseReturnsOnCall(i int, result1 error) {
	fake.ReleaseStub = nil
	if fake.releaseReturnsOnCall == nil {
		fake.releaseReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.releaseReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeLockDB) Fetch(logger lager.Logger, key string) (*db.Lock, error) {
	fake.fetchMutex.Lock()
	ret, specificReturn := fake.fetchReturnsOnCall[len(fake.fetchArgsForCall)]
	fake.fetchArgsForCall = append(fake.fetchArgsForCall, struct {
		logger lager.Logger
		key    string
	}{logger, key})
	fake.recordInvocation("Fetch", []interface{}{logger, key})
	fake.fetchMutex.Unlock()
	if fake.FetchStub != nil {
		return fake.FetchStub(logger, key)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.fetchReturns.result1, fake.fetchReturns.result2
}

func (fake *FakeLockDB) FetchCallCount() int {
	fake.fetchMutex.RLock()
	defer fake.fetchMutex.RUnlock()
	return len(fake.fetchArgsForCall)
}

func (fake *FakeLockDB) FetchArgsForCall(i int) (lager.Logger, string) {
	fake.fetchMutex.RLock()
	defer fake.fetchMutex.RUnlock()
	return fake.fetchArgsForCall[i].logger, fake.fetchArgsForCall[i].key
}

func (fake *FakeLockDB) FetchReturns(result1 *db.Lock, result2 error) {
	fake.FetchStub = nil
	fake.fetchReturns = struct {
		result1 *db.Lock
		result2 error
	}{result1, result2}
}

func (fake *FakeLockDB) FetchReturnsOnCall(i int, result1 *db.Lock, result2 error) {
	fake.FetchStub = nil
	if fake.fetchReturnsOnCall == nil {
		fake.fetchReturnsOnCall = make(map[int]struct {
			result1 *db.Lock
			result2 error
		})
	}
	fake.fetchReturnsOnCall[i] = struct {
		result1 *db.Lock
		result2 error
	}{result1, result2}
}

func (fake *FakeLockDB) FetchAndRelease(logger lager.Logger, lock *db.Lock) (bool, error) {
	fake.fetchAndReleaseMutex.Lock()
	ret, specificReturn := fake.fetchAndReleaseReturnsOnCall[len(fake.fetchAndReleaseArgsForCall)]
	fake.fetchAndReleaseArgsForCall = append(fake.fetchAndReleaseArgsForCall, struct {
		logger lager.Logger
		lock   *db.Lock
	}{logger, lock})
	fake.recordInvocation("FetchAndRelease", []interface{}{logger, lock})
	fake.fetchAndReleaseMutex.Unlock()
	if fake.FetchAndReleaseStub != nil {
		return fake.FetchAndReleaseStub(logger, lock)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.fetchAndReleaseReturns.result1, fake.fetchAndReleaseReturns.result2
}

func (fake *FakeLockDB) FetchAndReleaseCallCount() int {
	fake.fetchAndReleaseMutex.RLock()
	defer fake.fetchAndReleaseMutex.RUnlock()
	return len(fake.fetchAndReleaseArgsForCall)
}

func (fake *FakeLockDB) FetchAndReleaseArgsForCall(i int) (lager.Logger, *db.Lock) {
	fake.fetchAndReleaseMutex.RLock()
	defer fake.fetchAndReleaseMutex.RUnlock()
	return fake.fetchAndReleaseArgsForCall[i].logger, fake.fetchAndReleaseArgsForCall[i].lock
}

func (fake *FakeLockDB) FetchAndReleaseReturns(result1 bool, result2 error) {
	fake.FetchAndReleaseStub = nil
	fake.fetchAndReleaseReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeLockDB) FetchAndReleaseReturnsOnCall(i int, result1 bool, result2 error) {
	fake.FetchAndReleaseStub = nil
	if fake.fetchAndReleaseReturnsOnCall == nil {
		fake.fetchAndReleaseReturnsOnCall = make(map[int]struct {
			result1 bool
			result2 error
		})
	}
	fake.fetchAndReleaseReturnsOnCall[i] = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeLockDB) FetchAll(logger lager.Logger, lockType string) ([]*db.Lock, error) {
	fake.fetchAllMutex.Lock()
	ret, specificReturn := fake.fetchAllReturnsOnCall[len(fake.fetchAllArgsForCall)]
	fake.fetchAllArgsForCall = append(fake.fetchAllArgsForCall, struct {
		logger   lager.Logger
		lockType string
	}{logger, lockType})
	fake.recordInvocation("FetchAll", []interface{}{logger, lockType})
	fake.fetchAllMutex.Unlock()
	if fake.FetchAllStub != nil {
		return fake.FetchAllStub(logger, lockType)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.fetchAllReturns.result1, fake.fetchAllReturns.result2
}

func (fake *FakeLockDB) FetchAllCallCount() int {
	fake.fetchAllMutex.RLock()
	defer fake.fetchAllMutex.RUnlock()
	return len(fake.fetchAllArgsForCall)
}

func (fake *FakeLockDB) FetchAllArgsForCall(i int) (lager.Logger, string) {
	fake.fetchAllMutex.RLock()
	defer fake.fetchAllMutex.RUnlock()
	return fake.fetchAllArgsForCall[i].logger, fake.fetchAllArgsForCall[i].lockType
}

func (fake *FakeLockDB) FetchAllReturns(result1 []*db.Lock, result2 error) {
	fake.FetchAllStub = nil
	fake.fetchAllReturns = struct {
		result1 []*db.Lock
		result2 error
	}{result1, result2}
}

func (fake *FakeLockDB) FetchAllReturnsOnCall(i int, result1 []*db.Lock, result2 error) {
	fake.FetchAllStub = nil
	if fake.fetchAllReturnsOnCall == nil {
		fake.fetchAllReturnsOnCall = make(map[int]struct {
			result1 []*db.Lock
			result2 error
		})
	}
	fake.fetchAllReturnsOnCall[i] = struct {
		result1 []*db.Lock
		result2 error
	}{result1, result2}
}

func (fake *FakeLockDB) Count(logger lager.Logger, lockType string) (int, error) {
	fake.countMutex.Lock()
	ret, specificReturn := fake.countReturnsOnCall[len(fake.countArgsForCall)]
	fake.countArgsForCall = append(fake.countArgsForCall, struct {
		logger   lager.Logger
		lockType string
	}{logger, lockType})
	fake.recordInvocation("Count", []interface{}{logger, lockType})
	fake.countMutex.Unlock()
	if fake.CountStub != nil {
		return fake.CountStub(logger, lockType)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.countReturns.result1, fake.countReturns.result2
}

func (fake *FakeLockDB) CountCallCount() int {
	fake.countMutex.RLock()
	defer fake.countMutex.RUnlock()
	return len(fake.countArgsForCall)
}

func (fake *FakeLockDB) CountArgsForCall(i int) (lager.Logger, string) {
	fake.countMutex.RLock()
	defer fake.countMutex.RUnlock()
	return fake.countArgsForCall[i].logger, fake.countArgsForCall[i].lockType
}

func (fake *FakeLockDB) CountReturns(result1 int, result2 error) {
	fake.CountStub = nil
	fake.countReturns = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *FakeLockDB) CountReturnsOnCall(i int, result1 int, result2 error) {
	fake.CountStub = nil
	if fake.countReturnsOnCall == nil {
		fake.countReturnsOnCall = make(map[int]struct {
			result1 int
			result2 error
		})
	}
	fake.countReturnsOnCall[i] = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *FakeLockDB) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.lockMutex.RLock()
	defer fake.lockMutex.RUnlock()
	fake.releaseMutex.RLock()
	defer fake.releaseMutex.RUnlock()
	fake.fetchMutex.RLock()
	defer fake.fetchMutex.RUnlock()
	fake.fetchAndReleaseMutex.RLock()
	defer fake.fetchAndReleaseMutex.RUnlock()
	fake.fetchAllMutex.RLock()
	defer fake.fetchAllMutex.RUnlock()
	fake.countMutex.RLock()
	defer fake.countMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeLockDB) recordInvocation(key string, args []interface{}) {
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

var _ db.LockDB = new(FakeLockDB)
