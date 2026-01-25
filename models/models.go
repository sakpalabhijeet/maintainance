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
	OwnerId int64 `json:"owner_id"`
	AreaSqFt   float64 `json:"area_sq_ft" gorm:"column:area_sq_ft"`
	IsOccupied bool `json:"is_occupied"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Owners struct{
	OwnerId int64 `json:"id" gorm:"primaryKey"`
	UserId int64 `json:"user_id"`
	FullName string `json:"full_name"`
	Phone string `json:"phone"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Users struct {
    ID           uint      `json:"id"`
    Email        string    `json:"email"`
    PasswordHash string    `json:"password_hash" gorm:"column:password_hash"`
    Role         string    `json:"role"`
    CreatedAt    time.Time
}
