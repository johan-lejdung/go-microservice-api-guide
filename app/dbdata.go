package app

import (
	"time"
)

type DbData struct {
	ID   int       `json:"id"`
	Date time.Time `json:"date"`
	Name string    `json:"name"`
}
