package handler

import (
	"backend/sql"
	"backend/util"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetTestCodeListHandler(c *gin.Context) {
	testCodeList, tagsList := sql.GetTestCodeInfo()
	c.JSON(http.StatusOK, gin.H{
		"testcode": testCodeList,
		"tags":     tagsList,
	})
}

func GetImageTestCodeListHandler(c *gin.Context) {
	testCodeList, tagsList := sql.GetImageTestCodeInfo()
	c.JSON(http.StatusOK, gin.H{
		"testcode": testCodeList,
		"tags":     tagsList,
	})
}

// you have to use join or sql
func GetVideoListFromTagHandler(c *gin.Context) {
	tags := c.QueryArray("tag[]")
	var videoList []string
	for _, tag := range tags {
		videos, _ := sql.GetVideoListFromTag(tag)
		for _, video := range videos {
			videoList = append(videoList, video)
		}
	}
	videoList = util.RemoveDuplicates(videoList)
	originalVideo, _ := sql.GetVideoNameListFromVideoList(videoList)
	c.JSON(http.StatusOK, gin.H{
		"original_video_list": originalVideo,
	})
}

// use join
func GetImageListFromTagHandler(c *gin.Context) {
	tags := c.QueryArray("tag[]")
	var videoList []string
	for _, tag := range tags {
		videos, _ := sql.GetImageListFromTag(tag)
		for _, video := range videos {
			videoList = append(videoList, video)
		}
	}
	videoList = util.RemoveDuplicates(videoList)
	originalIamge, _ := sql.GetImageNameListFromVideoList(videoList)
	c.JSON(http.StatusOK, gin.H{
		"original_image_list": originalIamge,
	})

}

// use join
func GetTestCodeHandler(c *gin.Context) {
	var data RequestData
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	uuid, _ := uuid.NewUUID()
	var videoList []string
	for _, tag := range data.Tags {
		videos, _ := sql.GetVideoListFromTag(tag)
		for _, video := range videos {
			videoList = append(videoList, video)
		}
	}
	videoList = util.RemoveDuplicates(videoList)
	num, _ := sql.GetTestCodeCount()

	testcode := util.GenerateRandomString(8) + fmt.Sprint(num)
	sql.InsertTestCodeId(uuid, testcode, data.Tags, videoList)
	c.String(http.StatusOK, "Success insert testcode")
}

func GetImageTestCodeHandler(c *gin.Context) {
	var data RequestData
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	uuid, _ := uuid.NewUUID()
	var imageList []string
	for _, tag := range data.Tags {
		images, _ := sql.GetImageListFromTag(tag)
		for _, image := range images {
			imageList = append(imageList, image)
		}
	}
	imageList = util.RemoveDuplicates(imageList)
	num, _ := sql.GetImageTestCodeCount()

	testcode := util.GenerateRandomString(8) + fmt.Sprint(num)
	sql.InsertImageTestCodeId(uuid, testcode, data.Tags, imageList)
	c.String(http.StatusOK, "Success insert testcode")
}
