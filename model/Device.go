package model

import "time"

type Device struct {
	ID           int64     `json:"id" gorm:"unique;primaryKey;autoIncrement"`
	Name         string    `json:"name" db:"name"`
	Brand        string    `json:"brand" db:"brand"`
	State        string    `json:"state" db:"state"`
	CreationDate time.Time `json:"createdAt" db:"creation_date"`
}
