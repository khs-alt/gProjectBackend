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
