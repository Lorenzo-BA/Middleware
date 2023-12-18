package models

import (
	"github.com/gofrs/uuid"
	"time"
)

type User struct {
	Id              *uuid.UUID `json:"id"`
	InscriptionDate time.Time  `json:"inscription_date"`
	Name            string     `json:"name"`
	Username        string     `json:"username"`
}
