package models

import (
	"github.com/go-playground/validator"
	"github.com/gofrs/uuid"
	"time"
)

var validate = validator.New()

type Song struct {
	Id            *uuid.UUID `json:"id"`
	Title         string     `json:"title"`
	FileName      string     `json:"file_name"`
	Artist        string     `json:"artist"`
	PublishedDate time.Time  `json:"Published_date"`
}

type SongRequest struct {
	Title    string `json:"title" validate:"required"`
	FileName string `json:"file_name" validate:"required"`
	Artist   string `json:"artist" validate:"required"`
}

func ValidateSongRequest(song SongRequest) error {
	if err := validate.Struct(song); err != nil {
		return err
	}
	return nil
}
