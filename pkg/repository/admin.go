package repository

import (
	"github/MohdAhzan/internPortalAPP/pkg/repository/interfaces"
	"github/MohdAhzan/internPortalAPP/pkg/utils/domain"
	"github/MohdAhzan/internPortalAPP/pkg/utils/models"

	"gorm.io/gorm"
)

type AdminRepository struct{
  db *gorm.DB
}

func NewAdminRepository (dB *gorm.DB)interfaces.AdminRepository{
  return &AdminRepository{ 
    db: dB,
  }
}

func (u *AdminRepository) FetchAdminDetailsByEmail(email string)(domain.AdminDetails,error){
  
  var model domain.AdminDetails

	if err := u.db.Raw("select * from admins where email = ? ", email).Scan(&model).Error; err != nil {
		return domain.AdminDetails{}, err
  }
  return model,nil
}


func (u *AdminRepository) UserSignup(userModel models.UserSignup) (models.UserDetailsResponse, error) {
  var userDetails models.UserDetailsResponse
  // domain.User
  err := u.db.Raw("insert into users (name,email,password,role,phone,created_at,updated_at) values (?,?,?,?,?,?,?) RETURNING id,name,email,phone,role", userModel.Name, userModel.Email, userModel.Password,userModel.Role, userModel.Phone,userModel.CreatedAt,userModel.UpdatedAt).Scan(&userDetails).Error
  if err != nil {
    return models.UserDetailsResponse{}, err
  }
  return userDetails, nil

}

func (u *AdminRepository) AddDoctorDetail(userModel models.UserSignup,details domain.DoctorDetails) (models.UserDetailsResponse,error){
  var userDetails models.UserDetailsResponse
  err := u.db.Raw("insert into users (name,email,password,role,phone,created_at,updated_at) values (?,?,?,?,?,?,?) RETURNING id,name,email,phone,role", userModel.Name, userModel.Email, userModel.Password,userModel.Role, userModel.Phone,userModel.CreatedAt,userModel.UpdatedAt).Scan(&userDetails).Error
  if err != nil {
    return models.UserDetailsResponse{}, err
  }
  err=u.db.Exec("insert into doctor_details (user_id,department,created_at,updated_at) values (?,?,?,?)",userDetails.ID,userModel.Department,userModel.CreatedAt,userModel.UpdatedAt).Error
  if err != nil {
    return models.UserDetailsResponse{}, err
  }
  return userDetails,nil
}

