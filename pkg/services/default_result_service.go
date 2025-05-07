package services

import (
	"update-api/pkg/databases"
	"update-api/pkg/models"
)

type DefaultResultService struct {
	resultTable databases.ResultTable
}

func NewDefailtResultService(resultTable databases.ResultTable) *DefaultResultService {
	return &DefaultResultService{
		resultTable: resultTable,
	}
}

func (inst *DefaultResultService) Get(uuid string) (*models.Result, error) {
	return inst.resultTable.SelectOne(uuid)
}

func (inst *DefaultResultService) GetAll() ([]models.Result, error) {
	return inst.resultTable.SelectAll()
}
