//+build !test

package main

import (
	"github.com/webmalc/it-stats-backend/apps/languages"
	"github.com/webmalc/it-stats-backend/common/admin"
	"github.com/webmalc/it-stats-backend/common/api"
	"github.com/webmalc/it-stats-backend/common/cmd"
	"github.com/webmalc/it-stats-backend/common/config"
	"github.com/webmalc/it-stats-backend/common/db"
	"github.com/webmalc/it-stats-backend/common/logger"
)

func main() {
	config.Setup()
	log := logger.NewLogger()
	conn := db.NewConnection()
	langApp := languages.NewApp(conn, log, &admin.CrudMixin{})
	adm := admin.NewAdmin(conn.DB)
	server := api.NewServer(adm, langApp)
	cmdRouter := cmd.NewCommandRouter(log, server)
	defer conn.Close()

	// register the language app
	conn.RegisterApp(langApp)
	cmdRouter.RegisterApp(langApp)

	cmdRouter.Run()
}
