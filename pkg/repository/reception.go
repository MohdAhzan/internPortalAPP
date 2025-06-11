package repository

import (
	"github/MohdAhzan/internPortalAPP/pkg/repository/interfaces"

	"gorm.io/gorm"
)

type ReceptionRepository struct{
  db *gorm.DB
}

func NewReceptionRepository (dB *gorm.DB)interfaces.ReceptionRepository{
  return &ReceptionRepository{ 
    db: dB,
  }
}

// func (r *ReceptionRepository)  AddPatient(userDetails models.User)error{
//
  // query := `INSERT INTO patients (name, date_of_birth, gender, address, phone, email, medical_history, registered_by)
  // VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id`
  // r.db.Raw(query,)
//   return nil
// }

func (r *ReceptionRepository)CheckUserAvailability(email string)(bool,error){
 
  var count uint
  query:=`select count(*) from users where email = $1`
  dbRes:= r.db.Raw(query,email).Scan(&count)
  if dbRes.Error!= nil{
    return false,dbRes.Error
  }
  if count > 0{
    return true,nil
  }

  return false ,nil

}

