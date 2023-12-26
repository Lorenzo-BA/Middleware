package models

import (
	"github.com/gofrs/uuid"
	"time"
)

type Song struct {
	Id             *uuid.UUID `json:"id"`
	Content        string     `json:"content"`
	Music_title    string     `json:"title"`
	File_name      string     `json:"file_name"`
	Artist_name    string     `json:"artist"`
	Published_date time.Time  `json:"Published_date"`
	Code           int
}
