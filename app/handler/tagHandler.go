package handler

import (
	"backend/app/models"
	"backend/sql"
	"backend/util"
	"log"
	"net/http"
	"sort"

	"github.com/gin-gonic/gin"
)

// d
func DeleteVideotagHandler(c *gin.Context) {
	var data models.RequestData
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	tags := data.Tags
	for _, tag := range tags {
		err := sql.DeleteTagData(tag)
		if err != nil {
			log.Println(err)
		}
	}
	sort.Strings(tags)
	c.String(http.StatusOK, "Success delete tag")
}

// d
func DeleteImagetagHandler(c *gin.Context) {
	var data models.RequestData
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	tags := data.Tags
	for _, tag := range tags {
		err := sql.DeleteImageTagData(tag)
		if err != nil {
			log.Println(err)
		}
	}
	sort.Strings(tags)
	c.String(http.StatusOK, "Success delete tag")
}

// d
func ReceivedVideoTagHandler(c *gin.Context) {
	var data map[string]interface{}
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	tag := data["tag"].(string)
	uuid := util.MakeUUID()
	sql.InsertVideoTag(uuid, tag)
	c.String(http.StatusOK, "Success insert tag")
}

// d
func ReceivedImageTagHandler(c *gin.Context) {
	var data map[string]interface{}
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	tag := data["tag"].(string)
	uuid := util.MakeUUID()
	sql.InsertImageTag(uuid, tag)
	c.String(http.StatusOK, "Success insert image tag")
}

// TODO:테그를 삭제할 경우 response로 보내주는 값이 없는 것으로 예상됨.
// db에서 tag 데이터를 가져와서 json으로 변환 후 리스트 형태로 반환
// d
func GetVideoTagHandler(c *gin.Context) {
	tagDataList := sql.GetVideoTag()
	c.JSON(http.StatusOK, tagDataList)
}

func GetImageTagHandler(c *gin.Context) {
	tagDataList := sql.GetImageTag()
	c.JSON(http.StatusOK, tagDataList)
}
