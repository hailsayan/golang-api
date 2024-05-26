package main

import (
	"github.com/hailsayan/golang-api/api"
	"github.com/hailsayan/golang-api/config"
	"github.com/hailsayan/golang-api/data/cache"
)

func main() {
	cfg := config.GetConfig()
	cache.InitRedis(cfg)
	defer cache.CloseRedis()
	api.InitServer(cfg)
}
