package sql

import (
	"database/sql"
	"fmt"

	"github.com/joho/sqltocsv"
)

func ExportImageData(testcode string) (*sql.Rows, error) {
	app := SetDB()
	fmt.Println(testcode)
	rows, err := app.DB.Query("SELECT user_id, image_id, patch_score, time FROM image_scoring ORDER BY time")
	if err != nil {
		panic(err)
	}

	err = sqltocsv.WriteFile("./image_user_labeling.csv", rows)
	if err != nil {
		fmt.Println(err)
	}

	rows, err = app.DB.Query("SELECT user_id, image_id, patch_score, time FROM image_scoring ORDER BY time")
	if err != nil {
		panic(err)
	}

	return rows, err
}

func ExportVideoData(testcode string) (*sql.Rows, error) {
	app := SetDB()
	fmt.Println(testcode)

	rows, err := app.DB.Query("SELECT user_id, video_id, user_score, time FROM video_scoring ORDER BY time")
	if err != nil {
		panic(err)
	}

	err = sqltocsv.WriteFile("./video_user_scoring.csv", rows)
	if err != nil {
		fmt.Println(err)
	}

	rows, err = app.DB.Query("SELECT user_id, video_id, user_score, time FROM video_scoring ORDER BY time")
	if err != nil {
		panic(err)
	}

	return rows, err
}
