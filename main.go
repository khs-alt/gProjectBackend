package main

import (
	"GoogleProjectBackend/app"
	"GoogleProjectBackend/sql"
	"GoogleProjectBackend/util"
	"io"
	"log"
	"net/http"
	"os"
)

var logFile *os.File

func InitLogFile() {
	logFile, err := os.OpenFile("logfile.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	//defer logFile.Close()
	//log.SetOutput(logFile)

	// 화면과 파일에 동시에 로그 찍기
	multiWriter := io.MultiWriter(logFile, os.Stdout)
	log.SetOutput(multiWriter)

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetPrefix("INFO: ")
	log.Println("Logging")
}

// reset DB and system
func Init() {
	sql.DeleteDBTablbe()
	sql.CreateDBTalbe()
	util.DeleteAllFilesInFolder("originalVideos")
	util.DeleteAllFilesInFolder("artifactVideos")
}

func main() {
	InitLogFile()
	router := app.SetupRouter()
	log.Fatal(http.ListenAndServe(":8000", router))
	defer logFile.Close()

}
