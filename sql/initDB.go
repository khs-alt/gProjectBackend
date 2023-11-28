package sql

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

type App struct {
	DB   *sql.DB
	once sync.Once
}

var appInstance *App

func GetAppInstance() *App {
	if appInstance == nil {
		appInstance = &App{}
	}
	return appInstance
}

func (app *App) InitDB() {
	app.once.Do(func() {
		// dsn := "admin:QwR2]lPhV~4x^bx>E@/google_project"
		dsn := "root:1234@/google_project"
		db, err := sql.Open("mysql", dsn)
		if err != nil {
			log.Println(err)
			panic(err)
		}
		app.DB = db
	})
}

func (app *App) CloseDB() {
	if app.DB != nil {
		app.DB.Close()
	}
}

func SetDB() *App {
	app := GetAppInstance()
	app.InitDB()
	// If we use db again, you don't need to close the db
	// However you need to close rows whenever you call Query
	// defer app.CloseDB()
	return app
}

func DeleteDBTablbe() {
	app := SetDB()

	// 테이블 삭제
	tables := []string{
		"image_scoring",
		"video_scoring",
		"image_tag",
		"tag",
		"testcode",
		"image_testcode",
		"user",
		"user_testcode_info",
		"user_image_testcode_info",
		"video",
		"image",
		"image_patch",
	}
	for _, table := range tables {
		dropTableSQL := "DROP TABLE IF EXISTS " + table
		_, err := app.DB.Exec(dropTableSQL)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("테이블 %s 삭제 완료\n", table)
	}

}

func CreateDBTalbe() {
	app := SetDB()

	createTables := []string{
		`CREATE TABLE image_scoring (
            uuid BINARY(16) PRIMARY KEY,
            user_id VARCHAR(10) NOT NULL,
            image_id INT NOT NULL,
            patch_score VARCHAR(100) NOT NULL,
            time DATETIME NOT NULL
        )`,
		`CREATE TABLE video_scoring (
            uuid BINARY(16) PRIMARY KEY,
            user_id VARCHAR(20) NOT NULL,
            video_id INT NOT NULL,
            user_score INT NOT NULL,
            time DATETIME NOT NULL
        )`,
		`CREATE TABLE tag (
            uuid BINARY(16) PRIMARY KEY,
            tag VARCHAR(255) NOT NULL
        )`,
		`CREATE TABLE image_tag (
            uuid BINARY(16) PRIMARY KEY,
            tag VARCHAR(255) NOT NULL
        )`,
		`CREATE TABLE testcode (
            uuid BINARY(16) PRIMARY KEY,
            test_code VARCHAR(255) NOT NULL,
			tags VARCHAR(1000) NOT NULL,
            video_list VARCHAR(4000) NOT NULL
        )`,
		`CREATE TABLE image_testcode (
            uuid BINARY(16) PRIMARY KEY,
            test_code VARCHAR(255) NOT NULL,
			tags VARCHAR(1000) NOT NULL,
            image_list VARCHAR(4000) NOT NULL
        )`,
		`CREATE TABLE user (
            uuid BINARY(16),
            id VARCHAR(10),
            password VARCHAR(10) NOT NULL,
            PRIMARY KEY (uuid, id)
        )`,
		`CREATE TABLE user_testcode_info (
            uuid BINARY(16) PRIMARY KEY,
            user_id VARCHAR(10) NOT NULL,
            test_code VARCHAR(255) NOT NULL,
            current_page INT NOT NULL,
			time DATETIME NOT NULL
        )`,
		`CREATE TABLE user_image_testcode_info (
            uuid BINARY(16) PRIMARY KEY,
            user_id VARCHAR(10) NOT NULL,
            test_code VARCHAR(255) NOT NULL,
            current_page INT NOT NULL,
			time DATETIME NOT NULL
        )`,
		`CREATE TABLE video (
            uuid BINARY(16) PRIMARY KEY,
			original_video_name VARCHAR(255) NOT NULL,
            original_video VARCHAR(255) NOT NULL,
			original_video_fps float NOT NULL,
			artifact_video_name VARCHAR(255) NOT NULL,
            artifact_video VARCHAR(255) NOT NULL,
			artifact_video_fps float NOT NULL,
            tag VARCHAR(255) NOT NULL
        )`,
		`CREATE TABLE image (
            uuid BINARY(16) PRIMARY KEY,
            original_image_name VARCHAR(255) NOT NULL,
			original_image VARCHAR(255) NOT NULL,
            artifact_image_name VARCHAR(255) NOT NULL,
			artifact_image VARCHAR(255) NOT NULL,
			diff_image_name VARCHAR(255) NOT NULL,
			diff_image VARCHAR(255) NOT NULL,
            tag VARCHAR(255) NOT NULL
        )`,
		`CREATE TABLE image_patch (
            uuid BINARY(16),
            image_uuid BINARY(16),
            patch_id INT NOT NULL,
            auto_check_is_artifact BOOL NOT NULL,
            PRIMARY KEY (uuid, image_uuid)
        )`,
	}
	for _, createTableSQL := range createTables {
		_, err := app.DB.Exec(createTableSQL)
		if err != nil {
			log.Println(err)
		}
	}

	fmt.Println("테이블 생성 완료")
}
