package sql

import (
	"database/sql"
	"fmt"

	"github.com/joho/sqltocsv"
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
	fmt.Println("=====================================")
	rows, err := app.DB.Query("SELECT * FROM video_scoring")
	if err != nil {
		panic(err)
	}
	err = sqltocsv.WriteFile("./test.csv", rows)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	return rows, err
}
