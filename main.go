package main

import (
	"backend/app"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(cors.Default())
	app.Routes((router))
	router.Run(":8000")
}
