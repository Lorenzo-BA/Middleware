package models

import (
	"github.com/go-playground/validator"
	"github.com/gofrs/uuid"
	"time"
)

var validate = validator.New()

type User struct {
	Id              *uuid.UUID `json:"id"`
	InscriptionDate time.Time  `json:"inscription_date"`
	Name            string     `json:"name"`
	Username        string     `json:"username"`
}

type UserRequest struct {
	Name     string `json:"name" validate:"required"`
	Username string `json:"username" validate:"required"`
}

func ValidateUserRequest(rating UserRequest) error {
	if err := validate.Struct(rating); err != nil {
		return err
	}
	return nil
}
