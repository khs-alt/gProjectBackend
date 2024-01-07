package handler

import (
	"backend/sql"
	"backend/util"
	"log"
	"net/http"
	"sort"

	"github.com/gin-gonic/gin"
)

type RequestData struct {
	Tags []string `json:"tags"`
}

func DeletetagHandler(c *gin.Context) {
	var data RequestData
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

func DeleteImagetagHandler(c *gin.Context) {
	var data RequestData
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

func ReceivedTagHandler(c *gin.Context) {
	var data map[string]interface{}
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	tag := data["tag"].(string)
	uuid := util.MakeUUID()
	sql.InsertTagData(uuid, tag)
	c.String(http.StatusOK, "Success insert tag")
}

func ReceivedImageTagHandler(c *gin.Context) {
	var data map[string]interface{}
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	tag := data["tag"].(string)
	// 응답 보내기
	uuid := util.MakeUUID()
	sql.InsertImageTagData(uuid, tag)
	c.String(http.StatusOK, "Success insert image tag")
}

// TODO:테그를 삭제할 경우 response로 보내주는 값이 없는 것으로 예상됨.
// db에서 tag 데이터를 가져와서 json으로 변환 후 리스트 형태로 반환
func GetTagHandler(c *gin.Context) {
	tagDataList := sql.GetTagData()
	c.JSON(http.StatusOK, gin.H{
		"testcode_list": tagDataList,
	})
}

func GetImageTagHandler(c *gin.Context) {
	tagDataList := sql.GetImageTagData()
	c.JSON(http.StatusOK, gin.H{
		"image_testcode_list": tagDataList,
	})
}
