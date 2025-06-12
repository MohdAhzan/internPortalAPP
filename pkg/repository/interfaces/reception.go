package interfaces

type ReceptionRepository interface{
  // AddPatient(userDetails domain.Patient)error
  CheckUserAvailability(email string)(bool,error)

}
