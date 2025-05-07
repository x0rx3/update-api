package databases

import (
	"database/sql"
	"errors"
	"update-api/pkg/models"

	"gorm.io/gorm"
)

type MySqlResultTable struct {
	db *gorm.DB
}

func NewMySqlResultTable(db *gorm.DB) *MySqlResultTable {
	return &MySqlResultTable{
		db: db,
	}
}

func (inst *MySqlResultTable) SelectAll() ([]models.Result, error) {
	var results []models.Result

	if result := inst.db.Find(results); result.Error != nil {
		if errors.Is(result.Error, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, result.Error
	}

	return results, nil
}

func (inst *MySqlResultTable) SelectOne(uuid string) (*models.Result, error) {
	result := &models.Result{}
	if queryResult := inst.db.First(result).Where("uuid = ?", uuid); queryResult.Error != nil {
		if errors.Is(queryResult.Error, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, queryResult.Error
	}

	return result, nil
}
