package interfaces

import "github/MohdAhzan/internPortalAPP/pkg/utils/models"

type AdminUsecase interface{
    UserSignup(userModel models.UserSignup)(models.UserDetailsResponse,error)
}
