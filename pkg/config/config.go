package config

import "fmt"

type Config struct{
  DBHost     string `mapstructure:"DB_HOST"`
  DBName     string `mapstructure:"DB_NAME"`
  DBUser     string `mapstructure:"DB_USER"`
  DBPassword string `mapstructure:"DB_PASSWORD"`
  DBPort     string `mapstructure:"DB_PORT"`
}


var envs = []string{
  "DB_HOST", "DB_NAME", "DB_USER", "DB_PORT", "DB_PASSWORD" , "USER_SECRET"}

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
  fmt.Println("env in LOad Config",config)

  return config, nil

}


