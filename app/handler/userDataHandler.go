package handler

import (
	"backend/app/models"
	"backend/sql"
	"backend/util"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// login handler
// done
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

func GetUserScoringListHandler(c *gin.Context) {
	currentUser := c.Query("userID")
	testCode := c.Query("testcode")
	_, _, _, videoIndex, _ := sql.GetVideoListFromTestCode(testCode)

	userScoreList := sql.GetCurrentUserScoreList(currentUser, videoIndex)
	//TODO: videoIndex를 이용해서 userScoringList를 만들어서 보내주기

	userStringList := make([]string, len(userScoreList))
	for i, val := range userScoreList {
		userStringList[i] = strconv.Itoa(val)
	}
	randVideoIndex := util.ShuffleList(currentUser, userStringList)
	c.JSON(http.StatusOK, gin.H{
		"userScoringList": randVideoIndex,
	})
}

func GetImageScoreData(c *gin.Context) {
	var data models.UserImageScoreData
	if err := c.ShouldBindJSON(&data); err != nil {
		log.Println("GetImageScoreData error ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	uuid := util.MakeUUID()
	currentPage := data.ImageId
	SCVData := util.MakeIntListtoCSV(data.Score)
	isVideo := false
	go sql.InsertUserImageScoringInfo(data.CurrentUser, data.ImageId, data.TestCode, SCVData)
	go sql.InsertUserTestInfo(uuid, data.CurrentUser, data.TestCode, currentPage, isVideo)
	c.String(http.StatusOK, "Success insert user image score data")
}

func GetVideoScoringData(c *gin.Context) {
	var data models.UserScoreData
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	uuid := util.MakeUUID()
	isVideo := true
	go sql.InsertUserVideoScoringInfo(data.CurrentUser, data.ImageId, data.TestCode, data.Score)
	go sql.InsertUserTestInfo(uuid, data.CurrentUser, data.TestCode, data.ImageId, isVideo)
	_, userScore := sql.GetCurrentUserScore(data.CurrentUser, data.TestCode)

	c.JSON(http.StatusOK, userScore)
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

// done
func RequestLoginHandler(c *gin.Context) {
	var data models.UserLoginData
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}
	//check login process
	if !sql.IsUserIdExist(data.ID, data.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid ID or Password"})
		return
	}

	isVideoTestcodeExist, err := sql.GetVideoTestcodeExist(data.TestCode)
	if err != nil {
		log.Println("GetVideoTestcodeExist error")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
	}

	if isVideoTestcodeExist {
		var userIndex int
		userIndex, _ = sql.GetCurrentUserScore(data.ID, data.TestCode)
		c.JSON(http.StatusOK, gin.H{
			"path":     "scoring",
			"lastPage": userIndex,
		})
		return
	}

	isImageTestcodeExist, err := sql.GetImageTestcodeExist(data.TestCode)
	if err != nil {
		log.Println("GetImageTestcodeExist error")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
	}

	if isImageTestcodeExist {
		userIndex := sql.GetUserCurrentImagePageAboutTestCode(data.ID, data.TestCode)
		c.JSON(http.StatusOK, gin.H{
			"path":     "labeling",
			"lastPage": userIndex,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"path": "No Testcode",
	})
}

func GetUserLabelingListHandler(c *gin.Context) {
	var data models.UserLabelingListData
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}
	imageList, err := sql.GetImageListFromTestCode(data.TestCode)
	if err != nil {
		log.Println("GetImageListFromTestCode error: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
	}

	randImageList := util.ShuffleList(data.UserID, imageList)
	var intRandImageList []int
	for _, imageIndex := range randImageList {
		intImageIndex, _ := strconv.Atoi(imageIndex)
		intRandImageList = append(intRandImageList, intImageIndex)
	}
	userLabelingList := sql.GetUserLabelingList(data.UserID, intRandImageList)
	userLabelingIntList, err := util.ConvertTo2DIntSlice(userLabelingList)
	if err != nil {
		log.Println("ConvertTo2DIntSlice error: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
	}
	c.JSON(http.StatusOK, gin.H{
		"userLabelingList": userLabelingIntList,
	})
}

// func GetUserScoringListHandler(c *gin.Context) {
// 	var data models.UserScoringListData
// 	if err := c.ShouldBindJSON(&data); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	userScoringList := sql.GetUserScoringList(data.CurrentUser, data.TestCode)

// 	c.JSON(http.StatusOK, gin.H{
// 		"userScoringList": userScoringList,
// 	})
// }
