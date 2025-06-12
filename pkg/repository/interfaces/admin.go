package interfaces

import (
	"github/MohdAhzan/internPortalAPP/pkg/utils/domain"
	"github/MohdAhzan/internPortalAPP/pkg/utils/models"
)

type AdminRepository interface{
  FetchAdminDetailsByEmail(email string)(domain.AdminDetails,error)
  UserSignup(userModel models.UserSignup) (models.UserDetailsResponse, error) 
  AddDoctorDetail(userModel models.UserSignup,details domain.DoctorDetails)(models.UserDetailsResponse,error)

}
