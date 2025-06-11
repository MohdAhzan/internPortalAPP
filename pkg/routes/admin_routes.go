package routes

import (
	"github/MohdAhzan/internPortalAPP/pkg/api/handlers"

	"github.com/gin-gonic/gin"
)
func AdminRoutes(engine *gin.RouterGroup,adminAuth gin.HandlerFunc,userHandler *handlers.UserHandler){
  engine.POST("/signup",userHandler.UserSignup)

}

