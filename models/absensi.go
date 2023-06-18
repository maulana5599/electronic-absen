package models

import "time"

type Absensi struct {
	ID              int
	Nis             int
	MatapelajaranId int
	GuruId          int
	IsHadir         bool
	CreatedAt       time.Time
	UpdateAt        time.Time
}

func (Absensi) TableName() string {
	return "absensi_t"
}
