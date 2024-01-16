package sql

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	"github.com/google/uuid"
)

func GetUserScoreFromVideo(userId string, videoIndex string, videoTestcode string) int {
	app := SetDB()
	insertQuery := `
					SELECT
						vs.user_score
					FROM
						video_scoring vs
					JOIN
						user u ON u.uuid = vs.user_uuid
					JOIN
						video v ON vs.video_uuid = v.uuid
					WHERE
						u.user_name = ? AND v.video_index = ? AND vs.video_testcode = ?
					ORDER BY 
						time DESC
					LIMIT 1
					`
	var userScore int
	err := app.DB.QueryRow(insertQuery, userId, videoIndex, videoTestcode).Scan(&userScore)
	if err != nil {
		if err == sql.ErrNoRows {
			// 결과가 없을 때 -1 반환
			fmt.Println("GetUserScoreFromVideo no rows in result set")
			return -1
		} else {
			// 다른 오류가 발생한 경우 로그를 출력
			log.Println("GetUserScoreFromVideo is error: ", err)
		}
	}
	return userScore
}

// return videoIndex, score
func GetCurrentUserScore(userId string, videoTestCode string) (int, int) {
	app := SetDB()
	var userUUID uuid.UUID
	userUUIDQuery := "SELECT BIN_TO_UUID(uuid) FROM user WHERE user_name = ?"
	err := app.DB.QueryRow(userUUIDQuery, userId).Scan(&userUUID)
	if err != nil {
		log.Printf("Error finding user_uuid: %v\n", err)
	}

	insertQuery := `
					SELECT						
						COALESCE(user_score, -1) AS user_score, BIN_TO_UUID(video_uuid) AS video_uuid
					FROM
						video_scoring
					WHERE
						BIN_TO_UUID(user_uuid) = ? AND video_testcode =  ?
					ORDER BY
						time DESC
					LIMIT 1
					`
	var score int
	var videoUUID uuid.UUID
	err = app.DB.QueryRow(insertQuery, userUUID, videoTestCode).Scan(&score, &videoUUID)
	if err != nil {
		if err == sql.ErrNoRows {
			// 결과가 없을 때 -1 반환
			fmt.Println("GetCurrentUserScore: no rows in result set")
			return 1, -1
		} else {
			// 다른 오류가 발생한 경우 로그를 출력
			log.Println("GetCurrentUserScore: is error: ", err)
		}
	}
	var videoIndex int
	videoIndexQuery := `
					SELECT
						video_index
					FROM
						video
					WHERE
						BIN_TO_UUID(uuid) = ?
					`
	err = app.DB.QueryRow(videoIndexQuery, videoUUID).Scan(&videoIndex)
	if err != nil {
		log.Println("GetCurrentUserScore is error: ", err)
	}

	return videoIndex, score
}

func GetCurrentUserScoreList(userId string, videoIndexList []string) []int {
	app := SetDB()
	var userUUID uuid.UUID
	userUUIDQuery := "SELECT BIN_TO_UUID(uuid) FROM user WHERE user_name = ?"
	err := app.DB.QueryRow(userUUIDQuery, userId).Scan(&userUUID)
	if err != nil {
		log.Printf("Error finding user_uuid: %v\n", err)
	}

	videoUUIDQuery := `
					SELECT
						BIN_TO_UUID(uuid)
					FROM
						video
					WHERE
						video_index = ?
					`
	var videoUUID uuid.UUID
	var videoUUIDList []uuid.UUID
	for i := 0; i < len(videoIndexList); i++ {
		err = app.DB.QueryRow(videoUUIDQuery, videoIndexList[i]).Scan(&videoUUID)
		if err != nil {
			log.Println("GetCurrentUserScore is error: ", err)
		}
		videoUUIDList = append(videoUUIDList, videoUUID)
	}
	scoreListQuery := `
					SELECT
						user_score
					FROM
						video_scoring
					WHERE
						BIN_TO_UUID(user_uuid) = ? AND BIN_TO_UUID(video_uuid) =  ?
					ORDER BY
						time DESC
					LIMIT 1
					`
	var scoreList []int
	var score int
	for i := 0; i < len(videoUUIDList); i++ {
		err = app.DB.QueryRow(scoreListQuery, userUUID, videoUUIDList[i]).Scan(&score)
		if err != nil {
			if err == sql.ErrNoRows {
				// 결과가 없을 때 -1 반환
				fmt.Println("no rows in result set")
				scoreList = append(scoreList, -1)
				continue
			} else {
				// 다른 오류가 발생한 경우 로그를 출력
				log.Println("GetCurrentUserScore is error: ", err)
			}
		}
		scoreList = append(scoreList, score)
	}

	return scoreList
}

