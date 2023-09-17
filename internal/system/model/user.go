package model

import (
	"database/sql"
	"time"
)

type User struct {
	Id        uint   `gorm:"primaryKey;autoIncrement"`
	UserName  string `gorm:"not null"`
	Password  string `gorm:"not null"`
	Mobile    string `gorm:"not null"`
	Email     string `gorm:"null"`
	TenantId  string `gorm:"not null"`
	Status    int    `gorm:"not null"`
	Birthday  sql.NullTime
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *User) TableName() string {
	return "users"
}
