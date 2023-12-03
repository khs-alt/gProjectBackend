package sql

import "database/sql"

func ExportImageData(testcode string) (*sql.Rows, error) {
	app := SetDB()
	rows, err := app.DB.Query("SELECT * FROM image_scoring WHERE test_code = ?", testcode)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	return rows, err
}

func ExportVideoData(testcode string) (*sql.Rows, error) {
	app := SetDB()
	rows, err := app.DB.Query("SELECT * FROM video_scoring WHERE test_code = ?", testcode)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	return rows, err
}
