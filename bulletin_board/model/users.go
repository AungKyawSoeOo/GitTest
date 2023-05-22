package model

import "time"

type User struct {
	Id           int    `gorm:"type:int;primary_key"`
	Username     string `gorm:"type:varchar(255);not null"`
	Email        string `gorm:"uniqueIndex;not null"`
	Password     string `gorm:"not null"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	CreateUserId int     `gorm:"not null"`
	UpdateUserId int     `gorm:"not null"`
	Posts        []Posts `gorm:"foreignkey:CreateUserId"`
}
