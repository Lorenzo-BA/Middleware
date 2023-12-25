package models

import (
	"github.com/gofrs/uuid"
	"time"
)

type Song struct {
	Id             *uuid.UUID `json:"id"`
	Content        string     `json:"content"`
	Artist_name    string     `json:"artist name"`
	Music_title    string     `json:"music title"`
	File_name      string     `json:"file name"`
	Published_date time.Time  `json:"published date"`
	Code           int
}
