package handler

import (
	"backend/app/models"
	"backend/sql"
	"backend/util"
	"fmt"
	"log"
	"net/http"
	"os"
	"sort"
	"strconv"

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

//done

func UploadVideoHandler(c *gin.Context) {
	err := c.Request.ParseMultipartForm(5000 << 20)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	tagCSV, _ := c.GetPostForm("tags")
	form, err := c.MultipartForm()
	if err != nil {
		c.String(http.StatusBadRequest, "get form err: %s", err.Error())
		return
	}
	originalVideos := form.File["original"]
	originalVideosName, oriVideos, originalVideosFileForm := util.UploadVideo(c, originalVideos, "original")

	artifactVideos := form.File["artifact"]
	artifactVideosName, artiVideos, arfectVideosFileForm := util.UploadVideo(c, artifactVideos, "artifact")

	differentVideos := form.File["diff"]
	diffVideosName, _, _ := util.UploadVideo(c, differentVideos, "diff")
	tags := util.MakeCSVToStringList(tagCSV)
	sort.Strings(tags)

	//tags = strings.Split(tags[0], ",")
	for i := 0; i < len(oriVideos) && i < len(artiVideos); i++ {
		uuid := util.MakeUUID()
		originaVideoFPS := util.GetVideoFPS("./originalVideos/" + oriVideos[i] + originalVideosFileForm[i])
		artifactVideoFPS := util.GetVideoFPS("./artifactVideos/" + artiVideos[i] + arfectVideosFileForm[i])
		if originaVideoFPS != artifactVideoFPS {
			fmt.Printf("%s와 %s의 비디오 프레임이 다릅니다 \n", originalVideosName[i], artifactVideosName[i])
		}
		width, height, err := util.GetFileDimensions("./originalVideos/" + oriVideos[i] + originalVideosFileForm[i])
		if err != nil {
			log.Println(err)
		}
		err = sql.InsertVideo(uuid, originalVideosName[i], artifactVideosName[i], diffVideosName[i], originaVideoFPS, width, height)
		if err != nil {
			log.Println(err)
		}
		for _, tag := range tags {
			err = sql.InsertVideoTagLink(uuid, tag)
			if err != nil {
				log.Println(err)
			}
		}
	}
	c.String(http.StatusOK, "비디오가 성공적으로 업로드되었습니다.")
}

func PostVideoFrameTimeHandler(c *gin.Context) {
	var data models.VideoFrameTimeData
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	videoIndex := strconv.Itoa(data.VideoIndex)
	//videoIndex := fmt.Sprint(data.VideoIndex)
	videoFilePath := fmt.Sprintf("./artifactVideos/artifactVideo%s.mp4", videoIndex)
	videoCurrentTime := data.VideoCurrentTime
	outputImage := fmt.Sprintf("./selectedFrame/selectedFrame%s_%f.png", videoIndex, videoCurrentTime)
	err := util.ExtractFrame(videoFilePath, videoCurrentTime, outputImage)
	if err != nil {
		log.Println("error: ", err)
		return
	}

	c.String(http.StatusOK, "Success insert frame time")
}
