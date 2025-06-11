package db

import (
	"database/sql"
	"fmt"
	"github/MohdAhzan/internPortalAPP/pkg/config"
	"github/MohdAhzan/internPortalAPP/pkg/utils/domain"
	"time"

	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB(cfg config.Config)(*gorm.DB,error){

   str:=fmt.Sprintf("user=%s password=%s host=%s",cfg.DBUser,cfg.DBPassword,cfg.DBHost)
   DB,err:=sql.Open("postgres",str)
    if err != nil {
		    return nil, fmt.Errorf("failed to open postgres DB: %w", err)
	  }

  rows, err := DB.Query("SELECT 1 FROM pg_database WHERE datname = $1",cfg.DBName)
	if err != nil {
    return nil,err
	}

	defer rows.Close()

	if rows.Next(){
		fmt.Println("db already exists")
	} else {
		// If the database does not exist, create it
		_, err = DB.Exec("CREATE DATABASE "+cfg.DBName)
		if err != nil {
      return nil,err
		}
	}

	dsn := fmt.Sprintf("host=%s user=%s dbname=%s port=%s password=%s", cfg.DBHost, cfg.DBUser, cfg.DBName, cfg.DBPort, cfg.DBPassword)
	db, dbErr := gorm.Open(postgres.Open(dsn), &gorm.Config{SkipDefaultTransaction: true})
  err=db.AutoMigrate(&domain.AdminDetails{},&domain.User{},&domain.DoctorDetails{}) 
  if err!=nil{
    return nil,err
  }
  
  err=CheckAndCreateAdmin(cfg,db)
  if err!=nil{
    return nil,fmt.Errorf("error creating admin %w",err)
  }
	return db, dbErr
}

func CheckAndCreateAdmin(cfg config.Config, db *gorm.DB)error {
	var count int64
	db.Model(&domain.AdminDetails{}).Count(&count)
	if count == 0 {
		password := cfg.AdminPassword
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		admin := domain.AdminDetails{
			Name:     cfg.AdminName,
			Email:    cfg.AdminEmail,
			Password: string(hashedPassword),
      CreatedAt: time.Now(), 
      UpdatedAt: time.Now(),
		}
		db.Create(&admin)
	}
return nil
}

