package handler

import (
	"GoogleProjectBackend/app/models"
	"GoogleProjectBackend/sql"
	"GoogleProjectBackend/util"
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
)

func DeletetagHandler(w http.ResponseWriter, r *http.Request) {
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

		// 받은 데이터 출력
		fmt.Println("Received Data:", data)
		fmt.Println(data.Tags)
		tags := data.Tags
		for _, tag := range tags {
			err := sql.DeleteTagData(tag)
			if err != nil {
				fmt.Println(err)
			}
		}
		sort.Strings(tags)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Success delete tag"))
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func ReceivedTagHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		util.EnableCorsResponse(&w)
	}
	if r.Method == http.MethodPost {
		util.EnableCors(&w)
		body, _ := util.ProcessRequest(w, r)

		// JSON 데이터 디코딩
		var data map[string]interface{}
		err := json.Unmarshal(body, &data)
		if err != nil {
			http.Error(w, "Error decoding JSON data", http.StatusBadRequest)
			fmt.Printf("marshal err: %s\n", err)
			return
		}
		// 받은 데이터 출력
		fmt.Println("Received Data:", data)
		tag := data["tag"].(string)
		println(tag)
		// 응답 보내기
		uuid := util.MakeUUID()
		sql.InsertTagData(uuid, tag)
		fmt.Println("-----------")
		w.Write([]byte("Success")) //여기가 데이터 보내는 곳
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// TODO:테그를 삭제할 경우 response로 보내주는 값이 없는 것으로 예상됨.
// db에서 tag 데이터를 가져와서 json으로 변환 후 리스트 형태로 반환
func GetTagHandler(w http.ResponseWriter, r *http.Request) {
	util.EnableCors(&w)
	tagDataList := sql.GetTagData()
	jsonData, err := json.Marshal(tagDataList)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println("json marshal error")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}
