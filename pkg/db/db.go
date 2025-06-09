package db

import (
	"database/sql"
	"fmt"
	"github/MohdAhzan/internPortalAPP/pkg/config"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB(cfg config.Config)(*gorm.DB,error){

   str:=fmt.Sprintf("user=%s password=%s host=%s",cfg.DBUser,cfg.DBPassword,cfg.DBHost)
   DB,err:=sql.Open("postgres",str)
    if err != nil {
		    return nil, err
	  }

  rows, err := DB.Query("SELECT 1 FROM pg_database WHERE datname = $1",cfg.DBName)
	if err != nil {
    return nil,err
	}

	defer rows.Close()

	if rows.Next() {
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


	return db, dbErr

}
