package models

import "time"

type Society struct {
	Id             int64  `json:"id" gorm:"primaryKey"`
	Name           string `json:"name"`
	Address        string `json:"address"`
	RegistrationNo string `json:"registration_no"`
}
type Flat struct {
	Id         int64   `json:"id" gorm:"primaryKey"`
	FlatNumber string  `json:"flat_number"`
	OwnerName  string  `json:"owner_name"`
	AreaSqFt   float64 `json:"area_sqft" gorm:"column:area_sqft"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
