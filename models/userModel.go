package models

import "time"

type User struct {
	Id          string    `json:"id"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	DateCreated time.Time `json:"date_created"`
}
