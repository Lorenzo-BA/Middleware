package models

import (
	"github.com/gofrs/uuid"
	"time"
)

type Rating struct {
	Id         *uuid.UUID `json:"id"`
	RatingDate time.Time     `json:"rating_date"`
	Comment    string     `json:"comment"`
	SongId     *uuid.UUID     `json:"song_id"`
	UserId     *uuid.UUID     `json:"user_id"`
	Rating     int        	`json:"rating"`
}