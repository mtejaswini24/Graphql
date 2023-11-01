package service

import (
	"fmt"
	"graphql/graph/model"
	"graphql/models"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

func (s *Conn) CreateUser(nu model.NewUser) (*model.User, error) {
	//method that creates a new record in  db
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(nu.Password), bcrypt.DefaultCost)
	if err != nil {
		return &model.User{}, fmt.Errorf("generating password hash: %w", err)
	}
	//prepare user record
	u := models.User{
		Username:     nu.Username,
		Email:        nu.Email,
		HashPassword: string(hashedPass),
	}
	//calling default create method
	err = s.db.Create(&u).Error
	if err != nil {
		return &model.User{}, err
	}
	uid := strconv.FormatUint(uint64(u.ID), 10)
	u1 := model.User{
		ID:       uid,
		Username: u.Username,
		Email:    u.Email,
	}
	// Successfully created the record, return the user.
	return &u1, nil
}
