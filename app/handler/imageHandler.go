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

func ServeImage(c *gin.Context) {
	videoFilePath := "./image/google_logo.png"
	c.File(videoFilePath) // Serve file using Gin
}

func ServeOriginalImagesHandler(c *gin.Context) {
	ImageID := c.Param("id")
	imageFilePath := fmt.Sprintf("./originalImages/originalImage%s.png", ImageID)
	c.File(imageFilePath) // Serve file using Gin
}

func ServeArtifactImagesHandler(c *gin.Context) {
	ImageID := c.Param("id")
	imageFilePath := fmt.Sprintf("./artifactImages/artifactImage%s.png", ImageID)
	c.File(imageFilePath) // Serve file using Gin
}

func ServeDiffImagesHandler(c *gin.Context) {
	ImageID := c.Param("id")
	imageFilePath := fmt.Sprintf("./diffImages/diffImage%s.png", ImageID)
	c.File(imageFilePath) // Serve file using Gin
}

func UploadImageHandler(c *gin.Context) {

	err := c.Request.ParseMultipartForm(5000 << 20) // 5000MB
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
	originalImages := form.File["original"]
	originalImagesName, oriImages := util.UploadImage(c, originalImages, "original")

	artifactImages := form.File["artifact"]
	artifactImagesName, artiImages := util.UploadImage(c, artifactImages, "artifact")

	diffImages := form.File["diff"]
	diffImagesName, diffImgs := util.UploadImage(c, diffImages, "diff")

	// 모든 이미지 처리 후
	tags = strings.Split(tags[0], ",")
	for i := 0; i < len(oriImages) && i < len(artiImages); i++ {
		for _, tag := range tags {
			uuid := util.MakeUUID()
			if err := sql.InsertImageId(uuid, originalImagesName[i], oriImages[i], artifactImagesName[i], artiImages[i], diffImagesName[i], diffImgs[i], tag); err != nil {
				log.Println(err)
			}
		}
	}

	c.String(http.StatusOK, "이미지가 성공적으로 업로드되었습니다.")
}

func GetImageNameListHandler(c *gin.Context) {
	var data map[string]interface{}
	if err := c.ShouldBind(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	testcode := data["testcode"].(string)
	imageOriginalCSVList, _ := sql.GetImageListFromTestCode(testcode)

	imageList := util.MakeCSVToStringList(imageOriginalCSVList)
	var indexList []int
	for _, image := range imageList {
		id := strings.TrimLeft(image, "originalImage")
		num, _ := strconv.Atoi(id)
		indexList = append(indexList, num)
	}

	imageOriginalList, imageArtifactList := sql.GetImageNameListFromVideoList(imageList)
	c.JSON(http.StatusOK, gin.H{
		"ImageList":         imageList,
		"ImageOriginalList": imageOriginalList,
		"ImageArtifactList": imageArtifactList,
	})
}
