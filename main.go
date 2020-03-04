//+build !test

package main

import (
	"github.com/webmalc/it-stats-backend/apps/languages"
	"github.com/webmalc/it-stats-backend/cmd"
	"github.com/webmalc/it-stats-backend/config"
	"github.com/webmalc/it-stats-backend/db"
	"github.com/webmalc/it-stats-backend/logger"
)

func main() {
	config.Setup()
	log := logger.NewLogger()
	router := cmd.NewCommandRouter(log)
	conn := db.NewConnection()
	defer conn.Close()

	// register the language app
	langApp := languages.NewApp(conn)
	conn.RegisterApp(langApp)
	router.RegisterApp(langApp)

	router.Run()
}
