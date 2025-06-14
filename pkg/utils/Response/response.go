package response

import (
	"github/MohdAhzan/internPortalAPP/pkg/utils/domain"

	"github.com/google/uuid"
)

type Response struct {
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
	Error      interface{} `json:"error"`
}

func ClientResponse(statusCode int, message string, data interface{}, err interface{}) Response {
	return Response{
		StatusCode: statusCode,
		Message:    message,
		Data:       data,
		Error:      err,
	}
}

type UserDetailsResponse struct{
	ID       uuid.UUID `gorm:"primaryKey"`
	Email    string `gorm:"unique;not null"`
	Role     domain.Role `gorm:"not null"` 
}
