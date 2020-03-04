//+build !test

package main

import (
	"net/http"

	"github.com/webmalc/it-stats-backend/apps/languages"
	"github.com/webmalc/it-stats-backend/common/admin"
	"github.com/webmalc/it-stats-backend/common/cmd"
	"github.com/webmalc/it-stats-backend/common/config"
	"github.com/webmalc/it-stats-backend/common/db"
	"github.com/webmalc/it-stats-backend/common/logger"
)

func main() {
	config.Setup()
	log := logger.NewLogger()
	router := cmd.NewCommandRouter(log)
	conn := db.NewConnection()
	mux := http.NewServeMux()
	adm := admin.NewAdmin(conn.DB, mux)
	defer conn.Close()

	// register the language app
	langApp := languages.NewApp(conn)
	conn.RegisterApp(langApp)
	router.RegisterApp(langApp)
	adm.RegisterApp(langApp)

	// Or http.ListenAndServe(":9000", mux)
	router.Run()
}
