package handler

import (
	"GoogleProjectBackend/sql"
	"GoogleProjectBackend/util"
	"encoding/json"
	"fmt"
	"net/http"
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
		videoList, err := sql.GetVideoListFromTestCode(testCode)
		if err != nil {
			fmt.Println(err)
		}
		resData := map[string]string{
			"currentPage": currentPage,
			"videoList":   videoList,
		}

		// 응답 보내기
		w.WriteHeader(http.StatusOK)
		jsonData, err := json.Marshal(resData)
		w.Write(jsonData)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
