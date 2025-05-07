package databases

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MySqlUpdateDatabase struct {
	serverTable ServerTable
	resultTable ResultTable
}

func NewMySqlUpdateDatabase() *MySqlUpdateDatabase {
	return &MySqlUpdateDatabase{}
}

func (inst *MySqlUpdateDatabase) Connect(dsn string) error {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	inst.serverTable = NewMySQlServerTable(db)
	inst.resultTable = NewMySqlResultTable(db)

	return nil
}

func (inst *MySqlUpdateDatabase) ServerTable() ServerTable {
	return inst.serverTable
}

func (inst *MySqlUpdateDatabase) ResultTable() ResultTable {
	return inst.resultTable
}
