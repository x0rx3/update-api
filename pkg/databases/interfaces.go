package databases

import "update-api/pkg/models"

type UpdateDatabase interface {
	Connect(dsn string) error
	ServerTable() ServerTable
	ResultTable() ResultTable
}

type ServerTable interface {
	SelectAll() ([]models.Server, error)
	SelectOne(uuid string) (*models.Server, error)
	Delete(uuid string) error
	Update(server *models.Server) error
	Create(server *models.Server) (string, error)
}

type ResultTable interface {
	SelectAll() ([]models.Result, error)
	SelectOne(uuid string) (*models.Result, error)
}
