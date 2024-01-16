package main

import (
	"backend/app"
	"io"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitLogFile() {
	logFile, err := os.OpenFile("logfile.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	multiWriter := io.MultiWriter(logFile, os.Stdout)
	log.SetOutput(multiWriter)

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetPrefix("INFO: ")
	log.Println("Logging")
}

func main() {

	InitLogFile()
	router := gin.Default()
	router.Use(cors.Default())
	app.Routes((router))
	router.Run(":8000")
}
