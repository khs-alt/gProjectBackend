package sql

import (
	"database/sql"
	"fmt"
)

func ExportVideoData(testcode string) string {
	app := SetDB()
	fmt.Println(testcode)
	var data string
	insertQuery := "SELECT image_list FROM image_testcode WHERE test_code = ?"
	err := app.DB.QueryRow(insertQuery).Scan(&data)
	if err != nil {
		panic(err)
	}

	return data
}

func ExportImageData(testcode string) (*sql.Rows, error) {
	app := SetDB()
	fmt.Println(testcode)
	rows, err := app.DB.Query("SELECT * FROM video_scoring")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	return rows, err
}
