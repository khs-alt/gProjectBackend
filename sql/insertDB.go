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

// 비디오 정보를 넣는 sql 함수
// 주요 파라미터로는 오리지널 비디오 이름, 디고스팅된 비디오 이름, 태그가 있다.
// TODO: 프레임 정보를 넣어야 함.
func InsertVideoId(uuid uuid.UUID, originalVideoName string, orgin string, originalVideoFPS float32, artifactVideoName string, arti string, artifactVideoFPS float32, tag string) error {
	app := SetDB()

	insertQuery := "INSERT INTO video (uuid, original_video_name, original_video, original_video_fps, artifact_video_name, artifact_video, artifact_video_fps, tag) VALUES (UUID_TO_BIN(?), ?, ?, ?, ?, ?, ?, ?)"
	_, err := app.DB.Exec(insertQuery, uuid, originalVideoName, orgin, originalVideoFPS, artifactVideoName, arti, artifactVideoFPS, tag)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("Video Id insert")
	return nil
}

func IsUserIdExist(id string, password string) bool {
	app := SetDB()

	var resultID string
	var resultPassword string
	app.InitDB()

	query := "SELECT id, password FROM user WHERE id = ? AND password = ?"
	err := app.DB.QueryRow(query, id, password).Scan(&resultID, &resultPassword)
	fmt.Println(err)
	if err == sql.ErrNoRows {
		return false
	} else if err != nil {
		fmt.Print("Login error")
		panic(err)
	} else {
		return true
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
		return "No"
	} else {
		insertQuery := "INSERT INTO user (uuid, id, password) VALUES(UUID_TO_BIN(?), ?, ?)"
		_, err = app.DB.Exec(insertQuery, uuid, id, ps)
		if err != nil {
			panic(err)
		}
		fmt.Println("Insert success")
		return "Yes"
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
