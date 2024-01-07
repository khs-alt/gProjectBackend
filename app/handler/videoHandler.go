package handler

import (
	"backend/app/models"
	"backend/sql"
	"backend/util"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func ServeOriginalVideosHandler(c *gin.Context) {
	videoID := c.Param("id")
	videoFilePrefix := fmt.Sprintf("./originalVideos/originalVideo%s", videoID)
	var videoFilePath string
	var fileExtension string
	for ext := range models.MimeTypes {
		tempPath := videoFilePrefix + ext
		if _, err := os.Stat(tempPath); err == nil {
			videoFilePath = tempPath
			fileExtension = ext
			break
		}
	}
	if videoFilePath == "" {
		c.JSON(http.StatusNotFound, gin.H{"error": "Video not found"})
		return
	}
	mimeType, exists := models.MimeTypes[fileExtension]
	if !exists {
		c.JSON(http.StatusUnsupportedMediaType, gin.H{"error": "Unsupported video format"})
		return
	}
	c.Header("Content-Type", mimeType)
	c.File(videoFilePath) // Serve file using Gin

}

func ServeArtifactVideosHandler(c *gin.Context) {
	videoID := c.Param("id")
	videoFilePrefix := fmt.Sprintf("./artifactVideos/artifactVideo%s", videoID)
	var videoFilePath string
	var fileExtension string
	for ext := range models.MimeTypes {
		tempPath := videoFilePrefix + ext
		if _, err := os.Stat(tempPath); err == nil {
			videoFilePath = tempPath
			fileExtension = ext
			break
		}
	}
	if videoFilePath == "" {
		c.JSON(http.StatusNotFound, gin.H{"error": "Video not found"})
		return
	}
	mimeType, exists := models.MimeTypes[fileExtension]
	if !exists {
		c.JSON(http.StatusUnsupportedMediaType, gin.H{"error": "Unsupported video format"})
		return
	}
	c.Header("Content-Type", mimeType)
	c.File(videoFilePath) // Serve file using Gin

}

func UploadVideoHandler(c *gin.Context) {
	err := c.Request.ParseMultipartForm(5000 << 20)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	tags, _ := c.GetPostFormArray("tags")

	form, err := c.MultipartForm()
	if err != nil {
		c.String(http.StatusBadRequest, "get form err: %s", err.Error())
		return
	}
	originalVideos := form.File["original"]
	originalVideosName, oriVideos, originalVideosFileForm := util.UploadVideo(c, originalVideos, "original")

	artifactVideos := form.File["artifact"]
	artifactVideosName, artiVideos, arfectVideosFileForm := util.UploadVideo(c, artifactVideos, "artifact")

	tags = strings.Split(tags[0], ",")
	for i := 0; i < len(oriVideos) && i < len(artiVideos); i++ {
		for j := 0; j < len(tags); j++ {
			uuid := util.MakeUUID()
			originaVideoFPS := util.GetVideoFPS("./originalVideos/" + oriVideos[i] + originalVideosFileForm[i])
			artifactVideoFPS := util.GetVideoFPS("./artifactVideos/" + artiVideos[i] + arfectVideosFileForm[i])
			err := sql.InsertVideoId(uuid, originalVideosName[i], oriVideos[i], originaVideoFPS, artifactVideosName[i], artiVideos[i], artifactVideoFPS, tags[j])
			if err != nil {
				fmt.Println(err)
			}
		}
	}
	c.String(http.StatusOK, "비디오가 성공적으로 업로드되었습니다.")
}
