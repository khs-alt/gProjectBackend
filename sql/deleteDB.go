package sql

import (
	"fmt"
)

func DeleteTagData(tag string) error {
	app := SetDB()

	insertQuery := "DELETE FROM video_tag WHERE video_tag = ?"
	_, err := app.DB.Exec(insertQuery, tag)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("Delet Success")
	return nil

}

func DeleteImageTagData(tag string) error {
	app := SetDB()

	insertQuery := "DELETE FROM image_tag WHERE image_tag = ?"
	_, err := app.DB.Exec(insertQuery, tag)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("Delet Success")
	return nil

}
