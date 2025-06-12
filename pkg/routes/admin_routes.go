package routes

import (
	"github/MohdAhzan/internPortalAPP/pkg/api/handlers"

	"github.com/gin-gonic/gin"
)
func AdminRoutes(engine *gin.RouterGroup,adminAuth gin.HandlerFunc,adminHandler *handlers.AdminHandler){
  
  engine.POST("/login",adminHandler.AdminLogin)

  engine.Use(adminAuth)
  engine.POST("/signup",adminHandler.UserSignup)
}

