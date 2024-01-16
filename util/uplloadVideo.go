package util

import (
	"fmt"
	"log"
	"mime/multipart"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func UploadVideo(c *gin.Context, videos []*multipart.FileHeader, videoType string) ([]string, []string, []string) {
	var VideosName, Videos, VideosFileForm []string
	for _, VideoHeader := range videos {
		//들어온 비디오를 video+n 이름으로 바꿈
		VideosName = append(VideosName, filepath.Base(VideoHeader.Filename))
		VideoFileForm := filepath.Ext(VideoHeader.Filename)
		VideosFileForm = append(VideosFileForm, VideoFileForm)

		VideoPath := "./" + videoType + "Videos/"
		count, err := CountFile(VideoPath)
		if err != nil {
			fmt.Print("CountFile error : ")
			fmt.Println(err)
		}
		VideoName := videoType + "Video" + fmt.Sprint(count+1)
		Videos = append(Videos, VideoName)
		FilePath := VideoPath + VideoName + VideoFileForm
		if err := c.SaveUploadedFile(VideoHeader, FilePath); err != nil {
			log.Println("원본 비디오 저장 실패: ", err)
			continue
		}
	}
	return VideosName, Videos, VideosFileForm
}
