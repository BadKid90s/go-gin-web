package common

import (
	"fmt"
	"github.com/spf13/viper"
	"go-gin-demo/pkg/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Repository struct {
	DB     *gorm.DB
	Logger *log.Logger
}

func NewRepository(db *gorm.DB, logger *log.Logger) *Repository {
	return &Repository{
		DB:     db,
		Logger: logger,
	}
}

func NewDB(conf *viper.Viper) *gorm.DB {
	gormConfig := &gorm.Config{
		PrepareStmt: true,
		Plugins:     map[string]gorm.Plugin{},
	}
	db, err := gorm.Open(mysql.Open(conf.GetString("data.mysql.user")), gormConfig)
	if err != nil {
		panic(fmt.Sprintf("mysql error: %s", err.Error()))
	}
	return db
}
