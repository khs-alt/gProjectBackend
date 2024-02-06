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
	"github.com/google/uuid"
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

func ServeDiffVideosHandler(c *gin.Context) {
	videoID := c.Param("id")
	videoFilePrefix := fmt.Sprintf("./diffVideos/diffVideo%s", videoID)
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

// 비디오 아이디와 시간 리스트가 오면 그 시간에 해당하는 비디오 프레임의 잘라서 이미지로 생성함
// If you provide a video ID and a list of timestamps,
// generate images by cropping the corresponding video frames at those timestamps.
func PostVideoFrameTimeHandler(c *gin.Context) {
	var data models.VideoFrameTimeData
	if err := c.ShouldBindJSON(&data); err != nil {
		fmt.Println("PostVideoFrameTimeHandler error: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	videoIndex := strconv.Itoa(data.VideoIndex)
	//비디오 인덱스에 해당하는 비디오 path를 선택
	originalVideoPath := fmt.Sprintf("./originalVideos/originalVideo%s.mp4", videoIndex)
	artifactVideoPath := fmt.Sprintf("./artifactVideos/artifactVideo%s.mp4", videoIndex)
	diffVideoPath := fmt.Sprintf("./diffVideos/diffVideo%s.mp4", videoIndex)

	sql.DeleteVideoTime(data.VideoIndex)
	sort.Strings(data.VideoCurrentTimeList)
	sort.Strings(data.VideoFrame)

	fmt.Println("index: " + videoIndex)
	fmt.Println("data.VideoCurrentTimeList: ", data.VideoCurrentTimeList)
	fmt.Println("data.VideoFrame: ", data.VideoFrame)
	// 같은 비디오 인덱스를 가진 이미지를 삭제
	sql.DeleteImage(data.VideoIndex)
	for i, videoCurrentTime := range data.VideoCurrentTimeList {
		//저장할 비디오의 프레임 시간을 DB에 저장
		err := sql.InsertVideoTime(data.VideoIndex, data.VideoFrame[i], videoCurrentTime)
		if err != nil {
			log.Println("InsertVideoTime error: ", err)
			c.JSON(http.StatusBadRequest, gin.H{"InsertVideoTime error": err.Error()})
			return
		}
	}
	for _, videoCurrentTime := range data.VideoCurrentTimeList {

		// 구지 할 필요 없음 논리적 삭제를 하겠음
		//DeleteRealImage(videoIndex)
		originalVideoName, artifactVideoName, diffVideoName := sql.GetVideoNameForIndex(data.VideoIndex)
		originalVideoName = originalVideoName + "_frame" + videoCurrentTime
		artifactVideoName = artifactVideoName + "_frame" + videoCurrentTime
		diffVideoName = diffVideoName + "_frame" + videoCurrentTime

		imageUUID, _ := uuid.NewUUID()
		width, heigth, err := util.GetFileDimensions(originalVideoPath)
		if err != nil {
			log.Println("PostVideoFrameTimeHandler GetFileDimensions error: ", err)
			return
		}
		videoIntIndex, _ := strconv.Atoi(videoIndex)

		// 이미지 저장하기 이미지 데이터와 테그 데이터가 필요함
		// 하단은 이를 가져오는 함수들
		// fmt.Print("uuid: ", imageUUID, "originalVideoName : ", originalVideoName, " artifactVideoName : ", artifactVideoName, " diffVideoName : ", diffVideoName, " width : ", width, " heigth : ", heigth, " videoIntIndex : ", videoIntIndex, "\n")
		// 왜 한 개 밖에 저장이 안될까?
		err = sql.InsertImage(imageUUID, originalVideoName, artifactVideoName, diffVideoName, width, heigth, videoIntIndex)
		if err != nil {
			log.Println("PostVideoFrameTimeHandler InsertImage error: ", err)
			return
		}

		// 비디오에 인덱스가 추가 될 때는 비디오에서 추가될 때 이미지에도 추가되면 될 듯 한다.
		// 즉 비디오에 추가될 때 이미지에도 같이 추가되면 됨
		tags := sql.GetVideoTagList(videoIntIndex)
		tags = util.RemoveDuplicates(tags)
		for _, tag := range tags {
			err = sql.InsertImageTagLink(imageUUID, tag)
			if err != nil {
				log.Println("PostVideoFrameTimeHandler InsertImageTagLink error: ", err)
				return
			}

		}

		//각 이미지의 파일 개수를 세기 위한 작업

		orgiImagePath := "./" + "original" + "Images/"
		oriCount, err := util.CountFile(orgiImagePath)
		if err != nil {
			fmt.Print("CountFile error : ")
			log.Println(err)
		}

		artiImagePath := "./" + "artifact" + "Images/"
		artiCount, err := util.CountFile(artiImagePath)
		if err != nil {
			fmt.Print("CountFile error : ")
			log.Println(err)
		}

		diffImagePath := "./" + "diff" + "Images/"
		diffCount, err := util.CountFile(diffImagePath)
		if err != nil {
			fmt.Print("CountFile error : ")
			log.Println(err)
		}

		//센 파일의 개수를 이미지 이름에 넣어서 이미지를 저장

		oriImage := fmt.Sprintf("./originalImages/originalImage%d.png", oriCount+1)
		artImage := fmt.Sprintf("./artifactImages/artifactImage%d.png", artiCount+1)
		diffImage := fmt.Sprintf("./diffImages/diffImage%d.png", diffCount+1)

		//실제로 이미지를 잘라서 저장
		err = util.ExtractFrame(originalVideoPath, videoCurrentTime, oriImage)
		if err != nil {
			log.Println("error: ", err)
			return
		}
		err = util.ExtractFrame(artifactVideoPath, videoCurrentTime, artImage)
		if err != nil {
			log.Println("error: ", err)
			return
		}
		err = util.ExtractFrame(diffVideoPath, videoCurrentTime, diffImage)
		if err != nil {
			log.Println("error: ", err)
			return
		}
	}
	c.String(http.StatusOK, "Success insert frame time")
}

// 비디오 인덱스가 오면 어떤 프레임이 잘렸는지를 반환함
func GetSelectedFrameListHandler(c *gin.Context) {
	videoIndex := c.Query("video_index")
	videoIntIndex, _ := strconv.Atoi(videoIndex)
	timeList, selectedFrameList := sql.GetSelectedFrameList(videoIntIndex)
	if selectedFrameList == nil {
		selectedFrameList = []string{}
	}
	if timeList == nil {
		timeList = []string{}
	}
	fmt.Println("========= GetSelectedFrameListHandler =========")
	fmt.Println("selectedFrameList: ", selectedFrameList)
	fmt.Println("video_index: ", videoIndex)
	c.JSON(http.StatusOK, gin.H{
		"selected_video_frame_time_list": selectedFrameList,
		"selected_video_frame_list":      timeList,
	})
}
