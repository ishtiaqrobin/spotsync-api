package main

import (
	"github.com/ishtiaqrobin/spotsync-api/internal/config"
	"github.com/ishtiaqrobin/spotsync-api/internal/server"
)

func main() {
	cfg := config.LoadEnv()
	db := config.ConnectDatabase(cfg)
	server.Start(db, cfg)
}
