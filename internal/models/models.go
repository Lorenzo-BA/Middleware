package models

import (
	"github.com/gofrs/uuid"
	"time"
)

type Song struct {
	Id             *uuid.UUID `json:"id"`
	Content        string     `json:"content"`
	Title          string     `json:"title"`
	File_name      string     `json:"file_name"`
	Artist         string     `json:"artist"`
	Published_date time.Time  `json:"Published_date"`
	Code           int        `json:"Code"`
}
