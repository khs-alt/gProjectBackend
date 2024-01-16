package sql

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

// insert video testcode
// done
func InsertVideoTestCode(uuid uuid.UUID, videoTag string, videoTestCode string, description string) error {
	app := SetDB()

	insertQuery := "INSERT INTO video_testcode (uuid, video_tag, video_testcode, description) VALUES (UUID_TO_BIN(?),?,?,?)"
	_, err := app.DB.Exec(insertQuery, uuid, videoTag, videoTestCode, description)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

// insert image testcode
// done
func InsertImageTestCode(uuid uuid.UUID, imageTag string, imageTestCode string, description string) error {
	app := SetDB()

	insertQuery := "INSERT INTO image_testcode (uuid, image_tag, image_testcode, description) VALUES (UUID_TO_BIN(?),?,?,?)"
	_, err := app.DB.Exec(insertQuery, uuid, imageTag, imageTestCode, description)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// 비디오 정보를 넣는 sql 함수
// 주요 파라미터로는 오리지널 비디오 이름, 디고스팅된 비디오 이름, 태그가 있다.
// finished
func InsertVideo(uuid uuid.UUID, originalVideoName string, artifactVideoName string, diffVideoName string, videoFrame float32, width int, height int) error {
	app := SetDB()
	insertQuery := "INSERT INTO video (uuid, original_video_name, artifact_video_name, diff_video_name, video_frame, width, height) VALUES (UUID_TO_BIN(?), ?, ?, ?, ?, ?, ?)"
	_, err := app.DB.Exec(insertQuery, uuid, originalVideoName, artifactVideoName, diffVideoName, videoFrame, width, height)
	if err != nil {
		log.Println(err)
		return err
	}
	fmt.Println("Video Id insert")
	return nil
}

// done
func InsertImage(uuid uuid.UUID, originalImageName string, artifactImageName string, diffImageName string, width int, height int) error {
	app := SetDB()

	insertQuery := "INSERT INTO image (uuid, original_image_name, artifact_image_name, diff_image_name, width, height) VALUES (UUID_TO_BIN(?), ?, ?, ?, ?, ?)"
	_, err := app.DB.Exec(insertQuery, uuid, originalImageName, artifactImageName, diffImageName, width, height)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("Image Id insert")
	return nil
}

// check user id and password
// TODO: check id wrong, password wrong and both wrong
// done
func IsUserIdExist(id string, password string) bool {
	app := SetDB()

	var resultID string
	var resultPassword string
	app.InitDB()

	query := "SELECT user_name, user_password FROM user WHERE user_name = ? AND user_password = ?"
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

// Insert user id and password
// done
func InsertUserIdAndPassword(uuid uuid.UUID, id string, ps string) string {
	app := SetDB()

	// if user is exist, return "No"
	insertQuery := `
    INSERT INTO user (uuid, user_name, user_password) 
    VALUES (UUID_TO_BIN(?), ?, ?)`

	_, err := app.DB.Exec(insertQuery, uuid, id, ps)
	if err != nil {
		// if errcode is 1062 then return "Id is exist"
		if mysqlErr, ok := err.(*mysql.MySQLError); ok && mysqlErr.Number == 1062 {
			return "Id is exist"
		} else {
			panic(err) // 다른 유형의 에러 처리
		}
	}

	return "Yes"
}

// done
func InsertUserTestInfo(newUUID uuid.UUID, userId string, testCode string, lastPage int, isVideo bool) {
	app := SetDB()

	var userUUID uuid.UUID
	userUUIDQuery := "SELECT BIN_TO_UUID(uuid) FROM user WHERE user_name = ?"

	err := app.DB.QueryRow(userUUIDQuery, userId).Scan(&userUUID)
	if err != nil {
		log.Printf("Error finding user_uuid: %v\n", err)
		return
	}

	insertQuery := "INSERT INTO user_testcode_info (uuid, user_uuid, test_code, last_page, time, is_video) VALUES (UUID_TO_BIN(?), UUID_TO_BIN(?), ?, ?, NOW(), ?)"
	_, err = app.DB.Exec(insertQuery, newUUID, userUUID, testCode, lastPage, isVideo)
	if err != nil {
		log.Printf("Error inserting into user_testcode_info: %v\n", err)
		return
	}

	fmt.Println("Insert user test info")
}

// done
// InsertUserImageScoringInfo 함수는 userId, imageTestcode, patchScore를 사용하여 image_scoring 테이블에 데이터를 삽입합니다.
func InsertUserImageScoringInfo(userId string, imageID int, imageTestcode string, patchScore string) {
	app := SetDB()
	// 사용자의 uuid를 찾습니다.
	var userUUID uuid.UUID
	userUUIDQuery := "SELECT BIN_TO_UUID(uuid) FROM user WHERE user_name = ?"
	err := app.DB.QueryRow(userUUIDQuery, userId).Scan(&userUUID)
	if err != nil {
		log.Printf("Error finding user_uuid: %v\n", err)
		return
	}

	// imageTestcode에 해당하는 image_uuid를 찾습니다.
	var imageUUID uuid.UUID
	imageUUIDQuery := `
    SELECT BIN_TO_UUID(i.uuid)
    FROM image i
    WHERE i.image_index = ?
	`

	err = app.DB.QueryRow(imageUUIDQuery, imageID).Scan(&imageUUID)
	if err != nil {
		log.Printf("Error finding image_uuid: %v\n", err)
		return
	}

	// image_scoring에 데이터를 삽입합니다.
	insertQuery := `
    INSERT INTO image_scoring (uuid, user_uuid, image_uuid, patch_score, image_testcode, time)
    VALUES (UUID_TO_BIN(?), UUID_TO_BIN(?), UUID_TO_BIN(?), ?, ?, NOW())`

	scoringUUID := uuid.New() // 새로운 UUID 생성
	_, err = app.DB.Exec(insertQuery, scoringUUID, userUUID, imageUUID, patchScore, imageTestcode)
	if err != nil {
		log.Printf("Error inserting into image_scoring: %v\n", err)
		return
	}
}

// done
func InsertUserVideoScoringInfo(userId string, videoID int, testCode string, score int) {
	app := SetDB()

	var userUUID uuid.UUID
	userUUIDQuery := "SELECT BIN_TO_UUID(uuid) FROM user WHERE user_name = ?"

	err := app.DB.QueryRow(userUUIDQuery, userId).Scan(&userUUID)
	if err != nil {
		log.Printf("Error finding user_uuid: %v\n", err)
		return
	}

	// videoID에 해당하는 video_uuid를 찾습니다.
	var videoUUID uuid.UUID
	videoUUIDQuery := `
    SELECT BIN_TO_UUID(v.uuid)
    FROM video v
    WHERE v.video_index = ?`

	err = app.DB.QueryRow(videoUUIDQuery, videoID).Scan(&videoUUID)
	if err != nil {
		log.Printf("Error finding video_uuid: %v\n", err)
		return
	}

	// video_scoring에 데이터를 삽입합니다.
	insertQuery := `
    INSERT INTO video_scoring (uuid, user_uuid, video_uuid, user_score, video_testcode, time)
    VALUES (UUID_TO_BIN(?), UUID_TO_BIN(?), UUID_TO_BIN(?), ?, ?, NOW())`

	scoringUUID := uuid.New() // 새로운 UUID 생성
	_, err = app.DB.Exec(insertQuery, scoringUUID, userUUID, videoUUID, score, testCode)
	if err != nil {
		log.Printf("Error inserting into video_scoring: %v\n", err)
		return
	}

}

// done
func InsertVideoTag(uuid uuid.UUID, tag string) {
	app := SetDB()

	insertQuery := "INSERT INTO video_tag (uuid, tag) VALUES (UUID_TO_BIN(?),?)"
	_, err := app.DB.Exec(insertQuery, uuid, tag)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("INSERT tag SUCCESS")
	}
}

// done
func InsertImageTag(uuid uuid.UUID, tag string) {
	app := SetDB()

	insertQuery := "INSERT INTO image_tag (uuid, tag) VALUES (UUID_TO_BIN(?),?)"
	_, err := app.DB.Exec(insertQuery, uuid, tag)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("INSERT tag SUCCESS")
	}
}

func InsertVideoTagLink(videoUUID uuid.UUID, tag string) error {
	app := SetDB()
	insertQuery := "SELECT BIN_TO_UUID(uuid) FROM video_tag WHERE tag = ?"
	var tagUUID uuid.UUID
	err := app.DB.QueryRow(insertQuery, tag).Scan(&tagUUID)
	if err != nil {
		fmt.Println(err)
		return err
	}

	insertQuery = "INSERT INTO video_tag_link (video_uuid, tag_uuid) VALUES (UUID_TO_BIN(?), UUID_TO_BIN(?))"
	_, err = app.DB.Exec(insertQuery, videoUUID, tagUUID)
	if err != nil {
		fmt.Println(err)
		return err
	} else {
		fmt.Println("INSERT video tag link SUCCESS")
		return nil
	}
}

func InsertImageTagLink(imageUUID uuid.UUID, tag string) error {
	app := SetDB()
	insertQuery := "SELECT BIN_TO_UUID(uuid) FROM image_tag WHERE tag = ?"
	var tagUUID uuid.UUID
	err := app.DB.QueryRow(insertQuery, tag).Scan(&tagUUID)
	if err != nil {
		fmt.Println(err)
		return err
	}

	insertQuery = "INSERT INTO image_tag_link (image_uuid, tag_uuid) VALUES (UUID_TO_BIN(?), UUID_TO_BIN(?))"
	_, err = app.DB.Exec(insertQuery, imageUUID, tagUUID)
	if err != nil {
		fmt.Println(err)
		return err
	} else {
		fmt.Println("INSERT image tag link SUCCESS")
		return nil
	}
}
