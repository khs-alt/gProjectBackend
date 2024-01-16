package handler

import (
	"backend/sql"
	"backend/util"
	"fmt"
	"log"
	"net/http"

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
	imageList, _ := sql.GetImageListFromTestCode(testCode) // you have to control error

	c.JSON(http.StatusOK, gin.H{
		"current_page": currentPage,
		"image_list":   imageList,
	})
}

// 유저 ID와 Testcode를 받아서
// 유저가 테스트를 진행해야하는 Page를 알려줌
// 상단 함수와 중복됨
// func GetUserCurrentPage(c *gin.Context) {
// 	var data map[string]interface{}
// 	if err := c.ShouldBindJSON(&data); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	id := data["userID"].(string)
// 	testCode := data["testcode"].(string)

// 	currentPage := fmt.Sprint(sql.GetUserCurrentPageAboutTestCode(id, testCode))
// 	c.JSON(http.StatusOK, gin.H{
// 		"current_page": currentPage,
// 	})
// }

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
	fmt.Println("data: ", data)
	var originalVideoNameList []string
	var artifactVideoNameList []string
	var videoFPSList []string
	var indexList []string
	originalVideoNameList, artifactVideoNameList, videoFPSList, indexList, err := sql.GetVideoListFromTestCode(data.TestCode)
	if err != nil {
		log.Println(err)
	}
	//videoIndex, userScore := sql.GetCurrentUserScore(data.ID, data.TestCode)
	userScore := sql.GetUserScoreFromVideo(data.ID, data.CurrentPage, data.TestCode)
	//fmt.Println("videoIndex: ", videoIndex)
	fmt.Println("userScore: ", userScore)
	fmt.Println("videoList: ", indexList)
	randIndexList := util.ShuffleList(data.ID, indexList)
	randOrigianlVideoNameList := util.ShuffleList(data.ID, originalVideoNameList)
	randArtifactVideoNameList := util.ShuffleList(data.ID, artifactVideoNameList)
	randVideoFPSList := util.ShuffleList(data.ID, videoFPSList)
	c.JSON(http.StatusOK, gin.H{
		// "currentPage":           videoIndex,
		"videoList":             randIndexList,
		"originalVideoNameList": randOrigianlVideoNameList,
		"artifactVideoNameList": randArtifactVideoNameList,
		"originalVideoFPSList":  randVideoFPSList,
		"artifactVideoFPSList":  randVideoFPSList,
		"userScore":             userScore,
	})

}
