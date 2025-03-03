package main

import (

	"github.com/hailsayan/golang-api/api"
	"github.com/hailsayan/golang-api/config"
	"github.com/hailsayan/golang-api/data/cache"
	"github.com/hailsayan/golang-api/data/db"
	"github.com/hailsayan/golang-api/pkg/logging"
)

// @securityDefinitions.apikey AuthBearer
// @in header
// @name Authorization
func main() {

	cfg := config.GetConfig()
	logger := logging.NewLogger(cfg)

	err := cache.InitRedis(cfg)
	defer cache.CloseRedis()
	if err != nil {
		logger.Fatal(logging.Redis, logging.Startup, err.Error(), nil)
	}

	err = db.InitDb(cfg)
	defer db.CloseDb()
	if err != nil {
		logger.Fatal(logging.Postgres, logging.Startup, err.Error(), nil)
	}

	api.InitServer(cfg)
}
