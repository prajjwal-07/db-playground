package model

import (
	"time"
)

type User struct {
	Id        string `gorm:"primary_key;type:uuid;default:gen_random_uuid()"`
	Name      string
	Email     string
	Age       int
	CreatedAt time.Time
	DeletedAt time.Time
	UpdatedAt time.Time
}
