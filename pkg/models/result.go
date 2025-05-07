package models

import "time"

type Result struct {
	UUID       string    `json:"uuid" gorm:"column:uuid;primarykey"`
	ServerUUID string    `json:"server_uuid" gorm:"column:server_uuid"`
	TimeStart  time.Time `json:"time_start" gorm:"column:time_start"`
	TimeEnd    time.Time `json:"time_end" gorm:"column:time_end"`
	Errors     string    `json:"errors" gorm:"column:errors"`
	Malware    string    `json:"malware" gorm:"column:malware"`
	Rules      string    `json:"rules" gorm:"column:rules"`
}

func (inst Result) TableName() string {
	return "result"
}
