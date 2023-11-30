package sql

import "database/sql"

func ExportImageData() (*sql.Rows, error) {
	app := SetDB()
	rows, err := app.DB.Query("SELECT * FROM image_scoring")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	return rows, err
}

func ExportVideoData() (*sql.Rows, error) {
	app := SetDB()
	rows, err := app.DB.Query("SELECT * FROM video_scoring")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	return rows, err
}
