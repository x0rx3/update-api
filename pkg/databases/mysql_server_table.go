package databases

import (
	"database/sql"
	"errors"
	"fmt"
	"update-api/pkg/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MySqlServerTable struct {
	db *gorm.DB
}

func NewMySQlServerTable(db *gorm.DB) *MySqlServerTable {
	return &MySqlServerTable{
		db: db,
	}
}

func (inst *MySqlServerTable) SelectAll() ([]models.Server, error) {
	var servers []models.Server
	if result := inst.db.Find(&servers); result.Error != nil {
		if errors.Is(result.Error, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, result.Error
	}
	return servers, nil
}

func (inst *MySqlServerTable) SelectOne(uuid string) (*models.Server, error) {
	server := &models.Server{}
	if result := inst.db.First(server).Where("uuid = ?", uuid); result.Error != nil {
		if errors.Is(result.Error, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, result.Error
	}

	return server, nil
}

func (inst *MySqlServerTable) Delete(uuid string) error {
	return inst.db.Delete(&models.Server{}).Where("uuid = ?", uuid).Error
}

func (inst *MySqlServerTable) Update(server *models.Server) error {
	oldServer := &models.Server{}

	if server.UUID == "" {
		return fmt.Errorf("server uuid can't be null")
	}

	selectResult := inst.db.First(oldServer).Where("uuid = ?", server.UUID)
	if selectResult.Error != nil {
		return selectResult.Error
	}

	updatesField := map[string]any{}

	if oldServer.Url != server.Url {
		updatesField["url"] = server.Url
	}

	if oldServer.Login != server.Login {
		updatesField["login"] = server.Login
	}

	if oldServer.Password != server.Password {
		updatesField["password"] = server.Password
	}

	if oldServer.Name != server.Name {
		updatesField["name"] = server.Name
	}

	query := inst.db.Where("uuid = ?", server.UUID).Updates(updatesField)

	return query.Error
}

func (inst *MySqlServerTable) Create(server *models.Server) (string, error) {
	server.UUID = uuid.NewString()
	if result := inst.db.Create(server); result.Error != nil {
		return "", result.Error
	}
	return server.UUID, nil
}
