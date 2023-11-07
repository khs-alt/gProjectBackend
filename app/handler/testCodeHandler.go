package handler

import (
	"GoogleProjectBackend/app/models"
	"GoogleProjectBackend/sql"
	"GoogleProjectBackend/util"
	"encoding/json"
	"fmt"
	"sort"
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

func GetCSVVideoListFromTagHandler(w http.ResponseWriter, r *http.Request) {

}

func GetVideoListFromTagHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	util.EnableCors(&w)
	tags := r.URL.Query()["tag[]"]
	fmt.Print("tag data is: ")
	var videoList []string
	for _, tag := range tags {
		videos, _ := sql.GetVideoListFromTag(tag)
		for _, video := range videos {
			videoList = append(videoList, video)
		}
	}
	videoList = util.RemoveDuplicates(videoList) //중복된 비디오 리스트 제거
	originalVideo, _ := sql.GetVideoNameListFromVideoList(videoList)
	fmt.Println(originalVideo)
	jsonData, err := json.Marshal(originalVideo)
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
		//테그의 리스트를 받아서 각 테그에 해당하는 비디오를 가져옴
		//다만 중복이 발생할 수 있음
		//즉 videoList에 중복이 가능함
		for _, tag := range data.Tags {
			videos, _ := sql.GetVideoListFromTag(tag)
			for _, video := range videos {
				videoList = append(videoList, video)
			}
		}
		//중복된 요소를 제거함
		videoList = util.RemoveDuplicates(videoList)
		sort.Strings(videoList)
		//TODO: videoList를 sort해야함
		num, _ := sql.GetTestCodeCount()

		testcode := util.GenerateRandomString(8) + fmt.Sprint(num)
		sql.InsertTestCodeId(uuid, testcode, data.Tags, videoList)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(testcode))
	}
	//w.Write(jsonData)
}
