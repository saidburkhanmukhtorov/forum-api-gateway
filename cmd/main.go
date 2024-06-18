package main

import (
	"github.com/Forum-service/Forum-api-gateway/api"
	"github.com/Forum-service/Forum-api-gateway/config"
)

func main() {
	cfg := config.Load()
	r := api.NewEngine()
	if err := r.Run(cfg.HTTPPort); err != nil {
		panic(err)
	}
}
