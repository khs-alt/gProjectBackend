package main

import (
	"backend/sql"
	"backend/util"
	"fmt"
	"testing"
)

func TestDB(t *testing.T) {
	sql.DeleteDBTable()
	sql.CreateDBTable()
	util.DeleteAllFilesInFolder("originalVideos")
	util.DeleteAllFilesInFolder("artifactVideos")
	util.DeleteAllFilesInFolder("diffVideos")
	util.DeleteAllFilesInFolder("originalImages")
	util.DeleteAllFilesInFolder("artifactImages")
	util.DeleteAllFilesInFolder("diffImages")
}
func TestGetCurrentUserScoreList(t *testing.T) {
	userScoreList := sql.GetCurrentUserScoreList("kim", []string{"1", "2", "3", "4"})
	fmt.Println(userScoreList)
}

func TestGetUserScoreList(t *testing.T) {
	a, b := sql.GetCurrentUserScore("kim", "KC0rtr1o")
	fmt.Println(a)
	fmt.Println(b)
}

func TestDeletImageDB(t *testing.T) {
	sql.DeleteImageDBTablbe()
	sql.CreateImageDBTalbe()
	util.DeleteAllFilesInFolder("originalImages")
	util.DeleteAllFilesInFolder("artifactImages")
	util.DeleteAllFilesInFolder("diffImages")
}

// func TestGetLabelingList(t *testing.T) {
// 	labelingList := sql.GetUserLabelingList("kim", []int{1, 2, 3, 4, 5})
// 	a, _ := util.ConvertTo2DIntSlice(labelingList)
// 	fmt.Print(a)
// }

func TestInsertVideoScoring(t *testing.T) {
	sql.InsertUserVideoScoringInfo("kim", 1, "KC0rtr1o", 1)
	sql.InsertUserVideoScoringInfo("kim", 2, "KC0rtr1o", 4)
	sql.InsertUserVideoScoringInfo("kim", 2, "KC0rtr1o", 3)
	sql.InsertUserVideoScoringInfo("kim", 2, "KC0rtr1o", 5)
}

func TestInsertImageScoring(t *testing.T) {
	sql.InsertUserImageScoringInfo("kim", 1, "KC0rtr1o", "0,0,0,0,0,0,0,0,0,0,1,0,1")
	sql.InsertUserImageScoringInfo("kim", 2, "KC0rtr1o", "0,0,0,0,0,0,0,1,1,1,1,1,1")
	sql.InsertUserImageScoringInfo("kim", 2, "KC0rtr1o", "0,0,0,0,0,0,0,2,2,2,2,2,2")
}

// func TestSelectFram(t *testing.T) {
// 	videoIndex := "1"
// 	VideoCurrentTime := 02.05
// 	videoFilePath := fmt.Sprintf("./artifactVideos/artifactVideo%s.mp4", videoIndex)
// 	videoCurrentTime := VideoCurrentTime
// 	outputImage := fmt.Sprintf("./selectedFrame/selectedFrame%s_%s.png", videoIndex, videoCurrentTime)
// 	err := util.ExtractFrame(videoFilePath, videoCurrentTime, outputImage)
// 	if err != nil {
// 		log.Println("error: ", err)
// 		return
// 	}
// }

func TestGetVideoListInfoFromTestCode(t *testing.T) {
	sql.GetVideoListInfoFromTestCode("9jJ0A2mP")
}

// func TestGetSelectedFrameList(t *testing.T) {
// 	s := sql.GetSelectedFrameList(1)
// 	fmt.Println(s)
// }

func TestRemoveSpecificPart(t *testing.T) {
	s := "original_amusement_park2_p64_t0.3_n0.mp4"
	s1 := util.RemoveSpecificPart(s)
	fmt.Println(s1)
}
