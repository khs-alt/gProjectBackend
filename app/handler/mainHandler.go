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

	fmt.Println("GetUserCurrentImagePage, currentPage", currentPage)

	c.JSON(http.StatusOK, gin.H{
		"current_page": currentPage,
		"image_list":   imageList,
	})
}

type UserInfo struct {
	ID          string `json:"userID"`
	TestCode    string `json:"testcode"`
	CurrentPage string `json:"currentPage"`
}

// it too long we have to make short
func GetUserCurrentPageInfo(c *gin.Context) {
	var data UserInfo
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var originalVideoNameList []string
	var artifactVideoNameList []string
	var videoFPSList []string
	var indexList []string
	originalVideoNameList, artifactVideoNameList, _, videoFPSList, indexList, err := sql.GetVideoListFromTestCode(data.TestCode)
	if err != nil {
		log.Println(err)
	}

	userScore := sql.GetUserScoreFromVideo(data.ID, data.CurrentPage, data.TestCode)

	randIndexList := util.ShuffleList(data.ID, indexList)
	randOrigianlVideoNameList := util.ShuffleList(data.ID, originalVideoNameList)
	randArtifactVideoNameList := util.ShuffleList(data.ID, artifactVideoNameList)
	var originalVideoList []string
	var artifactVideoList []string
	for _, originalVideo := range randOrigianlVideoNameList {
		s := util.RemoveSpecificPart(originalVideo)
		originalVideoList = append(originalVideoList, s)
	}
	for _, artifactVideo := range randArtifactVideoNameList {
		s := util.RemoveSpecificPart(artifactVideo)
		artifactVideoList = append(artifactVideoList, s)
	}
	randVideoFPSList := util.ShuffleList(data.ID, videoFPSList)
	c.JSON(http.StatusOK, gin.H{
		// "currentPage":           videoIndex,
		"videoList":             randIndexList,
		"originalVideoNameList": originalVideoList,
		"artifactVideoNameList": artifactVideoList,
		"originalVideoFPSList":  randVideoFPSList,
		"artifactVideoFPSList":  randVideoFPSList,
		"userScore":             userScore,
	})

}
