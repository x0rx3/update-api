package services

import "update-api/pkg/models"

type ServiceCollector interface {
	ServerService() ServerService
	ResultService() ResultService
}

type ServerService interface {
	Get(uuid string) (*models.Server, error)
	GetAll() ([]models.Server, error)
	Update(server *models.Server) error
	Delete(uuid string) error
	Create(server *models.Server) (string, error)
}

type ResultService interface {
	Get(uuid string) (*models.Result, error)
	GetAll() ([]models.Result, error)
}
