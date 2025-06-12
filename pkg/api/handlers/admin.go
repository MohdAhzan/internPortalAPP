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

func (u *AdminHandler)AdminLogin(c *gin.Context){
  var adminLogin models.AdminLogin
  err:=c.BindJSON(&adminLogin)
  if err!=nil{
    errRes:=response.ClientResponse(http.StatusBadRequest,"fields provided are in wrong format",nil,err.Error())
    c.JSON(http.StatusBadRequest,errRes)  
    return
  }
  resp,err:=u.usecase.AdminLogin(adminLogin)
  if err!=nil{
    errRes:=response.ClientResponse(http.StatusInternalServerError,"failed to signup admin",nil,err.Error())
    c.JSON(http.StatusBadRequest,errRes)  
    return
  }
  successRes:=response.ClientResponse(http.StatusOK,"successfully signedup admin",resp,nil)
  c.JSON(http.StatusOK,successRes)
}



func (u *AdminHandler)UserSignup(c *gin.Context){
  var userModel models.UserSignup
  err:=c.BindJSON(&userModel)
  if err!=nil{
    errRes:=response.ClientResponse(http.StatusBadRequest,"fields provided are in wrong format",nil,err.Error())
    c.JSON(http.StatusBadRequest,errRes)  
    return
  }
  resp,err:=u.usecase.UserSignup(userModel)
  if err!=nil{
    errRes:=response.ClientResponse(http.StatusInternalServerError,"failed to signup user",nil,err.Error())
    c.JSON(http.StatusBadRequest,errRes)  
    return
  }
  successRes:=response.ClientResponse(http.StatusOK,"successfully signedup user",resp,nil)
  c.JSON(http.StatusOK,successRes)
}



