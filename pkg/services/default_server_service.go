package services

import (
	"update-api/pkg/databases"
	"update-api/pkg/models"
)

type DefauleServerService struct {
	serverTable databases.ServerTable
}

func NewDefaultServerService(serverTable databases.ServerTable) *DefauleServerService {
	return &DefauleServerService{
		serverTable: serverTable,
	}
}

func (inst *DefauleServerService) Get(uuid string) (*models.Server, error) {
	return inst.serverTable.SelectOne(uuid)
}

func (inst *DefauleServerService) GetAll() ([]models.Server, error) {
	return inst.serverTable.SelectAll()
}

func (inst *DefauleServerService) Update(server *models.Server) error {
	return inst.serverTable.Update(server)
}

func (inst *DefauleServerService) Delete(uuid string) error {
	return inst.serverTable.Delete(uuid)
}

func (inst *DefauleServerService) Create(server *models.Server) (string, error) {
	return inst.serverTable.Create(server)
}
