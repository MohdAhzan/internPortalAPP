package usecase

import (
	"errors"
	"github/MohdAhzan/internPortalAPP/pkg/config"
	"github/MohdAhzan/internPortalAPP/pkg/helper"
	"github/MohdAhzan/internPortalAPP/pkg/repository/interfaces"
	usecase "github/MohdAhzan/internPortalAPP/pkg/usecase/interfaces"
	response "github/MohdAhzan/internPortalAPP/pkg/utils/Response"
	"github/MohdAhzan/internPortalAPP/pkg/utils/domain"
	"github/MohdAhzan/internPortalAPP/pkg/utils/models"
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type AdminUsecase struct{
  recepRepo interfaces.ReceptionRepository
  adminRepo interfaces.AdminRepository
  cfg config.Config
  h *helper.Helper
}

func NewAdminUsecase(adminRepo interfaces.AdminRepository,recepRepo interfaces.ReceptionRepository,cfg config.Config,h *helper.Helper )usecase.AdminUsecase{
  return AdminUsecase{
    recepRepo: recepRepo,
    adminRepo: adminRepo,
    cfg: cfg,
    h: h,
  }
}

func (a AdminUsecase) AdminLogin(model models.AdminLogin)(models.AdminLoginResponse,error){
 
  adminCompareDetails, err := a.adminRepo.FetchAdminDetailsByEmail(model.Email)
	if err != nil {
		return models.AdminLoginResponse{}, err
	}


	err = bcrypt.CompareHashAndPassword([]byte(adminCompareDetails.Password), []byte(model.Password))
	if err != nil {
		return models.AdminLoginResponse{}, err
	}


	accessToken,refreshToken, err := a.h.GenerateAdminToken(response.UserDetailsResponse{
    ID: adminCompareDetails.ID ,
    Email: adminCompareDetails.Email,
    Role: domain.Admin,
  },a.cfg)

	if err != nil {
		return models.AdminLoginResponse{}, err
	}
    
return  models.AdminLoginResponse{
    ID: adminCompareDetails.ID,
    Name: adminCompareDetails.Name,
    Email: adminCompareDetails.Email,
    AccessToken: accessToken,
    RefreshToken: refreshToken,
  },nil


}


func (a AdminUsecase)UserSignup(userModel models.UserSignup)(models.UserDetailsResponse,error){
  

  userExist,err := a.recepRepo.CheckUserAvailability(userModel.Email)
  if userExist {
    return models.UserDetailsResponse{},errors.Join(errors.New("user already exist"),err)
  }

  if userModel.Password != userModel.ConfirmPassword {
    return models.UserDetailsResponse{},errors.New("password should'nt be different")
  }

  hashedPassword, err := a.h.PasswordHashing(userModel.Password)
  if err != nil {
    return models.UserDetailsResponse{},errors.Join(errors.New("error hashing password"),err)
  }

  userModel.Password = hashedPassword
  userModel.CreatedAt= time.Now()
  userModel.UpdatedAt= time.Now()

  if userModel.Role == domain.Receptionist{
    res,err:=  a.adminRepo.UserSignup(userModel) 
    if err!=nil{
      return models.UserDetailsResponse{},errors.Join(errors.New("error signin receptionist"),err)
    }
    return  res,nil
  }


  if userModel.Role != domain.Doctor{
    return models.UserDetailsResponse{},errors.New("enter only valid roles")
  }

  var docDetails domain.DoctorDetails
  docDetails.Department = userModel.Department 


  if  userModel.Department == ""{
    return models.UserDetailsResponse{},errors.New("enter department for doctors")
  }

   log.Println("CHCEKKKK",userModel) 
  res,err:=a.adminRepo.AddDoctorDetail(userModel,docDetails)
  if err!=nil{
    return models.UserDetailsResponse{}  ,errors.Join(errors.New("error signin doctor"),err)
  }
  return  res,nil
}


