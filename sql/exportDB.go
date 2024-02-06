package sql

import (
	"database/sql"
	"log"
)

func ExportImageData(testcode string) (*sql.Rows, error) {
	app := SetDB()
	insertQuery := `
				SELECT
					u.user_name, i.original_image_name, i.width, i.height, ims.patch_score, ims.image_testcode, ims.time
				FROM
					image_scoring AS ims
				JOIN
					user AS u ON u.uuid = ims.user_uuid
				JOIN
					image AS i ON i.uuid = ims.image_uuid
				WHERE
					ims.image_testcode = ?
				ORDER BY ims.time
				`
	rows, err := app.DB.Query(insertQuery, testcode)
	if err != nil {
		log.Println("ExportImageData error", err)
		return nil, err
	}

	return rows, err
}

func ExportVideoData(testcode string) (*sql.Rows, error) {
	app := SetDB()
	insertQuery := `
				SELECT 
					u.user_name, v.original_video_name, vs.user_score, vs.video_testcode, vs.time
				FROM 
					video_scoring AS vs
				JOIN 
					user AS u ON u.uuid = vs.user_uuid
				JOIN 
					video AS v ON v.uuid = vs.video_uuid
				WHERE 
					vs.video_testcode = ?
				ORDER BY vs.time
	`
	rows, err := app.DB.Query(insertQuery, testcode)
	if err != nil {
		log.Println("ExportVideoData error", err)
		return nil, err
	}

	return rows, err
}

func ExportVideoFrameData() (*sql.Rows, error) {
	app := SetDB()
	insertQuery := `
				SELECT
					v.original_video_name, v.artifact_video_name, v.diff_video_name, vst.video_frame, vst.time
				FROM
					video_selected_time AS vst
				JOIN
					video AS v ON v.uuid = vst.video_uuid
				ORDER BY vst.time
				`
	rows, err := app.DB.Query(insertQuery)
	if err != nil {
		log.Println("ExportVideoFrameData error", err)
		return nil, err
	}
	return rows, err
}
