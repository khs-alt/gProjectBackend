package sql

import (
	"log"

	"github.com/google/uuid"
)

func DeleteTagData(tag string) error {
	app := SetDB()

	insertQuery := "DELETE FROM video_tag WHERE video_tag = ?"
	_, err := app.DB.Exec(insertQuery, tag)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil

}

func DeleteImageTagData(tag string) error {
	app := SetDB()

	insertQuery := "DELETE FROM image_tag WHERE image_tag = ?"
	_, err := app.DB.Exec(insertQuery, tag)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil

}

func DeleteImage(videoIndex int) error {
	app := SetDB()

	deleteQuery := `
				DELETE 
					i, itl
				FROM 
					image AS i
				JOIN 
					image_tag_link AS itl ON i.uuid = itl.image_uuid
				WHERE
					i.video_index = ?
				`
	_, err := app.DB.Exec(deleteQuery, videoIndex)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func DeleteVideoTime(videoIndex int) error {
	app := SetDB()
	insertQuery := "SELECT BIN_TO_UUID(uuid) FROM video WHERE video_index = ?"
	var videoUUID uuid.UUID
	err := app.DB.QueryRow(insertQuery, videoIndex).Scan(&videoUUID)
	if err != nil {
		log.Println(err)
		return err
	}

	deleteQuery := "DELETE FROM video_selected_time WHERE BIN_TO_UUID(video_uuid) = ?"
	_, err = app.DB.Exec(deleteQuery, videoUUID)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
