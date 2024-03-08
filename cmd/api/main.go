package main

import (
	"os"
	"supertaltest/server"
	"supertaltest/server/router"

	"github.com/gin-gonic/gin"
	"github.com/subosito/gotenv"
)

func main() {
	if os.Getenv("ENV") != "PRODUCTION" {
		gotenv.Load()
	}

	r := gin.Default()

	apiHandler := server.NewApiHandler()
	router.NewRoute(r, apiHandler)
	r.Run(":" + os.Getenv("PORT"))
}
