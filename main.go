package main

import (
	"github.com/eXtern-OS/core9-common/db"
	"github.com/eXtern-OS/core9-common/utils"
	"github.com/eXtern-OS/core9-id/server"
	"github.com/gin-gonic/gin"
	"log"
)

type Config struct {
	MongoURI string `json:"mongo_uri"`
}

func main() {
	r := gin.Default()

	server.SetEndpoints(r)
	var c Config
	err := utils.ReadConfig(&c)
	if err != nil {
		log.Fatalln(err)
	}
	db.Init(c.MongoURI)
	err = r.Run(":8080")

}
