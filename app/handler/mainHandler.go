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

func GetUserCurrentImagePage(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		util.EnableCorsResponse(&w)
	}
	if r.Method == http.MethodPost {
		fmt.Println("SessionAuthMiddleware")
		session, err := util.Store.Get(r, "surveySession")
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		userId := session.Values["authenticated"]
		fmt.Println(session.IsNew, userId)
		if session.IsNew || userId != "true" {
			fmt.Println("=============")
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
		util.EnableCors(&w)
		body, _ := util.ProcessRequest(w, r)

		var data map[string]interface{}
		err = json.Unmarshal(body, &data)
		if err != nil {
			http.Error(w, "Error decoding JSON data", http.StatusBadRequest)
			return
		}
		id := data["userID"].(string)
		testCode := data["testcode"].(string)
		fmt.Println(id, testCode)
		currentPage := sql.GetUserCurrentImagePageAboutTestCode(id, testCode)
		imageCSVList, err := sql.GetImageListFromTestCode(testCode)
		imageList := util.MakeCSVToStringList(imageCSVList)

		data1 := struct {
			CurrentPage int      `json:"current_page"`
			ImageList   []string `json:"image_list"`
		}{
			CurrentPage: currentPage,
			ImageList:   imageList,
		}
		w.WriteHeader(http.StatusOK)
		jsonData, err := json.Marshal(data1)
		w.Write(jsonData)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
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

		requestData := struct {
			CurrentPage string `json:"current_page"`
		}{
			CurrentPage: currentPage,
		}
		w.WriteHeader(http.StatusOK)
		jsonData, err := json.Marshal(requestData)
		if err != nil {
			log.Println(err)
		}
		w.Write(jsonData)
	}
}

func GetUserCurrentPageInfo(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		util.EnableCorsResponse(&w)
	}
	if r.Method == http.MethodPost {
		fmt.Println("SessionAuthMiddleware")
		session, err := util.Store.Get(r, "surveySession")
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		userId := session.Values["authenticated"]
		fmt.Println(session.IsNew, userId)
		if session.IsNew || userId != "true" {
			fmt.Println("=============")
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
		util.EnableCors(&w)
		body, _ := util.ProcessRequest(w, r)

		var data map[string]interface{}
		err = json.Unmarshal(body, &data)
		if err != nil {
			http.Error(w, "Error decoding JSON data", http.StatusBadRequest)
			return
		}

		// 받은 데이터 출력
		id := data["userID"].(string)
		testCode := data["testcode"].(string)
		currentPage := data["currentPage"].(string)
		fmt.Println(id, testCode)
		//currentPage := fmt.Sprint(sql.GetUserCurrentPageAboutTestCode(id, testCode))
		videoCSVList, err := sql.GetVideoListFromTestCode(testCode)
		var videoNumList []string
		var originalVideoNameList []string
		var artifactVideoNameList []string
		var originalVideoFPSList []string
		var artifactVideoFPSList []string
		videoList := strings.Split(videoCSVList, ",")
		originalVideoNameList, artifactVideoNameList = sql.GetVideoNameListFromVideoList(videoList)
		for _, video := range videoList {
			num := strings.TrimLeft(video, "originalVideo")
			videoNumList = append(videoNumList, num)
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
		//fmt.Println(currentPage, videoNumCSVList)

		var initData models.UserVideoInitInfo
		initData.CurrentPage = fmt.Sprint(currentPage)
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
