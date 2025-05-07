package services

import "update-api/pkg/databases"

type DefaultServiceCollector struct {
	serverService ServerService
	resultService ResultService
}

func NewDefaultServiceCollector(updateDB databases.UpdateDatabase) *DefaultServiceCollector {
	return &DefaultServiceCollector{
		serverService: NewDefaultServerService(updateDB.ServerTable()),
		resultService: NewDefailtResultService(updateDB.ResultTable()),
	}
}

func (inst *DefaultServiceCollector) ServerService() ServerService {
	return inst.serverService
}

func (inst *DefaultServiceCollector) ResultService() ResultService {
	return inst.resultService
}
