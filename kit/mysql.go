package kit

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func NewMySQLClient(config *Config) (*gorm.DB, error) {
	db, err := gorm.Open("mysql", config.MySQL.User+":"+config.MySQL.Password+
		"@tcp("+config.MySQL.Host+":"+config.MySQL.Port+")/"+config.MySQL.Database+
		"?charset=utf8mb4&parseTime=True&loc=Local&timeout=90s")
	if err != nil {
		return nil, err
	}

	return db, nil
}
