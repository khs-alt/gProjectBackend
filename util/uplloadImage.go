package util

import (
	"fmt"
	"log"
	"mime/multipart"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func UploadImage(c *gin.Context, images []*multipart.FileHeader, imageType string) ([]string, []string) {
	var ImagesName, Images, VideosFileForm []string
	for _, ImageHeader := range images {
		//들어온 이미지를 video+n 이름으로 바꿈
		ImagesName = append(ImagesName, filepath.Base(ImageHeader.Filename))
		ImageFileForm := filepath.Ext(ImageHeader.Filename)
		VideosFileForm = append(VideosFileForm, ImageFileForm)

		ImagePath := "./" + imageType + "Images/"
		count, err := CountFile(ImagePath)
		if err != nil {
			fmt.Print("CountFile error : ")
			fmt.Println(err)
		}
		ImageName := imageType + "Image" + fmt.Sprint(count)
		Images = append(Images, ImageName)
		FilePath := ImagePath + ImageName + ImageFileForm
		if err := c.SaveUploadedFile(ImageHeader, FilePath); err != nil {
			log.Println("원본 이미지 저장 실패: ", err)
			continue
		}
	}
	return ImagesName, Images
}
