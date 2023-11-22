package main

import (
	"GoogleProjectBackend/sql"
	"GoogleProjectBackend/util"
	"fmt"
	"testing"

	"github.com/google/uuid"
)

func TestDeletTable(t *testing.T) {
	sql.DeleteDBTablbe()
}
func TestCreateTable(t *testing.T) {
	sql.CreateDBTalbe()
}
func TestInsertUserInfo(t *testing.T) {
	id := "kim"
	password := "123"
	uuid, _ := uuid.NewUUID()
	res := sql.InsertUserIdAndPassword(uuid, id, password)
	fmt.Println(res)
}
func TestInputTag1(t *testing.T) {
	uuid, _ := uuid.NewUUID()
	sql.InsertTagData(uuid, "dark")
}
func TestInputTag2(t *testing.T) {
	uuid, _ := uuid.NewUUID()
	sql.InsertTagData(uuid, "bright")
}

func TestMakeTestcode1(t *testing.T) {
	uuid, _ := uuid.NewUUID()
	tags := []string{"bright", "dark"}
	var videoList []string
	for _, tag := range tags {
		videos, _ := sql.GetVideoListFromTag(tag)
		for _, video := range videos {
			videoList = append(videoList, video)
		}
	}
	num, _ := sql.GetTestCodeCount()
	testcode := "A" + fmt.Sprint(num)
	sql.InsertTestCodeId(uuid, testcode, tags, videoList)
}

func TestMakeTestcode2(t *testing.T) {
	uuid, _ := uuid.NewUUID()
	tags := []string{"dark"}
	var videoList []string
	for _, tag := range tags {
		videos, _ := sql.GetVideoListFromTag(tag)
		for _, video := range videos {
			videoList = append(videoList, video)
		}
	}
	num, _ := sql.GetTestCodeCount()
	testcode := "A" + fmt.Sprint(num)
	sql.InsertTestCodeId(uuid, testcode, tags, videoList)
}
func TestMakeTestcode3(t *testing.T) {
	uuid, _ := uuid.NewUUID()
	tags := []string{"bright"}
	var videoList []string
	for _, tag := range tags {
		videos, _ := sql.GetVideoListFromTag(tag)
		for _, video := range videos {
			videoList = append(videoList, video)
		}
	}
	num, _ := sql.GetTestCodeCount()
	testcode := "A" + fmt.Sprint(num)
	sql.InsertTestCodeId(uuid, testcode, tags, videoList)
}
func TestGetVideoList(t *testing.T) {
	testCode := "A1"
	videoCSV, _ := sql.GetVideoListFromTestCode(testCode)
	videoList := util.MakeCSVToStringList(videoCSV)
	fmt.Println(videoList)
}

func TestGetUserTestcodeCurrentPage(t *testing.T) {
	id := "kim"
	testCode := "A3"
	currentPage := sql.GetUserCurrentPageAboutTestCode(id, testCode)
	fmt.Printf("User %s current page is %d\n", id, currentPage)
}

func TestInsertUserVideoScoring1(t *testing.T) {
	id := "kim"
	testCode := "A3"
	userScore := 10
	videoId := 1
	currentPage := 0
	uuid1, _ := uuid.NewUUID()
	sql.InsertUserVideoScoringInfo(uuid1, id, videoId, userScore)
	uuid2, _ := uuid.NewUUID()
	sql.InsertUserTestInfo(uuid2, id, testCode, currentPage)
}

func TestInsertUserVideoScoring2(t *testing.T) {
	id := "kim"
	testCode := "A3"
	userScore := 8
	videoId := 12
	currentPage := 1
	uuid1, _ := uuid.NewUUID()
	sql.InsertUserVideoScoringInfo(uuid1, id, videoId, userScore)
	uuid2, _ := uuid.NewUUID()
	sql.InsertUserTestInfo(uuid2, id, testCode, currentPage)
}

func TestInsertUserVideoScoring3(t *testing.T) {
	id := "kim"
	testCode := "A3"
	userScore := 7
	videoId := 16
	currentPage := 2
	uuid1, _ := uuid.NewUUID()
	sql.InsertUserVideoScoringInfo(uuid1, id, videoId, userScore)
	uuid2, _ := uuid.NewUUID()
	sql.InsertUserTestInfo(uuid2, id, testCode, currentPage)
}

func TestGetUserTestcodeCurrentPage1(t *testing.T) {
	id := "lee"
	testCode := "2pTIsuT62"
	currentPage := sql.GetUserCurrentPageAboutTestCode(id, testCode)
	fmt.Printf("User %s current page is %d\n", id, currentPage)
}

func TestGetUserScore(t *testing.T) {
	sql.DeleteDBTablbe()
	sql.CreateDBTalbe()
	util.DeleteAllFilesInFolder("originalVideos")
	util.DeleteAllFilesInFolder("artifactVideos")
	util.DeleteAllFilesInFolder("originalImages")
	util.DeleteAllFilesInFolder("artifactImages")
}
