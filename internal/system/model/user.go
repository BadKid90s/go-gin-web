package model

import "time"

type User struct {
	Id        uint
	CreatedAt time.Time
	UpdatedAt time.Time

	LoginName *string
	Password  *string
	UserName  *string
	Mobile    *string
	Email     *string
	Enabled   bool
}

func (u *User) TableName() string {
	return "sys_users"
}
