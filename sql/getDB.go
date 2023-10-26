package sql

import (
	"fmt"
	"log"

	"github.com/google/uuid"
)

func GetUserCurrentPageAboutTestCode(userId string, testCode string) int {
	app := SetDB()
	maxCurrentPage := 0

	insertQuery := "SELECT current_page FROM user_testcode_info WHERE user_id = ? AND test_code = ? ORDER BY uuid DESC LIMIT 1;"
	err := app.DB.QueryRow(insertQuery, userId, testCode).Scan(&maxCurrentPage)
	if err != nil {
		fmt.Println(err)
	}
	return maxCurrentPage
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
	insertQuery := "SELECT original_video_name FROM videos WHERE tag = ?"
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

func GetVideo(tag string) ([]string, error) {
	app := SetDB()

	// Query to fetch data from the table
	query := "SELECT original_video_name FROM video WHERE tag = ?"
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
