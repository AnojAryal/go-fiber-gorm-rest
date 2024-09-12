package models

import "time"

type User struct {
	ID        uint `json:"id" gorm:"primarykey"`
	CreatedAt time.Time
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}
