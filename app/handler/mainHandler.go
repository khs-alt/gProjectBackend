package handler

import (
	"backend/sql"
	"backend/util"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func MainHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello this is main Page!!!!",
	})
}

// 유저 ID와 Testcode를 받아서
// 유저가 테스트를 진행해야하는 Page와 image_list들을 알려줌
func GetUserCurrentImagePage(c *gin.Context) {
	var data map[string]interface{}
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id := data["userID"].(string)
	testCode := data["testcode"].(string)

	currentPage := sql.GetUserCurrentImagePageAboutTestCode(id, testCode)
	imageCSVList, _ := sql.GetImageListFromTestCode(testCode) // you have to control error

	imageList := util.MakeCSVToStringList(imageCSVList)
	c.JSON(http.StatusOK, gin.H{
		"current_page": currentPage,
		"image_list":   imageList,
	})
}

// 유저 ID와 Testcode를 받아서
// 유저가 테스트를 진행해야하는 Page를 알려줌
// 상단 함수와 중복됨
func GetUserCurrentPage(c *gin.Context) {
	var data map[string]interface{}
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id := data["userID"].(string)
	testCode := data["testcode"].(string)

	currentPage := fmt.Sprint(sql.GetUserCurrentPageAboutTestCode(id, testCode))
	c.JSON(http.StatusOK, gin.H{
		"current_page": currentPage,
	})
}

type UserInfo struct {
	ID          string `json:"userID"`
	TestCode    string `json:"testcode"`
	CurrentPage string `json:"currentPage"`
}

// it too long we have to make short
// TODO:use join
func GetUserCurrentPageInfo(c *gin.Context) {
	var data UserInfo
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	videoCSVList, err := sql.GetVideoListFromTestCode(data.TestCode)
	var videoNumList []string
	var originalVideoNameList []string
	var artifactVideoNameList []string
	var originalVideoFPSList []string
	var artifactVideoFPSList []string
	videoList := strings.Split(videoCSVList, ",")
	originalVideoNameList, artifactVideoNameList = sql.GetVideoNameListFromVideoList(videoList)
	for _, video := range videoList {
		num := strings.TrimLeft(video, "originalVideo")
		videoNumList = append(videoNumList, num)
	}
	for _, video := range videoList {
		originalVideoFPS, artifactVideoFPS := sql.GetFPSFromVideo(video)
		originalVideoFPSList = append(originalVideoFPSList, strconv.FormatFloat(float64(originalVideoFPS), 'f', 2, 32))
		artifactVideoFPSList = append(artifactVideoFPSList, strconv.FormatFloat(float64(artifactVideoFPS), 'f', 2, 32))
	}

	curPage, err := strconv.Atoi(data.CurrentPage)
	if err != nil {
		log.Println(err)
	}
	userScore := sql.GetCurrentUserScore(data.ID, curPage)
	c.JSON(http.StatusOK, gin.H{
		"CurrentPage":           data.CurrentPage,
		"VideoList":             videoList,
		"OriginalVideoNameList": originalVideoNameList,
		"ArtifactVideoNameList": artifactVideoNameList,
		"OriginalVideoFPSList":  originalVideoFPSList,
		"ArtifactVideoFPSList":  artifactVideoFPSList,
		"UserScore":             userScore,
	})

}
