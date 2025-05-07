package models

type Server struct {
	UUID          string `json:"uuid" gorm:"primaryKey;column:uuid"`
	Url           string `json:"url" gorm:"column:url"`
	Name          string `json:"name" gorm:"column:name"`
	Login         string `json:"login" gorm:"column:login"`
	Password      string `json:"password" gorm:"column:password"`
	SoftVersion   string `json:"soft_version" gorm:"column:soft_version"`
	RulesStatus   string `json:"rules_status" gorm:"column:rules_status"`
	MalwareStatus string `json:"malware_status" gorm:"column:malware_status"`
}

func (isnt Server) TableName() string {
	return "servers"
}
