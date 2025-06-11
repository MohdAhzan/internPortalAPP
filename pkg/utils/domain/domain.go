package domain

import (
	"time"

	"github.com/google/uuid"
)

type AdminDetails struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
  Name      string   `gorm:"type:varchar(100);not null" json:"name"`
	Email     string    `gorm:"type:varchar(100);unique;not null" json:"email"`
	Password  string    `gorm:"type:varchar(100);not null" json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Role string

const (
	Admin Role = "admin"
	Receptionist Role = "receptionist"
	Doctor       Role = "doctor"
)

type User struct {
	ID        uuid.UUID     `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
  Name      string         `gorm:"type:varchar(100);unique;not null" json:"name"`
	Email     string        `gorm:"type:varchar(100);unique;not null" json:"email"`
	Password  string        `gorm:"type:varchar(100);not null" json:"password"`
	Role      Role          `gorm:"type:varchar(20);not null" json:"role"`
	Phone     string        `gorm:"type:varchar(15);not null" json:"phone"`
	Doctor    *DoctorDetails `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"doctor,omitempty"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
}

type DoctorDetails struct {
	ID         uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	UserID     uuid.UUID `gorm:"type:uuid;uniqueIndex;not null" json:"user_id"`
	Department string    `gorm:"type:varchar(100);not null" json:"department"`
	Available  bool      `json:"available"`
	UpdatedAt  time.Time `json:"updated_at"`
}

//
// type Patient struct{
// }
