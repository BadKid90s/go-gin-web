package common

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type Repository struct {
	db     *gorm.DB
	logger *log.Logger
}

func NewRepository(db *gorm.DB, logger *log.Logger) *Repository {
	return &Repository{
		db:     db,
		logger: logger,
	}
}

func NewDB(conf *viper.Viper) *gorm.DB {
	db, err := gorm.Open(mysql.Open(conf.GetString("data.mysql.user")), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("mysql error: %s", err.Error()))
	}
	return db
}
