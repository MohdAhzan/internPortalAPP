package config

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

type Config struct{
  DBHost     string `mapstructure:"DB_HOST"`
  DBName     string `mapstructure:"DB_NAME"`
  DBUser     string `mapstructure:"DB_USER"`
  DBPassword string `mapstructure:"DB_PASSWORD"`
  DBPort     string `mapstructure:"DB_PORT"`
  Port     string `mapstructure:"PORT"`
  ReceptionistSecret  string `mapstructure:"RECEPTIONIST_SECRET"`
  ReceptionistRefreshSecret  string `mapstructure:"RECEPTIONIST_REFRESH_SECRET"`
  DoctorSecret  string `mapstructure:"DOCTOR_SECRET"`
  DoctorRefreshSecret  string `mapstructure:"DOCTOR_REFRESH_SECRET"`
  AdminSecret  string `mapstructure:"ADMIN_SECRET"`
  AdminRefreshSecret  string `mapstructure:"ADMIN_REFRESH_SECRET"`
  AdminName string `mapstructure:"ADMIN_NAME"`
  AdminEmail string `mapstructure:"ADMIN_EMAIL"`
  AdminPassword string `mapstructure:"ADMIN_PASSWORD"`

}

var config Config

var envs = []string{
  "DB_HOST", "DB_NAME", "DB_USER", "DB_PORT", "DB_PASSWORD" ,"PORT","USER_SECRET","ADMIN_NAME","ADMIN_EMAIL","ADMIN_PASSWORD","ADMIN_SECRET","ADMIN_REFRESH_SECRET","RECEPTIONIST_SECRET","RECEPTIONIST_REFRESH_SECRET","DOCTER_SECRET","DOCTOR_REFRESH_SECRET"}

func LoadConfig() (Config, error) {

  viper.AddConfigPath(".")
  viper.SetConfigFile(".env")

  if err := viper.ReadInConfig(); err != nil{
    return config, fmt.Errorf("error reading config file: %v", err)
  }
    

  for _, env := range envs {
    if err := viper.BindEnv(env); err != nil {
      return config, err
    }
  }

  if err := viper.Unmarshal(&config); err != nil {

    return config, err
  }
  if err := validator.New().Struct(&config); err != nil {
    return config, err
  }

  // fmt.Println("env in LOad Config",config)

  return config, nil

}


