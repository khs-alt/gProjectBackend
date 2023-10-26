package main

// import (
// 	"io"
// 	"log"
// 	"os"
// )

// func InitLogFile() {
// 	logFile, err := os.OpenFile("logfile.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer logFile.Close()
// 	//log.SetOutput(logFile)

// 	// 화면과 파일에 동시에 로그 찍기
// 	multiWriter := io.MultiWriter(logFile, os.Stdout)
// 	log.SetOutput(multiWriter)

// 	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
// 	log.SetPrefix("INFO: ")
// 	//log.Println("Logging")
// }
