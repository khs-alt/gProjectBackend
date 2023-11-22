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
	r.HandleFunc("/postimagedata", handler.GetImageScoreData)

	// 마지막으로 들어온 '/' 뒤의 값이 1-50사이라면 각 URL에 맞게 주소 부여
	r.HandleFunc("/postvideo/original/{id:[1-50]+}", handler.ServeOriginalVideosHandler)
	r.HandleFunc("/postvideo/artifact/{id:[1-50]+}", handler.ServeArtifactVideosHandler)
	r.HandleFunc("/postimage/original/{id:[1-50]+}", handler.ServeOriginalImagesHandler)
	r.HandleFunc("/postimage/artifact/{id:[1-50]+}", handler.ServeArtifactImagesHandler)
	//about Login
	r.HandleFunc("/login", handler.ReqeustLoginHandler)

	r.HandleFunc("/admin/login", handler.AdminLoginHandler)
	r.HandleFunc("/signUp", handler.SignupHandler)
	//upload Data
	r.HandleFunc("/upload/video", handler.UploadVideoHandler)
	r.HandleFunc("/upload/image", handler.UploadImageHandler)
	//about tag
	r.HandleFunc("/addTag", handler.ReceivedTagHandler)
	r.HandleFunc("/addImageTag", handler.ReceivedImageTagHandler)

	r.HandleFunc("/getTag", handler.GetTagHandler)
	r.HandleFunc("/getImageTag", handler.GetImageTagHandler)

	r.HandleFunc("/deleteTag", handler.DeletetagHandler)
	r.HandleFunc("/deleteImageTag", handler.DeleteImagetagHandler)

	r.HandleFunc("/generateTestcode", handler.GetTestCodeHandler)
	r.HandleFunc("/generateImageTestcode", handler.GetImageTestCodeHandler)

	r.HandleFunc("/getVideoIndexCurrentPage", handler.GetUserCurrentPage)
	r.HandleFunc("/getImageIndexCurrentPage", handler.GetUserCurrentImagePage)

	r.HandleFunc("/serveImage", handler.ServeImage)

	r.HandleFunc("/getTestcodeWithTag", handler.GetTestCodeListHandler)
	r.HandleFunc("/getImageTestcodeWithTag", handler.GetImageTestCodeListHandler)

	r.HandleFunc("/getVideoListFromTag", handler.GetVideoListFromTagHandler)
	r.HandleFunc("/getImageListFromTag", handler.GetImageListFromTagHandler)

	r.HandleFunc("/getCSVFile", handler.MakeCSVFromTestHandler)

	r.HandleFunc("/getUserScore", handler.GetScoreDataFromUser)
	r.HandleFunc("/getUserImageScore", handler.GetImageScoreDataFromUser)
	//r.HandleFunc("/getUserImageScore", handler.GetScoreImageDataFromUser)

	//labeling

	//Patch 이미지의 사이즈(총 개수, 가로, 세로)
	//r.HandleFunc("/patchsize", handler.GetPatchSizeHandler)

	return r
}
