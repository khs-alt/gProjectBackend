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
		dsn := "admin:QwR2]lPhV~4x^bx>E@/google_project"
		//dsn := "root:1234@/google_project"
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

func DeleteImageDBTablbe() {
	app := SetDB()

	// 테이블 삭제
	tables := []string{
		"image_scoring",
		"image_tag",
		"image_testcode",
		"image",
		"image_tag_link",
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

func DeleteDBTable() {
	app := SetDB()

	// 테이블 삭제
	tables := []string{
		"video_tag",
		"image_tag",
		"user_testcode_info",
		// "image_scoring",
		// "video_scoring",
		//"user",
		"video",
		"image",
		"video_testcode",
		"image_testcode",
		"video_tag_link",
		"image_tag_link",
		"video_selected_time",
	}
	for _, table := range tables {
		app.DB.Exec("SET FOREIGN_KEY_CHECKS = 0")
		dropTableSQL := "DROP TABLE IF EXISTS " + table
		_, err := app.DB.Exec(dropTableSQL)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("테이블 %s 삭제 완료\n", table)
		app.DB.Exec("SET FOREIGN_KEY_CHECKS = 1")
	}

}

func CreateDBTable() {
	app := SetDB()

	createTables := []string{
		`CREATE TABLE video (
            uuid binary(16)  NOT NULL,
            original_video_name varchar(255),
            artifact_video_name varchar(255),
            diff_video_name varchar(255),
            video_frame float NOT NULL,
            video_index INT AUTO_INCREMENT PRIMARY KEY,
            width int,
            height int
        )`,
		`CREATE TABLE video_tag (
            uuid binary(16) PRIMARY KEY NOT NULL,
            tag varchar(255) UNIQUE NOT NULL
        )`,
		`CREATE TABLE image_tag (
            uuid binary(16) PRIMARY KEY NOT NULL,
            tag varchar(255) UNIQUE NOT NULL
        )`,
		`CREATE TABLE image (
            uuid binary(16) UNIQUE NOT NULL,
            original_image_name varchar(255),
            artifact_image_name varchar(255),
			diff_image_name varchar(255),
			image_index int AUTO_INCREMENT PRIMARY KEY,
			video_index int,
            width int,
            height int
        )`,
		`CREATE TABLE user (
            uuid binary(16) PRIMARY KEY NOT NULL,
            user_name varchar(255) NOT NULL,
            user_password varchar(255) NOT NULL
        )`,
		`CREATE TABLE user_testcode_info (
            uuid binary(16) PRIMARY KEY NOT NULL,
            user_uuid binary(16) NOT NULL,
            test_code varchar(255) NOT NULL,
            last_page int NOT NULL,
            time datetime,
            is_video boolean NOT NULL
        )`,
		`CREATE TABLE video_testcode (
            uuid binary(16) PRIMARY KEY NOT NULL,
            video_tag varchar(255) NOT NULL,
            video_testcode varchar(255) NOT NULL,
            description varchar(1000)
        )`,
		`CREATE TABLE video_scoring (
            uuid binary(16) PRIMARY KEY NOT NULL,
            user_uuid binary(16) NOT NULL,
            video_uuid binary(16) NOT NULL,
            user_score INT NOT NULL,
            video_testcode varchar(255) NOT NULL,
            time datetime,
			unique key user_video_testcode_unique(user_uuid, video_uuid, video_testcode)
        )`,
		`CREATE TABLE image_scoring (
            uuid binary(16) PRIMARY KEY NOT NULL,
            user_uuid binary(16),
            image_uuid binary(16),
            patch_score varchar(2000) NOT NULL,
            image_testcode varchar(255) NOT NULL,
            time datetime,
			unique key user_image_testcode_unique(user_uuid, image_uuid, image_testcode)
        )`,
		`CREATE TABLE image_testcode (
            uuid binary(16) PRIMARY KEY NOT NULL,
            image_tag varchar(255) NOT NULL,
            image_testcode varchar(255) NOT NULL,
            description varchar(1000)
        )`,
		`CREATE TABLE video_tag_link (
            video_uuid binary(16) NOT NULL,
            tag_uuid binary(16) NOT NULL,
            PRIMARY KEY (video_uuid, tag_uuid)
        )`,
		`CREATE TABLE image_tag_link (
            image_uuid binary(16) NOT NULL,
            tag_uuid binary(16) NOT NULL,
            PRIMARY KEY (image_uuid, tag_uuid)
        )`,
		`CREATE TABLE video_selected_time (
			uuid binary(16) PRIMARY KEY NOT NULL,
			video_uuid binary(16) NOT NULL,
			video_frame varchar(255) NOT NULL,
			time varchar(255) NOT NULL
		)`,
	}

	for _, createTableSQL := range createTables {
		_, err := app.DB.Exec(createTableSQL)
		if err != nil {
			log.Println("Error creating table: ", err)
		}
	}

	fmt.Println("All tables created successfully")
}

func CreateImageDBTalbe() {
	app := SetDB()

	createTables := []string{
		`CREATE TABLE image_tag (
            uuid binary(16) PRIMARY KEY NOT NULL,
            tag varchar(255) NOT NULL
        )`,
		`CREATE TABLE image (
            uuid binary(16) UNIQUE NOT NULL,
            original_image_name varchar(255),
            artifact_image_name varchar(255),
			diff_image_name varchar(255),
			image_index int AUTO_INCREMENT PRIMARY KEY,
			video_index int,
            width int,
            height int
        )`,
		`CREATE TABLE image_scoring (
            uuid binary(16) PRIMARY KEY NOT NULL,
            user_uuid binary(16),
            image_uuid binary(16),
            patch_score varchar(2000) NOT NULL,
            image_testcode varchar(255) NOT NULL,
            time datetime,
			unique key user_image_testcode_unique(user_uuid, image_uuid, image_testcode)
        )`,
		`CREATE TABLE image_testcode (
            uuid binary(16) PRIMARY KEY NOT NULL,
            image_tag varchar(255) NOT NULL,
            image_testcode varchar(255) NOT NULL,
            description varchar(1000)
        )`,
		`CREATE TABLE image_tag_link (
            image_uuid binary(16) NOT NULL,
            tag_uuid binary(16) NOT NULL,
            PRIMARY KEY (image_uuid, tag_uuid)
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
