package models

import "time"

type User struct {
	Id         int       `json:"id"`
	Name       string    `json:"name" binding:"required"`
	Email      string    `json:"email" binding:"required"`
	Password   string    `json:"password" binding:"required"`
	Created_at time.Time `json:"created_at"`
}
