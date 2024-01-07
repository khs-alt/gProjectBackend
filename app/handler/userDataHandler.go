package handler

import (
	"backend/app/models"
	"backend/sql"
	"backend/util"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func SignupHandler(c *gin.Context) {
	var data models.UserNewIdAndPassword
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newId := data.NewId
	newPassword := data.NewPassword
	uuid, _ := uuid.NewUUID()
	res := sql.InsertUserIdAndPassword(uuid, newId, newPassword)
	c.String(http.StatusOK, res)
}

func GetImageScoreDataFromUser(c *gin.Context) {
	var data models.UserImageInfo
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userScore := sql.GetCurrentUserImageScore(data.CurrentUser, data.ImageID)
	userIntScore := util.MakeCSVtoIntList(userScore)
	c.JSON(http.StatusOK, gin.H{
		"patch": userIntScore,
	})
}

func GetScoreDataFromUser(c *gin.Context) {
	var data models.UserInfoData
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userScore := sql.GetCurrentUserScore(data.CurrentUser, data.ImageId)
	replyData := fmt.Sprint(userScore)
	c.String(http.StatusOK, replyData)
}

func GetImageScoreData(c *gin.Context) {
	var data models.UserImageScoreData
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	uuid := util.MakeUUID()
	currentPage := data.ImageId
	SCVData := util.MakeIntListtoCSV(data.Score)
	go sql.InsertUserImageScoringInfo(uuid, data.CurrentUser, data.ImageId, SCVData)
	go sql.InsertUserImageTestInfo(uuid, data.CurrentUser, data.TestCode, currentPage)
	c.String(http.StatusOK, "Success insert user image score data")
}

func GetScoringData(c *gin.Context) {
	var data models.UserScoreData
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	uuid := util.MakeUUID()
	currentPage := data.ImageId
	go sql.InsertUserVideoScoringInfo(uuid, data.CurrentUser, data.ImageId, data.Score)
	go sql.InsertUserTestInfo(uuid, data.CurrentUser, data.TestCode, currentPage)
	userScore := sql.GetCurrentUserScore(data.CurrentUser, data.ImageId+1)
	var res models.UserCurrentScore
	res.Score = userScore
	c.String(http.StatusOK, "Success insert user score data")
}

func AdminLoginHandler(c *gin.Context) {
	var data map[string]interface{}
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id := data["adminId"].(string)
	password := data["adminPassword"].(string)
	res := ""
	if id == "admin" && password == "c404b!pipi" {
		res = "yes"
	} else {
		res = "no"
	}
	c.String(http.StatusOK, res)
}

func ReqeustLoginHandler(c *gin.Context) {
	var data models.UserLoginData
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	IsUserIdExist := sql.IsUserIdExist(data.ID, data.Password)
	IsVideoTestcodeExist := sql.GetTestcodeExist(data.TestCode)
	var currentPage string
	IsImageTestcodeExist := sql.GetImageTestcodeExist(data.TestCode)
	var res string
	if IsVideoTestcodeExist == true {
		res = "scoring"
	}
	if IsImageTestcodeExist == true {
		res = "labeling"
	}
	if IsVideoTestcodeExist == false && IsImageTestcodeExist == false {
		res = "No TestCode"
	}
	if IsUserIdExist == false {
		res = "Wrong ID or Password"
	}
	c.JSON(http.StatusOK, gin.H{
		"path":     res,
		"lastPage": currentPage,
	})
}
