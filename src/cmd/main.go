package main

import (
	"log"

	"github.com/hailsayan/golang-api/api"
	"github.com/hailsayan/golang-api/config"
	"github.com/hailsayan/golang-api/data/cache"
	"github.com/hailsayan/golang-api/data/db"
)

// @securityDefinitions.apikey AuthBearer
// @in header
// @name Authorization
func main() {
	cfg := config.GetConfig()

	err := cache.InitRedis(cfg)
	defer cache.CloseRedis()
	if err != nil {
		log.Fatal(err)
	}

	err = db.InitDb(cfg)
	defer db.CloseDb()
	if err != nil {
		log.Fatal(err)
	}

	api.InitServer(cfg)
}
