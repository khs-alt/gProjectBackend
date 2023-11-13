package handler

import (
	"GoogleProjectBackend/app/models"
	"GoogleProjectBackend/sql"
	"GoogleProjectBackend/util"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func MainHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello this is main Page!!!!")
}

func GetUserCurrentPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		util.EnableCorsResponse(&w)
	}
	if r.Method == http.MethodPost {
		util.EnableCors(&w)
		body, _ := util.ProcessRequest(w, r)

		var data map[string]interface{}
		err := json.Unmarshal(body, &data)
		if err != nil {
			http.Error(w, "Error decoding JSON data", http.StatusBadRequest)
			return
		}

		// 받은 데이터 출력
		id := data["userID"].(string)
		testCode := data["testcode"].(string)
		fmt.Println(id, testCode)
		currentPage := fmt.Sprint(sql.GetUserCurrentPageAboutTestCode(id, testCode))
		videoCSVList, err := sql.GetVideoListFromTestCode(testCode)
		var videoNumList []string
		var originalVideoNameList []string
		var artifactVideoNameList []string
		var originalVideoFPSList []string
		var artifactVideoFPSList []string
		videoList := strings.Split(videoCSVList, ",")
		originalVideoNameList, artifactVideoNameList = sql.GetVideoNameListFromVideoList(videoList)
		for _, video := range videoList {
			videoNumList = append(videoNumList, string(video[len(video)-1]))
		}
		for _, video := range videoList {
			originalVideoFPS, artifactVideoFPS := sql.GetFPSFromVideo(video)
			originalVideoFPSList = append(originalVideoFPSList, strconv.FormatFloat(float64(originalVideoFPS), 'f', 2, 32))
			artifactVideoFPSList = append(artifactVideoFPSList, strconv.FormatFloat(float64(artifactVideoFPS), 'f', 2, 32))
		}
		curPage, err := strconv.Atoi(currentPage)
		if err != nil {
			log.Println(err)
		}
		userScore := sql.GetCurrentUserScore(id, curPage)
		// videoNumCSVList := util.MakeStringListtoCSV(videoNumList)
		// originalVideoCSVList := util.MakeStringListtoCSV(originalVideoNameList)
		// artifactVideoCSVList := util.MakeStringListtoCSV(artifactVideoNameList)
		// originalVideoFPSCSVList := util.MakeStringListtoCSV(originalVideoFPSList)
		// artifactVideoFPSCSVList := util.MakeStringListtoCSV(artifactVideoFPSList)

		if err != nil {
			log.Println(err)
		}
		//fmt.Println(currentPage, videoNumCSVList)
		var initData models.UserVideoInitInfo
		initData.CurrentPage = currentPage
		initData.VideoList = videoNumList
		initData.OriginalVideoNameList = originalVideoNameList
		initData.ArtifactVideoNameList = artifactVideoNameList
		initData.OriginalVideoFPSList = originalVideoFPSList
		initData.ArtifactVideoFPSList = artifactVideoFPSList
		initData.UserScore = userScore

		// 응답 보내기
		w.WriteHeader(http.StatusOK)
		jsonData, err := json.Marshal(initData)
		w.Write(jsonData)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
