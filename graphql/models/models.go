package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username     string `json:"username" validate:"required"`
	Email        string `json:"email" validate:"required"`
	HashPassword string `json:"-"`
}
type Company struct {
	gorm.Model
	Name     string `json:"name" validate:"required"`
	Location string `json:"location" validate:"required"`
}
type Job struct {
	gorm.Model
	Company  Company `json:"-" gorm:"ForeignKey:cid"`
	Cid      string  `json:"cid"`
	JobTitle string  `json:"jobTitle" validate:"required"`
	Salary   string  `json:"salary" validate:"required"`
}
