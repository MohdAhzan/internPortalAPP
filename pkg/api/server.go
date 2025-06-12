package server

import (
	"github/MohdAhzan/internPortalAPP/pkg/api/handlers"
	"github/MohdAhzan/internPortalAPP/pkg/api/middleware"
	"github/MohdAhzan/internPortalAPP/pkg/config"
	"github/MohdAhzan/internPortalAPP/pkg/routes"
	"log"

	"github.com/gin-gonic/gin"
)


type ServerHTTP struct{

  engine *gin.Engine
}

func NewServeHTTP(adminHandler *handlers.AdminHandler,cfg config.Config,adminAuth middleware.AdminAuth,receptionistAuth middleware.ReceptionistAuth,doctorAuth middleware.DoctorAuth) (*ServerHTTP){
  newEngine:=gin.New()
  logger:=gin.Logger()
  newEngine.Use(logger)
  newEngine.Use(gin.Recovery())


  routes.ReceptionistRoutes(newEngine.Group("/receptionist"),gin.HandlerFunc(receptionistAuth))
  routes.AdminRoutes(newEngine.Group("/admin"),gin.HandlerFunc(adminAuth),adminHandler)
  routes.DoctorRoutes(newEngine.Group("/doctor"),gin.HandlerFunc(doctorAuth))

  return &ServerHTTP{
    engine: newEngine,
  }

}



func (s *ServerHTTP)Start(cfg config.Config){
  err:=s.engine.Run(cfg.Port)
  if err!=nil{
    log.Fatal("server failed to start at",err)
  }
}

