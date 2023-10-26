package app

import (
	"GoogleProjectBackend/app/handler"

	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	//gorilla mux easy for routing
	r := mux.NewRouter()
	r.HandleFunc("/", handler.MainHandler)
	r.HandleFunc("/postdata", handler.GetScoringData)
	// 마지막으로 들어온 '/' 뒤의 값이 1-50사이라면 각 URL에 맞게 주소 부여
	r.HandleFunc("/postvideo/original/{id:[1-50]+}", handler.ServeOriginalVideosHandler)
	r.HandleFunc("/postvideo/artifact/{id:[1-50]+}", handler.ServeArtifactVideosHandler)
	r.HandleFunc("/postimage/original/{id:[1-50]+}", handler.ServeOriginalImagesHandler)
	r.HandleFunc("/postimage/artifact/{id:[1-50]+}", handler.ServeArtifactImagesHandler)
	//about Login
	r.HandleFunc("/login", handler.ReqeustLoginHandler)
	r.HandleFunc("/admin/login", handler.AdminLoginHandler)
	r.HandleFunc("/signUp", handler.SighupHandler)
	//upload Data
	r.HandleFunc("/upload/video", handler.UploadVideoHandler)
	r.HandleFunc("/upload/image", handler.UploadImageHandler)
	//about tag
	r.HandleFunc("/addTag", handler.ReceivedTagHandler)
	r.HandleFunc("/getTag", handler.GetTagHandler)
	r.HandleFunc("/deleteTag", handler.DeletetagHandler)
	r.HandleFunc("/generateTestcode", handler.GetTestCodeHandler)
	r.HandleFunc("/getVideoIndexCurrentPage", handler.GetUserCurrentPage)
	r.HandleFunc("/serveImage", handler.ServeImage)
	r.HandleFunc("/getTestcodeWithTag", handler.GetTestCodeListHandler)

	return r
}
