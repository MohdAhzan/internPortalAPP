package models

import (
	"github/MohdAhzan/internPortalAPP/pkg/utils/domain"
	"time"

	"github.com/google/uuid"
)

type AdminLogin struct{
    Email string
    Password string
}

type UserSignup struct{
  Name     string
	Email    string 
  Phone    string
	Password string 
  ConfirmPassword string 
  Role   domain.Role
  Department string 
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`

}

type UserLogin struct {
	Email    string 
	Password string 
}
type UserDetailsResponse struct{
	ID       uuid.UUID   
  Name      string
	Email    string 
  Phone  string
  Role   domain.Role
}

type TokenUsers struct {
	Users        UserDetailsResponse
	AccessToken  string
	RefreshToken string
}

