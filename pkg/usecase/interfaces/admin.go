package interfaces

import (
	"github/MohdAhzan/internPortalAPP/pkg/utils/models"
)

type AdminUsecase interface{
    AdminLogin(model models.AdminLogin)(models.AdminLoginResponse,error)
    UserSignup(userModel models.UserSignup)(models.UserDetailsResponse,error)
}
