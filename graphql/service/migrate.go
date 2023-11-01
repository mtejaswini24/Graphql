package service

import (
	"errors"

	"gorm.io/gorm"
)

type Conn struct {

	// db is an instance of the SQLite database.
	db *gorm.DB
}

// NewService is the constructor for the Conn struct.
func NewService(db *gorm.DB) (*Conn, error) {

	// We check if the database instance is nil, which would indicate an issue.
	if db == nil {
		return nil, errors.New("please provide a valid connection")
	}

	// We initialize our service with the passed database instance.
	s := &Conn{db: db}
	return s, nil

}
