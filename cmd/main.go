package main

import (
	"fmt"
	"github/MohdAhzan/internPortalAPP/pkg/config"
	"github/MohdAhzan/internPortalAPP/pkg/db"
	"log"

	"github.com/gin-gonic/gin"
)

func main(){
  
  cfg,err:=config.LoadConfig()
  if err!=nil{
    log.Fatal("Error loading Configs")
  }
  DB,err:=db.ConnectDB(cfg)
  if err!=nil{
    log.Fatal("Error connecting db",err)
  }

  testQuery:=fmt.Sprintf("select * from %s",cfg.DBName)
  res:=DB.Raw(testQuery)
  if res.Error !=nil{
    fmt.Print("error fetching database info")
  }
    fmt.Println(res)

  engine:=gin.New()
  logger:=gin.Logger()
  engine.Use(logger)
  err=engine.Run(cfg.Port)
  if err!=nil{
    log.Fatal("server failed to start at ",cfg.Port)
  }





  
}
