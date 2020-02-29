//+build !test

package main

import (
	"github.com/webmalc/it-stats-backend/cmd"
	"github.com/webmalc/it-stats-backend/config"
	"github.com/webmalc/it-stats-backend/logger"
)

func main() {
	config.Setup()
	log := logger.NewLogger()
	router := cmd.NewCommandRouter(log)
	router.Run()
}
