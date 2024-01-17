package handler

import (
	"backend/sql"
	"backend/util"
	"fmt"
	"log"
	"net/http"
	"sort"
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
	originalImagesName, oriImages, originalImageFileForm := util.UploadImage(c, originalImages, "original")

	artifactImages := form.File["artifact"]
	artifactImagesName, artiImages, _ := util.UploadImage(c, artifactImages, "artifact")

	diffImages := form.File["diff"]
	diffImagesName, diffImgs, _ := util.UploadImage(c, diffImages, "diff")

	// 모든 이미지 처리 후
	tags = strings.Split(tags[0], ",")
	sort.Strings(tags)
	for i := 0; i < len(oriImages) && i < len(artiImages) && i < len(diffImgs); i++ {
		uuid := util.MakeUUID()
		width, height, err := util.GetFileDimensions("./originalImages/" + oriImages[i] + originalImageFileForm[i])
		if err != nil {
			log.Println(err)
		}
		if err := sql.InsertImage(uuid, originalImagesName[i], artifactImagesName[i], diffImagesName[i], width, height); err != nil {
			log.Println(err)
		}
		for _, tag := range tags {

			err = sql.InsertImageTagLink(uuid, tag)
			if err != nil {
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
	userID := data["user_id"].(string)
	imageList, _ := sql.GetImageListFromTestCode(testcode)

	// var indexList []int
	// for _, image := range imageList {
	// 	id := strings.TrimLeft(image, "originalImage")
	// 	num, _ := strconv.Atoi(id)
	// 	indexList = append(indexList, num)
	// }

	imageOriginalList, imageArtifactList := sql.GetImageNameListFromVideoList(imageList)

	randImageList := util.ShuffleList(userID, imageList)
	randImageOriginalList := util.ShuffleList(userID, imageOriginalList)
	randImageArtifactList := util.ShuffleList(userID, imageArtifactList)

	var imageOriginalList1 []string
	var imageArtifactList1 []string
	for _, image := range randImageOriginalList {
		imageOriginalList1 = append(imageOriginalList1, util.RemoveSpecificPart(image))
	}
	for _, image := range randImageArtifactList {
		imageArtifactList1 = append(imageArtifactList1, util.RemoveSpecificPart(image))
	}

	c.JSON(http.StatusOK, gin.H{
		"image_list":    randImageList,
		"original_list": imageOriginalList1,
		"artifact_list": imageArtifactList1,
	})
}
