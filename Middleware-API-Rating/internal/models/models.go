package models

import (
	"github.com/go-playground/validator"
	"github.com/gofrs/uuid"
	"time"
)

var validate = validator.New()

type Rating struct {
	Id         *uuid.UUID `json:"id"`
	RatingDate time.Time  `json:"rating_date"`
	Comment    string     `json:"comment"`
	SongId     *uuid.UUID `json:"song_id"`
	UserId     *uuid.UUID `json:"user_id"`
	Rating     int        `json:"rating"`
}

type RatingCreateRequest struct {
	Comment string `json:"comment" validate:"required"`
	UserId  string `json:"user_id" validate:"required"`
	Rating  int    `json:"rating" validate:"required,min=1,max=5"`
}

type RatingUpdateRequest struct {
	Comment string `json:"comment" validate:"required"`
	Rating  int    `json:"rating" validate:"required,min=1,max=5"`
}

func ValidateRatingCreateRequest(rating RatingCreateRequest) error {
	if err := validate.Struct(rating); err != nil {
		return err
	}
	if _, err := uuid.FromString(rating.UserId); err != nil {
		return err
	}
	return nil
}

func ValidateRatingUpdateRequest(rating RatingUpdateRequest) error {
	if err := validate.Struct(rating); err != nil {
		return err
	}
	return nil
}
