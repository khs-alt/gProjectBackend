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

func TestGetLabelingList(t *testing.T) {
	labelingList := sql.GetUserLabelingList("kim", []int{1, 2, 3, 4, 5})
	a, _ := util.ConvertTo2DIntSlice(labelingList)
	fmt.Print(a)
}

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
