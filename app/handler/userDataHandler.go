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

func SighupHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		util.EnableCorsResponse(&w)
	}
	if r.Method == http.MethodPost {
		util.EnableCors(&w)
		body, _ := util.ProcessRequest(w, r)

		var data models.UserNewIdAndPassword
		err := json.Unmarshal(body, &data)
		if err != nil {
			http.Error(w, "Error decoding JSON data", http.StatusBadRequest)
			return
		}
		fmt.Println("Received Data:", data)
		newId := data.NewId
		newPassword := data.NewPassword
		uuid, _ := uuid.NewUUID()
		res := sql.InsertUserIdAndPassword(uuid, newId, newPassword)
		fmt.Println(newId, newPassword, res)
		// 응답 보내기
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(res)) //여기가 데이터 보내는 곳
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func GetUserScore(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		util.EnableCorsResponse(&w)
	}
	if r.Method == http.MethodPost {
		util.EnableCors(&w)
		body, _ := util.ProcessRequest(w, r)
		var data models.UserInfoData
		err := json.Unmarshal(body, &data)
		if err != nil {
			http.Error(w, "Error decoding JSON data", http.StatusBadRequest)
			return
		}
	}
}

func GetScoreDataFromUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		util.EnableCorsResponse(&w)
	}
	if r.Method == http.MethodPost {
		util.EnableCors(&w)
		body, _ := util.ProcessRequest(w, r)

		var data models.UserInfoData
		err := json.Unmarshal(body, &data)
		if err != nil {
			http.Error(w, "Error decoding JSON data", http.StatusBadRequest)
			return
		}
		fmt.Println("Received Data:", data)
		userScore := sql.GetCurrentUserScore(data.CurrentUser, data.ImageId)
		var res models.UserCurrentScore
		res.Score = userScore

		// JSON으로 응답 데이터 마샬링
		// jsonResponse, err := json.Marshal(res)
		// if err != nil {
		// 	http.Error(w, err.Error(), http.StatusInternalServerError)
		// 	return
		// }

		// Content-Type 설정 및 JSON 데이터 전송
		//w.Header().Set("Content-Type", "application/json")
		//w.Write(jsonResponse)
		// 응답 보내기
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprint(userScore)))
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func GetScoringData(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		util.EnableCorsResponse(&w)
	}
	if r.Method == http.MethodPost {
		util.EnableCors(&w)
		body, _ := util.ProcessRequest(w, r)

		var data models.UserScoreData
		err := json.Unmarshal(body, &data)
		if err != nil {
			http.Error(w, "Error decoding JSON data", http.StatusBadRequest)
			return
		}
		if data.Title == "scoring data" {

			fmt.Print("scoring data part::::")

			uuid := util.MakeUUID()
			currentPage := data.ImageId
			fmt.Println(currentPage)
			sql.InsertUserVideoScoringInfo(uuid, data.CurrentUser, data.ImageId, data.Score)
			sql.InsertUserTestInfo(uuid, data.CurrentUser, data.TestCode, currentPage)
			userScore := sql.GetCurrentUserScore(data.CurrentUser, data.ImageId)
			var res models.UserCurrentScore
			res.Score = userScore
			fmt.Println(res.Score)
			// JSON으로 응답 데이터 마샬링
			//jsonResponse, err := json.Marshal(res)
			// if err != nil {
			// 	http.Error(w, err.Error(), http.StatusInternalServerError)
			// 	return
			// }

			// Content-Type 설정 및 JSON 데이터 전송
			//w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(fmt.Sprint(userScore)))
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}

func AdminLoginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
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
		// 응답 보내기
		fmt.Println("Received Data:", data)
		id := data["adminId"].(string)
		password := data["adminPassword"].(string)
		res := ""
		if id == "admin" && password == "admin" {
			res = "yes"
		} else {
			res = "no"
		}
		// 응답 보내기
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(res))
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// 어떤 유저의 테스트코드에 따른 비디오 리스트와 현재 페이지를 반환 및 로그인
func ReqeustLoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		util.EnableCorsResponse(&w)
	}
	if r.Method == http.MethodPost {
		util.EnableCors(&w)
		body, _ := util.ProcessRequest(w, r)

		var data models.UserLoginData
		err := json.Unmarshal(body, &data)
		if err != nil {
			http.Error(w, "Error decoding JSON data", http.StatusBadRequest)
			return
		}
		//
		IsUserIdExist := sql.IsUserIdExist(data.ID, data.Password)
		IsTestcodeExist := sql.GetTestcodeExist(data.TestCode)

		var res string
		if IsTestcodeExist != true {
			w.WriteHeader(http.StatusOK)
			res = "No TestCode"
		} else if IsUserIdExist != true {
			res = "No UserId"

		} else {
			w.WriteHeader(http.StatusOK)
			res = "Yes"
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(res))
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
