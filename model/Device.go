package model

import "time"

const (
	AVAILABLE = iota
	INUSE
	INACTIVE
)

type Device struct {
	ID           int64     `json:"id" db:"id"`
	Name         string    `json:"name" db:"name"`
	Brand        string    `json:"brand" db:"brand"`
	State        int       `json:"state" db:"state"`
	CreationDate time.Time `json:"creation_date" db:"creation_date"`
}
