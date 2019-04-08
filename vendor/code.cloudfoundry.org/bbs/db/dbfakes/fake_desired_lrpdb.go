// Code generated by counterfeiter. DO NOT EDIT.
package dbfakes

import (
	"sync"

	"code.cloudfoundry.org/bbs/db"
	"code.cloudfoundry.org/bbs/models"
	"code.cloudfoundry.org/lager"
)

type FakeDesiredLRPDB struct {
	DesiredLRPsStub        func(logger lager.Logger, filter models.DesiredLRPFilter) ([]*models.DesiredLRP, error)
	desiredLRPsMutex       sync.RWMutex
	desiredLRPsArgsForCall []struct {
		logger lager.Logger
		filter models.DesiredLRPFilter
	}
	desiredLRPsReturns struct {
		result1 []*models.DesiredLRP
		result2 error
	}
	desiredLRPsReturnsOnCall map[int]struct {
		result1 []*models.DesiredLRP
		result2 error
	}
	DesiredLRPByProcessGuidStub        func(logger lager.Logger, processGuid string) (*models.DesiredLRP, error)
	desiredLRPByProcessGuidMutex       sync.RWMutex
	desiredLRPByProcessGuidArgsForCall []struct {
		logger      lager.Logger
		processGuid string
	}
	desiredLRPByProcessGuidReturns struct {
		result1 *models.DesiredLRP
		result2 error
	}
	desiredLRPByProcessGuidReturnsOnCall map[int]struct {
		result1 *models.DesiredLRP
		result2 error
	}
	DesiredLRPSchedulingInfosStub        func(logger lager.Logger, filter models.DesiredLRPFilter) ([]*models.DesiredLRPSchedulingInfo, error)
	desiredLRPSchedulingInfosMutex       sync.RWMutex
	desiredLRPSchedulingInfosArgsForCall []struct {
		logger lager.Logger
		filter models.DesiredLRPFilter
	}
	desiredLRPSchedulingInfosReturns struct {
		result1 []*models.DesiredLRPSchedulingInfo
		result2 error
	}
	desiredLRPSchedulingInfosReturnsOnCall map[int]struct {
		result1 []*models.DesiredLRPSchedulingInfo
		result2 error
	}
	DesireLRPStub        func(logger lager.Logger, desiredLRP *models.DesiredLRP) error
	desireLRPMutex       sync.RWMutex
	desireLRPArgsForCall []struct {
		logger     lager.Logger
		desiredLRP *models.DesiredLRP
	}
	desireLRPReturns struct {
		result1 error
	}
	desireLRPReturnsOnCall map[int]struct {
		result1 error
	}
	UpdateDesiredLRPStub        func(logger lager.Logger, processGuid string, update *models.DesiredLRPUpdate) (beforeDesiredLRP *models.DesiredLRP, err error)
	updateDesiredLRPMutex       sync.RWMutex
	updateDesiredLRPArgsForCall []struct {
		logger      lager.Logger
		processGuid string
		update      *models.DesiredLRPUpdate
	}
	updateDesiredLRPReturns struct {
		result1 *models.DesiredLRP
		result2 error
	}
	updateDesiredLRPReturnsOnCall map[int]struct {
		result1 *models.DesiredLRP
		result2 error
	}
	RemoveDesiredLRPStub        func(logger lager.Logger, processGuid string) error
	removeDesiredLRPMutex       sync.RWMutex
	removeDesiredLRPArgsForCall []struct {
		logger      lager.Logger
		processGuid string
	}
	removeDesiredLRPReturns struct {
		result1 error
	}
	removeDesiredLRPReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeDesiredLRPDB) DesiredLRPs(logger lager.Logger, filter models.DesiredLRPFilter) ([]*models.DesiredLRP, error) {
	fake.desiredLRPsMutex.Lock()
	ret, specificReturn := fake.desiredLRPsReturnsOnCall[len(fake.desiredLRPsArgsForCall)]
	fake.desiredLRPsArgsForCall = append(fake.desiredLRPsArgsForCall, struct {
		logger lager.Logger
		filter models.DesiredLRPFilter
	}{logger, filter})
	fake.recordInvocation("DesiredLRPs", []interface{}{logger, filter})
	fake.desiredLRPsMutex.Unlock()
	if fake.DesiredLRPsStub != nil {
		return fake.DesiredLRPsStub(logger, filter)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.desiredLRPsReturns.result1, fake.desiredLRPsReturns.result2
}

func (fake *FakeDesiredLRPDB) DesiredLRPsCallCount() int {
	fake.desiredLRPsMutex.RLock()
	defer fake.desiredLRPsMutex.RUnlock()
	return len(fake.desiredLRPsArgsForCall)
}

func (fake *FakeDesiredLRPDB) DesiredLRPsArgsForCall(i int) (lager.Logger, models.DesiredLRPFilter) {
	fake.desiredLRPsMutex.RLock()
	defer fake.desiredLRPsMutex.RUnlock()
	return fake.desiredLRPsArgsForCall[i].logger, fake.desiredLRPsArgsForCall[i].filter
}

func (fake *FakeDesiredLRPDB) DesiredLRPsReturns(result1 []*models.DesiredLRP, result2 error) {
	fake.DesiredLRPsStub = nil
	fake.desiredLRPsReturns = struct {
		result1 []*models.DesiredLRP
		result2 error
	}{result1, result2}
}

func (fake *FakeDesiredLRPDB) DesiredLRPsReturnsOnCall(i int, result1 []*models.DesiredLRP, result2 error) {
	fake.DesiredLRPsStub = nil
	if fake.desiredLRPsReturnsOnCall == nil {
		fake.desiredLRPsReturnsOnCall = make(map[int]struct {
			result1 []*models.DesiredLRP
			result2 error
		})
	}
	fake.desiredLRPsReturnsOnCall[i] = struct {
		result1 []*models.DesiredLRP
		result2 error
	}{result1, result2}
}

func (fake *FakeDesiredLRPDB) DesiredLRPByProcessGuid(logger lager.Logger, processGuid string) (*models.DesiredLRP, error) {
	fake.desiredLRPByProcessGuidMutex.Lock()
	ret, specificReturn := fake.desiredLRPByProcessGuidReturnsOnCall[len(fake.desiredLRPByProcessGuidArgsForCall)]
	fake.desiredLRPByProcessGuidArgsForCall = append(fake.desiredLRPByProcessGuidArgsForCall, struct {
		logger      lager.Logger
		processGuid string
	}{logger, processGuid})
	fake.recordInvocation("DesiredLRPByProcessGuid", []interface{}{logger, processGuid})
	fake.desiredLRPByProcessGuidMutex.Unlock()
	if fake.DesiredLRPByProcessGuidStub != nil {
		return fake.DesiredLRPByProcessGuidStub(logger, processGuid)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.desiredLRPByProcessGuidReturns.result1, fake.desiredLRPByProcessGuidReturns.result2
}

func (fake *FakeDesiredLRPDB) DesiredLRPByProcessGuidCallCount() int {
	fake.desiredLRPByProcessGuidMutex.RLock()
	defer fake.desiredLRPByProcessGuidMutex.RUnlock()
	return len(fake.desiredLRPByProcessGuidArgsForCall)
}

func (fake *FakeDesiredLRPDB) DesiredLRPByProcessGuidArgsForCall(i int) (lager.Logger, string) {
	fake.desiredLRPByProcessGuidMutex.RLock()
	defer fake.desiredLRPByProcessGuidMutex.RUnlock()
	return fake.desiredLRPByProcessGuidArgsForCall[i].logger, fake.desiredLRPByProcessGuidArgsForCall[i].processGuid
}

func (fake *FakeDesiredLRPDB) DesiredLRPByProcessGuidReturns(result1 *models.DesiredLRP, result2 error) {
	fake.DesiredLRPByProcessGuidStub = nil
	fake.desiredLRPByProcessGuidReturns = struct {
		result1 *models.DesiredLRP
		result2 error
	}{result1, result2}
}

func (fake *FakeDesiredLRPDB) DesiredLRPByProcessGuidReturnsOnCall(i int, result1 *models.DesiredLRP, result2 error) {
	fake.DesiredLRPByProcessGuidStub = nil
	if fake.desiredLRPByProcessGuidReturnsOnCall == nil {
		fake.desiredLRPByProcessGuidReturnsOnCall = make(map[int]struct {
			result1 *models.DesiredLRP
			result2 error
		})
	}
	fake.desiredLRPByProcessGuidReturnsOnCall[i] = struct {
		result1 *models.DesiredLRP
		result2 error
	}{result1, result2}
}

func (fake *FakeDesiredLRPDB) DesiredLRPSchedulingInfos(logger lager.Logger, filter models.DesiredLRPFilter) ([]*models.DesiredLRPSchedulingInfo, error) {
	fake.desiredLRPSchedulingInfosMutex.Lock()
	ret, specificReturn := fake.desiredLRPSchedulingInfosReturnsOnCall[len(fake.desiredLRPSchedulingInfosArgsForCall)]
	fake.desiredLRPSchedulingInfosArgsForCall = append(fake.desiredLRPSchedulingInfosArgsForCall, struct {
		logger lager.Logger
		filter models.DesiredLRPFilter
	}{logger, filter})
	fake.recordInvocation("DesiredLRPSchedulingInfos", []interface{}{logger, filter})
	fake.desiredLRPSchedulingInfosMutex.Unlock()
	if fake.DesiredLRPSchedulingInfosStub != nil {
		return fake.DesiredLRPSchedulingInfosStub(logger, filter)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.desiredLRPSchedulingInfosReturns.result1, fake.desiredLRPSchedulingInfosReturns.result2
}

func (fake *FakeDesiredLRPDB) DesiredLRPSchedulingInfosCallCount() int {
	fake.desiredLRPSchedulingInfosMutex.RLock()
	defer fake.desiredLRPSchedulingInfosMutex.RUnlock()
	return len(fake.desiredLRPSchedulingInfosArgsForCall)
}

func (fake *FakeDesiredLRPDB) DesiredLRPSchedulingInfosArgsForCall(i int) (lager.Logger, models.DesiredLRPFilter) {
	fake.desiredLRPSchedulingInfosMutex.RLock()
	defer fake.desiredLRPSchedulingInfosMutex.RUnlock()
	return fake.desiredLRPSchedulingInfosArgsForCall[i].logger, fake.desiredLRPSchedulingInfosArgsForCall[i].filter
}

func (fake *FakeDesiredLRPDB) DesiredLRPSchedulingInfosReturns(result1 []*models.DesiredLRPSchedulingInfo, result2 error) {
	fake.DesiredLRPSchedulingInfosStub = nil
	fake.desiredLRPSchedulingInfosReturns = struct {
		result1 []*models.DesiredLRPSchedulingInfo
		result2 error
	}{result1, result2}
}

func (fake *FakeDesiredLRPDB) DesiredLRPSchedulingInfosReturnsOnCall(i int, result1 []*models.DesiredLRPSchedulingInfo, result2 error) {
	fake.DesiredLRPSchedulingInfosStub = nil
	if fake.desiredLRPSchedulingInfosReturnsOnCall == nil {
		fake.desiredLRPSchedulingInfosReturnsOnCall = make(map[int]struct {
			result1 []*models.DesiredLRPSchedulingInfo
			result2 error
		})
	}
	fake.desiredLRPSchedulingInfosReturnsOnCall[i] = struct {
		result1 []*models.DesiredLRPSchedulingInfo
		result2 error
	}{result1, result2}
}

func (fake *FakeDesiredLRPDB) DesireLRP(logger lager.Logger, desiredLRP *models.DesiredLRP) error {
	fake.desireLRPMutex.Lock()
	ret, specificReturn := fake.desireLRPReturnsOnCall[len(fake.desireLRPArgsForCall)]
	fake.desireLRPArgsForCall = append(fake.desireLRPArgsForCall, struct {
		logger     lager.Logger
		desiredLRP *models.DesiredLRP
	}{logger, desiredLRP})
	fake.recordInvocation("DesireLRP", []interface{}{logger, desiredLRP})
	fake.desireLRPMutex.Unlock()
	if fake.DesireLRPStub != nil {
		return fake.DesireLRPStub(logger, desiredLRP)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.desireLRPReturns.result1
}

func (fake *FakeDesiredLRPDB) DesireLRPCallCount() int {
	fake.desireLRPMutex.RLock()
	defer fake.desireLRPMutex.RUnlock()
	return len(fake.desireLRPArgsForCall)
}

func (fake *FakeDesiredLRPDB) DesireLRPArgsForCall(i int) (lager.Logger, *models.DesiredLRP) {
	fake.desireLRPMutex.RLock()
	defer fake.desireLRPMutex.RUnlock()
	return fake.desireLRPArgsForCall[i].logger, fake.desireLRPArgsForCall[i].desiredLRP
}

func (fake *FakeDesiredLRPDB) DesireLRPReturns(result1 error) {
	fake.DesireLRPStub = nil
	fake.desireLRPReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeDesiredLRPDB) DesireLRPReturnsOnCall(i int, result1 error) {
	fake.DesireLRPStub = nil
	if fake.desireLRPReturnsOnCall == nil {
		fake.desireLRPReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.desireLRPReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeDesiredLRPDB) UpdateDesiredLRP(logger lager.Logger, processGuid string, update *models.DesiredLRPUpdate) (beforeDesiredLRP *models.DesiredLRP, err error) {
	fake.updateDesiredLRPMutex.Lock()
	ret, specificReturn := fake.updateDesiredLRPReturnsOnCall[len(fake.updateDesiredLRPArgsForCall)]
	fake.updateDesiredLRPArgsForCall = append(fake.updateDesiredLRPArgsForCall, struct {
		logger      lager.Logger
		processGuid string
		update      *models.DesiredLRPUpdate
	}{logger, processGuid, update})
	fake.recordInvocation("UpdateDesiredLRP", []interface{}{logger, processGuid, update})
	fake.updateDesiredLRPMutex.Unlock()
	if fake.UpdateDesiredLRPStub != nil {
		return fake.UpdateDesiredLRPStub(logger, processGuid, update)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.updateDesiredLRPReturns.result1, fake.updateDesiredLRPReturns.result2
}

func (fake *FakeDesiredLRPDB) UpdateDesiredLRPCallCount() int {
	fake.updateDesiredLRPMutex.RLock()
	defer fake.updateDesiredLRPMutex.RUnlock()
	return len(fake.updateDesiredLRPArgsForCall)
}

func (fake *FakeDesiredLRPDB) UpdateDesiredLRPArgsForCall(i int) (lager.Logger, string, *models.DesiredLRPUpdate) {
	fake.updateDesiredLRPMutex.RLock()
	defer fake.updateDesiredLRPMutex.RUnlock()
	return fake.updateDesiredLRPArgsForCall[i].logger, fake.updateDesiredLRPArgsForCall[i].processGuid, fake.updateDesiredLRPArgsForCall[i].update
}

func (fake *FakeDesiredLRPDB) UpdateDesiredLRPReturns(result1 *models.DesiredLRP, result2 error) {
	fake.UpdateDesiredLRPStub = nil
	fake.updateDesiredLRPReturns = struct {
		result1 *models.DesiredLRP
		result2 error
	}{result1, result2}
}

func (fake *FakeDesiredLRPDB) UpdateDesiredLRPReturnsOnCall(i int, result1 *models.DesiredLRP, result2 error) {
	fake.UpdateDesiredLRPStub = nil
	if fake.updateDesiredLRPReturnsOnCall == nil {
		fake.updateDesiredLRPReturnsOnCall = make(map[int]struct {
			result1 *models.DesiredLRP
			result2 error
		})
	}
	fake.updateDesiredLRPReturnsOnCall[i] = struct {
		result1 *models.DesiredLRP
		result2 error
	}{result1, result2}
}

func (fake *FakeDesiredLRPDB) RemoveDesiredLRP(logger lager.Logger, processGuid string) error {
	fake.removeDesiredLRPMutex.Lock()
	ret, specificReturn := fake.removeDesiredLRPReturnsOnCall[len(fake.removeDesiredLRPArgsForCall)]
	fake.removeDesiredLRPArgsForCall = append(fake.removeDesiredLRPArgsForCall, struct {
		logger      lager.Logger
		processGuid string
	}{logger, processGuid})
	fake.recordInvocation("RemoveDesiredLRP", []interface{}{logger, processGuid})
	fake.removeDesiredLRPMutex.Unlock()
	if fake.RemoveDesiredLRPStub != nil {
		return fake.RemoveDesiredLRPStub(logger, processGuid)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.removeDesiredLRPReturns.result1
}

func (fake *FakeDesiredLRPDB) RemoveDesiredLRPCallCount() int {
	fake.removeDesiredLRPMutex.RLock()
	defer fake.removeDesiredLRPMutex.RUnlock()
	return len(fake.removeDesiredLRPArgsForCall)
}

func (fake *FakeDesiredLRPDB) RemoveDesiredLRPArgsForCall(i int) (lager.Logger, string) {
	fake.removeDesiredLRPMutex.RLock()
	defer fake.removeDesiredLRPMutex.RUnlock()
	return fake.removeDesiredLRPArgsForCall[i].logger, fake.removeDesiredLRPArgsForCall[i].processGuid
}

func (fake *FakeDesiredLRPDB) RemoveDesiredLRPReturns(result1 error) {
	fake.RemoveDesiredLRPStub = nil
	fake.removeDesiredLRPReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeDesiredLRPDB) RemoveDesiredLRPReturnsOnCall(i int, result1 error) {
	fake.RemoveDesiredLRPStub = nil
	if fake.removeDesiredLRPReturnsOnCall == nil {
		fake.removeDesiredLRPReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.removeDesiredLRPReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeDesiredLRPDB) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.desiredLRPsMutex.RLock()
	defer fake.desiredLRPsMutex.RUnlock()
	fake.desiredLRPByProcessGuidMutex.RLock()
	defer fake.desiredLRPByProcessGuidMutex.RUnlock()
	fake.desiredLRPSchedulingInfosMutex.RLock()
	defer fake.desiredLRPSchedulingInfosMutex.RUnlock()
	fake.desireLRPMutex.RLock()
	defer fake.desireLRPMutex.RUnlock()
	fake.updateDesiredLRPMutex.RLock()
	defer fake.updateDesiredLRPMutex.RUnlock()
	fake.removeDesiredLRPMutex.RLock()
	defer fake.removeDesiredLRPMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeDesiredLRPDB) recordInvocation(key string, args []interface{}) {
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

var _ db.DesiredLRPDB = new(FakeDesiredLRPDB)
