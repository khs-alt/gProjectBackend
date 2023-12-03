package sql

import (
	"database/sql"
	"fmt"
)

func ExportImageData(testcode string) (*sql.Rows, error) {
	app := SetDB()
	fmt.Println(testcode)
	rows, err := app.DB.Query("SELECT * FROM image_scoring")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	return rows, err
}

func ExportVideoData(testcode string) (*sql.Rows, error) {
	app := SetDB()
	fmt.Println(testcode)
	rows, err := app.DB.Query("SELECT * FROM video_scoring")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	return rows, err
}
