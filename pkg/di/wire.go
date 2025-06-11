//go:build wireinject
// +build wireinject

package di

import (
	server "github/MohdAhzan/internPortalAPP/pkg/api"
	"github/MohdAhzan/internPortalAPP/pkg/api/handlers"
	"github/MohdAhzan/internPortalAPP/pkg/api/middleware"
	"github/MohdAhzan/internPortalAPP/pkg/config"
	"github/MohdAhzan/internPortalAPP/pkg/db"
	"github/MohdAhzan/internPortalAPP/pkg/helper"
	"github/MohdAhzan/internPortalAPP/pkg/repository"
	"github/MohdAhzan/internPortalAPP/pkg/usecase"

	"github.com/google/wire"
)


func InitializeAPI(cfg config.Config) (*server.ServerHTTP,error) {
  wire.Build(

    db.ConnectDB,
    server.NewServeHTTP,

    NewAdminAuth,
    NewReceptionistAuth,
    NewDoctorAuth,

    handlers.NewUserHandler,
    usecase.NewAdminUsecase,
    repository.NewReceptionRepository,
    repository.NewAdminRepository,

    helper.NewHelper,
    
    )
  return &server.ServerHTTP{},nil
}


func NewAdminAuth(cfg config.Config)  middleware.AdminAuth{
  return middleware.AdminAuth(middleware.AdminAuthMiddleware(cfg))
}

func NewReceptionistAuth(cfg config.Config) middleware.ReceptionistAuth{
  return    middleware.ReceptionistAuth(middleware.ReceptionistAuthMiddleware(cfg))
}
func NewDoctorAuth(cfg config.Config) middleware.DoctorAuth{
  return    middleware.DoctorAuth(middleware.DoctorAuthMiddleware(cfg))
}
