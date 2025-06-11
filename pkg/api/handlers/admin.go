package handlers

import (
	"github/MohdAhzan/internPortalAPP/pkg/usecase/interfaces"
	response "github/MohdAhzan/internPortalAPP/pkg/utils/Response"
	"github/MohdAhzan/internPortalAPP/pkg/utils/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AdminHandler struct{
  usecase  interfaces.AdminUsecase
}

func NewUserHandler(u interfaces.AdminUsecase )*AdminHandler{
  return &AdminHandler{
    usecase: u,
  }
}


func (u *AdminHandler)UserSignup(c *gin.Context){
  var userLogin models.UserSignup
  err:=c.BindJSON(&userLogin)
  if err!=nil{
    errRes:=response.ClientResponse(http.StatusBadRequest,"fields provided are in wrong format",nil,err)
    c.JSON(http.StatusBadRequest,errRes)  
    return
  }
  resp,err:=u.usecase.UserSignup(userLogin)
  if err!=nil{
    errRes:=response.ClientResponse(http.StatusInternalServerError,"failed to signup user",nil,err)
    c.JSON(http.StatusBadRequest,errRes)  
    return
  }
  successRes:=response.ClientResponse(http.StatusOK,"successfully signedup user",resp,nil)
  c.JSON(http.StatusOK,successRes)
}




