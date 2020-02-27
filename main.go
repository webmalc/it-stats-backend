//+build !test

package main

import (
	"github.com/webmalc/it-stats-backend/logger"
)

func main() {
	log := logger.NewLogger(true)
	router := NewCommandRouter(log)
	router.Run()
}
