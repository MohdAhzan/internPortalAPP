package main

import (
	"github/MohdAhzan/internPortalAPP/pkg/config"
	"github/MohdAhzan/internPortalAPP/pkg/di"
	"log"
)

func main(){
  
  cfg,err:=config.LoadConfig()
  if err!=nil{
    log.Fatal("Error loading Configs")
  }


  server,err:=di.InitializeAPI(cfg)
  if err!=nil{
    log.Fatal(err,"FAILED TO START THE SERVER")
  }else{
   server.Start(cfg)
  }


  
}
