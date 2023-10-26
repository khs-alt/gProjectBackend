package handler

import (
	"GoogleProjectBackend/app/models"
	"GoogleProjectBackend/sql"
	"GoogleProjectBackend/util"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

func GetTestCodeListHandler(w http.ResponseWriter, r *http.Request) {
	util.EnableCors(&w)
	testCodeList, tagsList := sql.GetTestCodeInfo()
	fmt.Print(tagsList)
	ressult := struct {
		TestCode []string `json:"testcode"`
		Tags     []string `json:"tags"`
	}{
		TestCode: testCodeList,
		Tags:     tagsList,
	}
	jsonData, err := json.Marshal(ressult)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println("json marshal error")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func GetTestCodeHandler(w http.ResponseWriter, r *http.Request) {
	// util.EnableCors(&w)
	// //TODO: 이거 뭐임?
	// //testCode := r.URL.Query().Get("testcode")
	// testCodeInfo := sql.GetTestCode(testCode)
	// jsonData, err := json.Marshal(testCodeInfo)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	fmt.Println("json marshal error")
	// 	return
	// }

	if r.Method == http.MethodOptions {
		util.EnableCorsResponse(&w)
	}
	if r.Method == http.MethodPost {
		util.EnableCors(&w)
		body, _ := util.ProcessRequest(w, r)
		var data models.RequestData
		err := json.Unmarshal(body, &data)
		if err != nil {
			http.Error(w, "Error decoding JSON data", http.StatusBadRequest)
			return
		}
		uuid, _ := uuid.NewUUID()
		var videoList []string
		for _, tag := range data.Tags {
			videos, _ := sql.GetVideo(tag)
			for _, video := range videos {
				videoList = append(videoList, video)
			}
		}
		num, _ := sql.GetTestCodeCount()
		testcode := "Test" + fmt.Sprint(num)
		sql.InsertTestCodeId(uuid, testcode, data.Tags, videoList)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(testcode))

	}

	//w.Write(jsonData)
}
