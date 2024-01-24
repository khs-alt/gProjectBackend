package handler

import (
	"backend/app/models"
	"backend/sql"
	"backend/util"
	"log"
	"net/http"
	"sort"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// done
func GetTestCodeListHandler(c *gin.Context) {
	testCodeList, tagList := sql.GetTestCodeInfo()
	newTestCodeList, newtagList := util.MakeNewTestCodeTagList(testCodeList, tagList)
	c.JSON(http.StatusOK, gin.H{
		"testcode": newTestCodeList,
		"tags":     newtagList,
	})
}

// done
func GetImageTestCodeListHandler(c *gin.Context) {
	testCodeList, tagList := sql.GetImageTestCodeInfo()
	newTestCodeList, newtagList := util.MakeNewTestCodeTagList(testCodeList, tagList)
	c.JSON(http.StatusOK, gin.H{
		"testcode": newTestCodeList,
		"tags":     newtagList,
	})
}

// you have to use join or sql
func GetVideoListFromTagHandler(c *gin.Context) {
	// 이 함수는 제대로 []string으로 들어옴
	tags := c.QueryArray("tag[]")
	sort.Strings(tags)

	originalVideo, err := sql.GetVideoListFromTag(tags)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"original_video_list": originalVideo,
	})
}

// use join
func GetImageListFromTagHandler(c *gin.Context) {
	tags := c.QueryArray("tag[]")
	sort.Strings(tags)

	originalIamge, err := sql.GetImageListFromTag(tags)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"original_image_list": originalIamge,
	})

}

// use join it can't
// TODO: description is not used
// done
func GetVideoTestCodeHandler(c *gin.Context) {
	var data models.RequestData
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	testcode := util.GenerateRandomString(8)
	description := ""
	for _, tag := range data.Tags {
		uuid, _ := uuid.NewUUID()
		sql.InsertVideoTestCode(uuid, tag, testcode, description)
		c.String(http.StatusOK, "Success insert testcode")
	}
}

// TODO: description is not used
// done
func GetImageTestCodeHandler(c *gin.Context) {
	var data models.RequestData
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	testcode := util.GenerateRandomString(8)
	description := ""
	for _, tag := range data.Tags {
		uuid, _ := uuid.NewUUID()
		sql.InsertImageTestCode(uuid, tag, testcode, description)
		c.String(http.StatusOK, "Success insert testcode")
	}
}

func GetVideoListFromTestCodeHandler(c *gin.Context) {
	testCode := c.Query("testcode")
	isVideoTestcodeExist, err := sql.GetVideoTestcodeExist(testCode)
	if err != nil {
		log.Println("GetVideoTestcodeExist error")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
	}
	if !isVideoTestcodeExist {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Wrong Testcode"})
	}
	originalVideoNameList, arfectVideosNameList, videoFrameList, videoList, err := sql.GetVideoListFromTestCode(testCode)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"video_list":          videoList,
		"original_video_list": originalVideoNameList,
		"artifact_video_list": arfectVideosNameList,
		"video_frame_list":    videoFrameList,
	})
}