func GetCurrentUserImageScore(userId string, imageId int) string {
	app := SetDB()

	insertQuery := `				
				SELECT	
				iscore.patch_score
				FROM	
					image_scoring AS iscore
				JOIN
					user AS u ON u.uuid = iscore.user_uuid
				JOIN
					image AS i ON iscore.image_uuid = i.uuid
				WHERE
					u.user_name = ? AND i.image_index = ?
				ORDER BY
					iscore.time DESC
				LIMIT 1				
			`
	var score string
	err := app.DB.QueryRow(insertQuery, userId, imageId).Scan(&score)
	if err != nil {
		if err == sql.ErrNoRows {
			// 결과가 없을 때 -1 반환
			return "-1"
		} else {
			// 다른 오류가 발생한 경우 로그를 출력
			fmt.Println("GetCurrentUserImageScore error: ", err)
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

// TODO:
func GetUserCurrentImagePageAboutTestCode(userId string, testCode string) int {
	app := SetDB()
	maxCurrentPage := 0

	insertQuery := `
				SELECT 
					uti.last_page
				FROM 
					user_testcode_info uti
				JOIN 
					user u ON u.uuid = uti.user_uuid 
				WHERE 
					u.user_name = ? AND uti.test_code = ? AND uti.is_video = 0
				ORDER BY 
					time DESC 
				LIMIT 1
	`
	err := app.DB.QueryRow(insertQuery, userId, testCode).Scan(&maxCurrentPage)
	if err == sql.ErrNoRows {
		// 결과가 없을 때 1 반환
		log.Println("GetUserCurrentPageAboutTestCode no rows in result set")
		return 1
	} else {
		// 다른 오류가 발생한 경우 로그를 출력
		log.Println("GetUserCurrentPageAboutTestCode is error: ", err)
	}
	return maxCurrentPage
}

// TODO:
// func GetUserCurrentPageAboutTestCode(userId string, testCode string) int {
// 	app := SetDB()
// 	currentPage := 0

// 	insertQuery := "SELECT current_page FROM user_testcode_info WHERE user_id = ? AND test_code = ? ORDER BY time DESC LIMIT 1"
// 	err := app.DB.QueryRow(insertQuery, userId, testCode).Scan(&currentPage)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	return currentPage
// }

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

// done
func GetVideoTestcodeExist(testCode string) (bool, error) {
	app := SetDB()

	query := "SELECT EXISTS (SELECT * FROM video_testcode WHERE video_testcode = ?)"
	var exist bool
	err := app.DB.QueryRow(query, testCode).Scan(&exist)
	if err != nil {
		log.Println(err)
		return false, err
	}
	return exist, nil
}

// done
func GetImageTestcodeExist(testCode string) (bool, error) {
	app := SetDB()

	query := "SELECT EXISTS (SELECT * FROM image_testcode WHERE image_testcode = ?)"
	var exist bool
	err := app.DB.QueryRow(query, testCode).Scan(&exist)
	if err != nil {
		log.Println(err)
		return false, err
	}
	return exist, nil
}

// func GetTestCodeCount() (int, error) {
// 	app := SetDB()

// 	var count int
// 	err := app.DB.QueryRow("SELECT COUNT(*) FROM testcode").Scan(&count)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return count, nil
// }

// func GetImageTestCodeCount() (int, error) {
// 	app := SetDB()

// 	var count int
// 	err := app.DB.QueryRow("SELECT COUNT(*) FROM image_testcode").Scan(&count)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return count, nil
// }

func GetTestCodeInfo() ([]string, []string) {
	app := SetDB()

	insertQuery := "SELECT video_testcode, video_tag FROM video_testcode"
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

func GetImageTestCodeInfo() ([]string, []string) {
	app := SetDB()

	insertQuery := "SELECT image_testcode, image_tag FROM image_testcode"
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

func GetImageNameListFromVideoList(videoList []string) ([]string, []string) {
	app := SetDB()

	var originalImageNameList []string
	var artifactimageNameList []string
	for _, videoNum := range videoList {
		query := "SELECT original_image_name, artifact_image_name FROM image WHERE image_index = ?"
		var originalImageName string
		var artifactImageName string
		err := app.DB.QueryRow(query, videoNum).Scan(&originalImageName, &artifactImageName)
		if err != nil {
			fmt.Println("GetImageNameListFromVideoList ", err)
		}
		originalImageNameList = append(originalImageNameList, originalImageName)
		artifactimageNameList = append(artifactimageNameList, artifactImageName)
	}
	return originalImageNameList, artifactimageNameList
}

func GetVideoListInfoFromTestCode(testCode string) ([]string, []string, []string, []string, error) {
	app := SetDB()

	qeury := `SELECT v.video_index, v.original_video_name, v.artifact_video_name, v.video_frame
	FROM	
		video v
	JOIN
		video_testcode vt ON v.video_tag = vt.video_tag
	WHERE
		vt.video_testcode = ?`
	var videoList, originalVideoNameList, arfectVideosNameList, videoFrameList []string
	err := app.DB.QueryRow(qeury, testCode).Scan(&videoList, &originalVideoNameList, &arfectVideosNameList, &videoFrameList)
	if err != nil {
		return nil, nil, nil, nil, err
	}
	return videoList, originalVideoNameList, arfectVideosNameList, videoFrameList, nil

}

// done
func GetVideoListFromTestCode(testCode string) ([]string, []string, []string, []string, error) {
	app := SetDB()

	query := `
	SELECT DISTINCT
		v.original_video_name, v.artifact_video_name, v.video_frame, v.video_index
	FROM
    	video_testcode vt
	JOIN
    	video_tag t ON vt.video_tag = t.tag
	JOIN
    	video_tag_link vtl ON t.uuid = vtl.tag_uuid
	JOIN
    	video v ON vtl.video_uuid = v.uuid
	WHERE
    	vt.video_testcode = ?
	`
	rows, err := app.DB.Query(query, testCode)
	if err != nil {
		log.Println(err)
		return nil, nil, nil, nil, err
	}
	defer rows.Close()

	var originalVideoNameList, arfectVideosNameList, videoFrameList, indexList []string
	var originalVideoName, arfectVideosName, videoFrame, index string
	for rows.Next() {
		err := rows.Scan(&originalVideoName, &arfectVideosName, &videoFrame, &index)
		if err != nil {
			log.Println(err)
			return nil, nil, nil, nil, err
		}
		originalVideoNameList = append(originalVideoNameList, originalVideoName)
		arfectVideosNameList = append(arfectVideosNameList, arfectVideosName)
		videoFrameList = append(videoFrameList, videoFrame)
		indexList = append(indexList, index)
	}
	if err := rows.Err(); err != nil {
		log.Println(err)
		return nil, nil, nil, nil, err
	}
	return originalVideoNameList, arfectVideosNameList, videoFrameList, indexList, nil
}

func GetImageListFromTestCode(testCode string) ([]string, error) {
	app := SetDB()

	query := `
			SELECT DISTINCT
				i.image_index
			FROM
				image i
			JOIN
				image_tag_link itl ON i.uuid = itl.image_uuid
			JOIN
				image_tag it ON itl.tag_uuid = it.uuid
			JOIN
				image_testcode itc ON it.tag = itc.image_tag
			WHERE
				itc.image_testcode = ?
			`
	var videoIdList []string
	raw, err := app.DB.Query(query, testCode)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer raw.Close()

	for raw.Next() {
		var videoId string
		if err := raw.Scan(&videoId); err != nil {
			log.Println(err)
			return nil, err
		}
		videoIdList = append(videoIdList, videoId)
	}
	if err := raw.Err(); err != nil {
		log.Println(err)
		return nil, err
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
	// for _, data := range videoDataSlice {
	// 	fmt.Println("Tag:", data)
	// }
	return videoDataSlice
}

func GetVideoTag() []string {
	app := SetDB()
	// Query to fetch data from the table
	rows, err := app.DB.Query("SELECT tag FROM video_tag")
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
	return tags
}
func GetImageTag() []string {
	app := SetDB()
	// Query to fetch data from the table
	rows, err := app.DB.Query("SELECT tag FROM image_tag")
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
func GetVideoListFromTag(tags []string) ([]string, error) {
	app := SetDB()

	new_tags := `"` + strings.Join(tags, `","`) + `"`
	fmt.Println(new_tags)
	// Query to fetch data from the table
	query := fmt.Sprintf(`
			SELECT DISTINCT
    			v.original_video_name
			FROM
    			video v
			JOIN
    			video_tag_link vtl ON v.uuid = vtl.video_uuid
			JOIN
    			video_tag vt ON vtl.tag_uuid = vt.uuid
			WHERE
    			vt.tag IN (%s)
			`, new_tags)
	rows, err := app.DB.Query(query)
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
			log.Println(err)
		}
		videoList = append(videoList, video)
	}

	// Check for errors from iterating over rows
	if err := rows.Err(); err != nil {
		log.Println(err)
	}
	return videoList, nil
}

// TODO: Here has a problem maybe
func GetImageListFromTag(tags []string) ([]string, error) {
	app := SetDB()

	// Query to fetch data from the table
	new_tags := `"` + strings.Join(tags, `","`) + `"`
	fmt.Println(new_tags)
	// Query to fetch data from the table
	query := fmt.Sprintf(`
			SELECT DISTINCT
    			i.original_image_name
			FROM
				image i
			JOIN
    			image_tag_link itl ON i.uuid = itl.image_uuid
			JOIN
    			image_tag it ON itl.tag_uuid = it.uuid
			WHERE
    			it.tag IN (%s)
			`, new_tags)
	rows, err := app.DB.Query(query)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	// Slice to store the retrieved data
	var imageList []string

	// Iterate through the result set and populate the struct
	for rows.Next() {
		image := ""
		if err := rows.Scan(&image); err != nil {
			log.Fatal(err)
		}
		imageList = append(imageList, image)
	}

	// Check for errors from iterating over rows
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	return imageList, nil
}

func GetUserScoringList(user string, testCode string) []int {
	app := SetDB()
	// query := `SELECT
	// 			user_score
	// 		  FROM
	// 		  	video_scoring
	// 		  JOIN
	// 		  	user ON video_scoring.user_uuid = user.uuid
	// 		  WHERE
	// 		  	user.user_name = ? AND video_scoring.video_testcode = ?`
	query := `
	SELECT 
    	IFNULL(vs.user_score, -1) AS user_score 
	FROM 
    	video v 
	JOIN 
    	video_testcode vtc ON v.tag = vtc.video_tag 
	LEFT JOIN 
    	video_scoring vs ON BIN_TO_UUID(v.uuid) = BIN_TO_UUID(vs.video_uuid) AND BIN_TO_UUID(vs.user_uuid) = (SELECT BIN_TO_UUID(uuid) FROM user WHERE user_name = ?) 
	WHERE 
    	vtc.video_testcode = ?;
`
	rows, err := app.DB.Query(query, user, testCode)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	// Slice to store the retrieved data
	var userScoringList []int

	for rows.Next() {
		score := 0
		if err := rows.Scan(&score); err != nil {
			log.Println(err)
		}
		userScoringList = append(userScoringList, score)
	}
	if err := rows.Err(); err != nil {
		log.Println(err)
	}
	return userScoringList
}
