package sql

import (
	"fmt"
)

func DeleteTagData(tag string) error {
	app := SetDB()

	insertQuery := "DELETE FROM tag WHERE tag = ?"
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

	insertQuery := "DELETE FROM image_tag WHERE tag = ?"
	_, err := app.DB.Exec(insertQuery, tag)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("Delet Success")
	return nil

}
