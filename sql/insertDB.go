package sql

import (
	"GoogleProjectBackend/util"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
)

func InsertTestCodeId(uuid uuid.UUID, testCode string, tags []string, videoList []string) error {
	app := SetDB()
	tagsCSV := util.MakeStringListtoCSV(tags)
	videoCSV := util.MakeStringListtoCSV(videoList)
	insertQuery := "INSERT INTO testcode (uuid, test_code, tags, video_list) VALUES (UUID_TO_BIN(?),?,?,?)"
	_, err := app.DB.Exec(insertQuery, uuid, testCode, tagsCSV, videoCSV)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("Tag Id insert")
	return nil
}

func InsertVideoId(uuid uuid.UUID, originalVideoName string, orgin string, artifactVideoName string, arti string, tag string) error {
	app := SetDB()

	insertQuery := "INSERT INTO video (uuid, original_video_name, original_video, artifact_video_name, artifact_video, tag) VALUES (UUID_TO_BIN(?), ?, ?, ?, ?, ?)"
	_, err := app.DB.Exec(insertQuery, uuid, originalVideoName, orgin, artifactVideoName, arti, tag)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("Video Id insert")
	return nil
}

func IsUserIdExist(id string, password string) string {
	app := SetDB()

	var resultID string
	var resultPassword string
	app.InitDB()

	query := "SELECT id, password FROM user WHERE id = ? AND password = ?"
	err := app.DB.QueryRow(query, id, password).Scan(&resultID, &resultPassword)
	fmt.Println(err)
	if err == sql.ErrNoRows {
		fmt.Println("No")
		return "No"
	} else if err != nil {
		fmt.Print("Login error")
		fmt.Println(err)
		return "error"
	} else {
		fmt.Println("Yes")
		return "Yes"
	}
}

func InsertUserIdAndPassword(uuid uuid.UUID, id string, ps string) string {
	app := SetDB()

	query := "SELECT COUNT(*) FROM user WHERE id = ?;"
	var count int
	err := app.DB.QueryRow(query, id).Scan(&count)
	if err != nil {
		panic(err)
	}

	if count > 0 {
		fmt.Println("Id is exist")
		return "ID is exist"
	} else {
		insertQuery := "INSERT INTO user (uuid, id, password) VALUES(UUID_TO_BIN(?), ?, ?)"
		_, err = app.DB.Exec(insertQuery, uuid, id, ps)
		if err != nil {
			panic(err)
		}
		fmt.Println("Insert success")
		return "signup success"
	}
}

func InsertUserTestInfo(uuid uuid.UUID, userId string, testCode string, currentPage int) {
	app := SetDB()

	insertQuert := "INSERT INTO user_testcode_info (uuid, user_id, test_code, current_page, time) VALUES (UUID_TO_BIN(?), ?, ?, ?, NOW())"
	_, err := app.DB.Exec(insertQuert, uuid, userId, testCode, currentPage)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("TestCode info inserted")
}

func InsertUserVideoScoringInfo(uuid uuid.UUID, userId string, videoId int, userScore int) {
	app := SetDB()

	insertQuery := "INSERT INTO video_scoring (uuid, user_id, video_id, user_score, time) VALUES (UUID_TO_BIN(?), ?, ?, ?, NOW())"
	_, err := app.DB.Exec(insertQuery, uuid, userId, videoId, userScore)
	if err != nil {
		fmt.Println(err)
	}
}

func InsertTagData(uuid uuid.UUID, tag string) {
	app := SetDB()

	insertQuery := "INSERT INTO tag (uuid, tag) VALUES (UUID_TO_BIN(?),?)"
	_, err := app.DB.Exec(insertQuery, uuid, tag)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("INSERT SUCCESS")
	}
}
