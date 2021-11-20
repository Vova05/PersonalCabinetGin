package repository

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const(
	usersTable = "users"
	applicationsTable = "applications"
)

type Config struct {
	MdbUser string
	MdbPass string
	MdbHost string
	MdbName string
}

func NewMySQLDB(cfg Config) (*gorm.DB, error ){



	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True&loc=Local", cfg.MdbUser, cfg.MdbPass, cfg.MdbHost, cfg.MdbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil,err
	}

	sqlDB, err2 := db.DB()
	err2 = sqlDB.Ping()
	if err2 != nil{
		return nil, err2
	}
	return db, nil
}
