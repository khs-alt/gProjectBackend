package main

import (
	"GoogleProjectBackend/sql"
	"fmt"
	"testing"
)

func TestGettagList(t *testing.T) {
	tagDataList := sql.GetTagData()
	fmt.Println(tagDataList)
}

func TestGetTescodeList(t *testing.T) {
	a, b := sql.GetTestCodeInfo()
	fmt.Println(a, b)
}
