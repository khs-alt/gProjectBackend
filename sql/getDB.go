package sql

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/google/uuid"
)

func GetCurrentUserScore(userId string, videoId int) int {
	app := SetDB()

	insertQuery := "SELECT user_score FROM video_scoring WHERE user_id = ? AND video_id = ? ORDER BY time DESC LIMIT 1"
	var score int
	err := app.DB.QueryRow(insertQuery, userId, videoId).Scan(&score)
	if err != nil {
		if err == sql.ErrNoRows {
			// 결과가 없을 때 -1 반환
			return -1
		} else {
			// 다른 오류가 발생한 경우 로그를 출력
			fmt.Println(err)
		}
	}
	return score
}

func GetVideoAverageScore(video string) int {
	app := SetDB()
	n := len(video)
	fmt.Println(video)
	videoId := string(video[n-1])
	fmt.Println(videoId)
	insertQuery := "SELECT user_score FROM video_scoring WHERE video_id = ?"
	rows, err := app.DB.Query(insertQuery, videoId)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()
	var scoreList []int
	for rows.Next() {
		score := 0
		if err := rows.Scan(&score); err != nil {
			log.Fatal(err)
		}
		scoreList = append(scoreList, score)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	var sum int
	//평균 구하기
	for _, score := range scoreList {
		sum += score
	}
	fmt.Println("====================")
	averageScore := sum / len(scoreList)
	return averageScore
}

func GetUserCurrentPageAboutTestCode(userId string, testCode string) int {
	app := SetDB()
	maxCurrentPage := 0

	insertQuery := "SELECT current_page FROM user_testcode_info WHERE user_id = ? AND test_code = ? ORDER BY time DESC LIMIT 1"
	err := app.DB.QueryRow(insertQuery, userId, testCode).Scan(&maxCurrentPage)
	if err != nil {
		fmt.Println(err)
	}
	return maxCurrentPage
}

func GetFPSFromVideo(videoId string) (float32, float32) {
	app := SetDB()

	query := "SELECT original_video_fps, artifact_video_fps FROM video WHERE original_video = ?"
	var originalVideoFPS float32
	var artifactVideoFPS float32
	err := app.DB.QueryRow(query, videoId).Scan(&originalVideoFPS, &artifactVideoFPS)
	if err != nil {
		fmt.Println(err)
	}
	return originalVideoFPS, artifactVideoFPS
}

func GetTestcodeExist(testCode string) bool {
	app := SetDB()

	query := "SELECT COUNT(*) FROM testcode WHERE test_code = ?"
	var count int
	err := app.DB.QueryRow(query, testCode).Scan(&count)
	if err != nil {
		panic(err)
	}
	if count > 0 {
		return true
	}
	return false
}

func GetTestCodeCount() (int, error) {
	app := SetDB()

	var count int
	err := app.DB.QueryRow("SELECT COUNT(*) FROM testcode").Scan(&count)
	if err != nil {
		panic(err)
	}
	return count, nil
}

func GetTestCodeInfo() ([]string, []string) {
	app := SetDB()

	insertQuery := "SELECT test_code, tags FROM testcode"
	rows, err := app.DB.Query(insertQuery)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var testCodeList []string
	var tagList []string
	for rows.Next() {
		testCode := ""
		tag := ""
		if err := rows.Scan(&testCode, &tag); err != nil {
			log.Fatal(err)
		}
		testCodeList = append(testCodeList, testCode)
		tagList = append(tagList, tag)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	return testCodeList, tagList
}

// give originalVideos and return real name of original and artifact Videos name
func GetVideoNameListFromVideoList(videoList []string) ([]string, []string) {
	app := SetDB()

	var originalVideoNameList []string
	var artifactVideoNameList []string
	for _, videoNum := range videoList {
		query := "SELECT original_video_name, artifact_video_name FROM video WHERE original_video = ?"
		var originalVideoName string
		var artifactVideoName string
		err := app.DB.QueryRow(query, videoNum).Scan(&originalVideoName, &artifactVideoName)
		if err != nil {
			fmt.Println(err)
		}
		originalVideoNameList = append(originalVideoNameList, originalVideoName)
		artifactVideoNameList = append(artifactVideoNameList, artifactVideoName)
	}
	return originalVideoNameList, artifactVideoNameList
}

func GetVideoListFromTestCode(testCode string) (string, error) {
	app := SetDB()

	query := "SELECT video_list FROM testcode WHERE test_code = ?"
	var videoIdList string
	err := app.DB.QueryRow(query, testCode).Scan(&videoIdList)
	if err != nil {
		return "", err
	}
	return videoIdList, nil
}

func GetSameTagVideo(tag string) []string {
	app := SetDB()
	// Query to fetch data from the table
	insertQuery := "SELECT original_video FROM videos WHERE tag = ?"
	rows, err := app.DB.Query(insertQuery, tag)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Slice to store the retrieved data
	var videoDataSlice []string

	// Iterate through the result set and populate the struct
	for rows.Next() {
		videoData := ""
		if err := rows.Scan(&videoData); err != nil {
			log.Fatal(err)
		}
		videoDataSlice = append(videoDataSlice, videoData)
	}

	// Check for errors from iterating over rows
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	// Now, you can use the tagDataSlice in your Go code
	for _, data := range videoDataSlice {
		fmt.Println("Tag:", data)
	}
	return videoDataSlice
}

func GetTagData() []string {
	app := SetDB()
	// Query to fetch data from the table
	rows, err := app.DB.Query("SELECT tag FROM tag")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Slice to store the retrieved data
	var tags []string

	// Iterate through the result set and populate the struct
	for rows.Next() {
		tag := ""
		if err := rows.Scan(&tag); err != nil {
			log.Fatal(err)
		}
		tags = append(tags, tag)
	}

	// Check for errors from iterating over rows
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	// Now, you can use the tagDataSlice in your Go code
	for _, data := range tags {
		fmt.Println("Tag:", data)
	}
	return tags
}

func GetTagUUID(tag string) (uuid.UUID, error) {
	app := SetDB()

	query := "SELECT BIN_TO_UUID(uuid) FROM tag WHERE tag = ?"
	var tagUUID uuid.UUID
	err := app.DB.QueryRow(query, tag).Scan(&tagUUID)
	if err != nil {
		return uuid.Nil, err
	}

	return tagUUID, nil
}

func GetTagUUIDFromTestCode(testcode string) ([]uuid.UUID, error) {
	app := SetDB()

	query := "SELECT BIN_TO_UUID(tag_uuid) From testcode WHERE test_code = ?"
	rows, err := app.DB.Query(query, testcode)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	var tagsUUID []uuid.UUID

	// Iterate through the result set and populate the struct
	for rows.Next() {
		tagUUID := uuid.Nil
		if err := rows.Scan(&tagUUID); err != nil {
			log.Fatal(err)
		}
		tagsUUID = append(tagsUUID, tagUUID)
	}

	// Check for errors from iterating over rows
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	return tagsUUID, nil
}

// func GetTagList() ([]string, error) {
// 	// app := SetDB()

// 	// query :=
// }

func GetTag(testcodeUUID uuid.UUID) (string, error) {
	app := SetDB()

	query := "SELECT tag FROM tag WHERE BIN_TO_UUID(uuid) = ?"
	var tag string
	err := app.DB.QueryRow(query, testcodeUUID).Scan(&tag)
	if err != nil {
		return "uuid.Nil", err
	}

	return tag, nil
}

// return original_video
func GetVideoListFromTag(tag string) ([]string, error) {
	app := SetDB()

	// Query to fetch data from the table
	query := "SELECT original_video FROM video WHERE tag = ?"
	rows, err := app.DB.Query(query, tag)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	// Slice to store the retrieved data
	var videoList []string

	// Iterate through the result set and populate the struct
	for rows.Next() {
		video := ""
		if err := rows.Scan(&video); err != nil {
			log.Fatal(err)
		}
		videoList = append(videoList, video)
	}

	// Check for errors from iterating over rows
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	return videoList, nil
}
