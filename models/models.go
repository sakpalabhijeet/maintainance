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

type Owners struct {
    Id        int64    `json:"id" gorm:"primaryKey"`
    UserId   int64     `json:"user_id"`
    FullName  string    `json:"full_name"`
    Phone     string    `json:"phone"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}


type Users struct {
    Id          uint      `json:"id"`
    Email        string    `json:"email"`
    PasswordHash string    `json:"password_hash" gorm:"column:password_hash"`
    Role         string    `json:"role"`
    CreatedAt    time.Time
}

type Bill struct {
    BillID          int64     `json:"bill_id" gorm:"primaryKey"`
    FlatID      int64     `json:"flat_id" gorm:"column:flat_id;not null"`
    BillMonth   time.Time `json:"bill_month" gorm:"column:bill_month;not null"`
    TotalAmount float64   `json:"total_amount" gorm:"column:total_amount;not null"`
    Status      string    `json:"status" gorm:"column:status;default:'pending'"`
    DueDate     time.Time `json:"due_date" gorm:"column:due_date;not null"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
}

type BillItem struct {
	ID          int64     `json:"id" gorm:"primaryKey"`
	BillID      int64     `json:"bill_id" gorm:"not null;index"`
	Description string    `json:"description" gorm:"type:varchar(100);not null"`
	Rate        float64   `json:"rate" gorm:"type:numeric(10,2);not null"`
	Quantity    float64   `json:"quantity" gorm:"type:numeric(10,2);not null"`
	Amount      float64   `json:"amount" gorm:"type:numeric(10,2);not null"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type MaintenanceRate struct {
	ID              int64     `json:"id" gorm:"primaryKey"`
	RateType        string    `json:"rate_type"` // maintenance, sinking_fund
	RatePerSqFt     float64   `json:"rate_per_sq_ft"`
	GSTPercent      float64   `json:"gst_percent"` // 18 for 18%
	IsActive        bool      `json:"is_active"`
	StartDate       time.Time `json:"start_date"`
	EndDate         *time.Time `json:"end_date"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
