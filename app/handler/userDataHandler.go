package handler

import (
	"GoogleProjectBackend/app/models"
	"GoogleProjectBackend/sql"
	"GoogleProjectBackend/util"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/sessions"
)

func SignupHandler(w http.ResponseWriter, r *http.Request) {
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

func GetImageScoreDataFromUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		util.EnableCorsResponse(&w)
	}
	if r.Method == http.MethodPost {
		util.EnableCors(&w)
		body, _ := util.ProcessRequest(w, r)
		//axios로 front에서 받을 데이터
		data := struct {
			CurrentUser string `json:"current_user"`
			ImageID     int    `json:"image_id"`
		}{
			CurrentUser: "",
			ImageID:     0,
		}
		err := json.Unmarshal(body, &data)
		if err != nil {
			http.Error(w, "Error decoding JSON data", http.StatusBadRequest)
			return
		}
		userScore := sql.GetCurrentUserImageScore(data.CurrentUser, data.ImageID)
		userIntScore := util.MakeCSVtoIntList(userScore)
		replyData := struct {
			Patch []int `json:"patch"`
		}{
			Patch: userIntScore,
		}
		jsonData, err := json.Marshal(replyData)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(jsonData)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
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
		//w.Header().Set("Content-Type", "application/json")
		//w.Write(jsonResponse)
		// 응답 보내기
		fmt.Printf("userScore is %d\n", userScore)
		w.WriteHeader(http.StatusOK)
		replyData := fmt.Sprint(userScore)
		w.Write([]byte(replyData))
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func GetImageScoreData(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		util.EnableCorsResponse(&w)
	}
	if r.Method == http.MethodPost {
		util.EnableCors(&w)
		body, _ := util.ProcessRequest(w, r)

		var data models.UserImageScoreData
		err := json.Unmarshal(body, &data)
		if err != nil {
			http.Error(w, "Error decoding JSON data", http.StatusBadRequest)
			return
		}
		uuid := util.MakeUUID()
		currentPage := data.ImageId
		SCVData := util.MakeIntListtoCSV(data.Score)
		go sql.InsertUserImageScoringInfo(uuid, data.CurrentUser, data.ImageId, SCVData)
		go sql.InsertUserImageTestInfo(uuid, data.CurrentUser, data.TestCode, currentPage)
		// userScore := sql.GetCurrentUserImageScore(data.CurrentUser, data.ImageId+1)
		// userIntScore := util.MakeCSVtoIntList(userScore)
		// sendData := struct {
		// 	Score []int `json:"score"`
		// }{
		// 	Score: userIntScore,
		// }
		// finalSendData, err := json.Marshal(sendData)
		// if err != nil {
		// 	http.Error(w, err.Error(), http.StatusInternalServerError)
		// 	return
		// }
		// //json으로 만들어야 함.
		// w.Write(finalSendData)
		w.Write([]byte("Success"))
		w.WriteHeader(http.StatusOK)
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

			uuid := util.MakeUUID()
			currentPage := data.ImageId
			fmt.Printf("user %s videoId %d score %d\n", data.CurrentUser, data.ImageId, data.Score)
			go sql.InsertUserVideoScoringInfo(uuid, data.CurrentUser, data.ImageId, data.Score)
			go sql.InsertUserTestInfo(uuid, data.CurrentUser, data.TestCode, currentPage)
			userScore := sql.GetCurrentUserScore(data.CurrentUser, data.ImageId+1)
			var res models.UserCurrentScore
			res.Score = userScore
			// JSON으로 응답 데이터 마샬링
			//jsonResponse, err := json.Marshal(res)
			// if err != nil {
			// 	http.Error(w, err.Error(), http.StatusInternalServerError)
			// 	return
			// }

			// Content-Type 설정 및 JSON 데이터 전송
			//w.Header().Set("Content-Type", "application/json")
			fmt.Printf("user %s videoId %d currentScore %d\n", data.CurrentUser, data.ImageId, res.Score)
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
		if id == "admin" && password == "c404b!pipi" {
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
		fmt.Println(data)
		IsUserIdExist := sql.IsUserIdExist(data.ID, data.Password)

		IsVideoTestcodeExist := sql.GetTestcodeExist(data.TestCode)
		var currentPage string
		IsImageTestcodeExist := sql.GetImageTestcodeExist(data.TestCode)
		fmt.Println(IsUserIdExist, IsVideoTestcodeExist, IsImageTestcodeExist)
		var res string
		if IsVideoTestcodeExist == true {
			currentPage = fmt.Sprint(sql.GetUserCurrentPageAboutTestCode(data.ID, data.TestCode))
			session, _ := util.Store.Get(r, "surveySession")
			session.Values["authenticated"] = "true"
			session.Options = &sessions.Options{
				MaxAge: 1800, // 초 단위
			}
			session.Save(r, w)
			res = "scoring"
		}
		if IsImageTestcodeExist == true {
			currentPage = fmt.Sprint(sql.GetUserCurrentImagePageAboutTestCode(data.ID, data.TestCode))
			session, _ := util.Store.Get(r, "surveySession")
			session.Values["authenticated"] = "true"
			session.Options = &sessions.Options{
				MaxAge: 1800, // 초 단위
			}
			session.Save(r, w)
			res = "labeling"
		}
		if IsVideoTestcodeExist == false && IsImageTestcodeExist == false {
			res = "No TestCode"
		}
		if IsUserIdExist == false {
			res = "Wrong ID or Password"
		}

		resData := struct {
			Path     string `json:"path"`
			LastPage string `json:"lastPage"`
		}{
			Path:     res,
			LastPage: currentPage,
		}
		w.WriteHeader(http.StatusOK)
		jsonData, err := json.Marshal(resData)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(jsonData)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
