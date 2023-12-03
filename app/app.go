package app

import (
	"GoogleProjectBackend/app/handler"

	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	//gorilla mux easy for routing
	r := mux.NewRouter()
	authRoutes := r.PathPrefix("/label/api").Subrouter()
	authRoutes.Use(handler.SessionAuthMiddleware)
	r.HandleFunc("/", handler.MainHandler)
	r.HandleFunc("/label/api/postdata", handler.GetScoringData)
	r.HandleFunc("/label/api/postimagedata", handler.GetImageScoreData)

	// 마지막으로 들어온 '/' 뒤의 값이 1-50사이라면 각 URL에 맞게 주소 부여
	authRoutes.HandleFunc("/postvideo/original/{id:[0-9]+}", handler.ServeOriginalVideosHandler)
	r.HandleFunc("/label/api/postvideo/artifact/{id:[0-9]+}", handler.ServeArtifactVideosHandler)
	authRoutes.HandleFunc("/postimage/original/{id:[0-9]+}", handler.ServeOriginalImagesHandler)
	r.HandleFunc("/label/api/postimage/artifact/{id:[0-9]+}", handler.ServeArtifactImagesHandler)
	r.HandleFunc("/label/api/postimage/difference/{id:[0-9]+}", handler.ServeDiffImagesHandler)

	//about Login
	r.HandleFunc("/label/api/login", handler.ReqeustLoginHandler)

	r.HandleFunc("/label/api/admin/login", handler.AdminLoginHandler)
	r.HandleFunc("/label/api/signUp", handler.SignupHandler)
	//upload Data
	r.HandleFunc("/label/api/upload/video", handler.UploadVideoHandler)
	r.HandleFunc("/label/api/upload/image", handler.UploadImageHandler)
	//about tag
	r.HandleFunc("/label/api/addTag", handler.ReceivedTagHandler)
	r.HandleFunc("/label/api/addImageTag", handler.ReceivedImageTagHandler)

	r.HandleFunc("/label/api/getTag", handler.GetTagHandler)
	r.HandleFunc("/label/api/getImageTag", handler.GetImageTagHandler)

	//r.HandleFunc("/label/api/lastPage", handler.GetUserCurrentPage)

	r.HandleFunc("/label/api/deleteTag", handler.DeletetagHandler)
	r.HandleFunc("/label/api/deleteImageTag", handler.DeleteImagetagHandler)

	r.HandleFunc("/label/api/generateTestcode", handler.GetTestCodeHandler)
	r.HandleFunc("/label/api/generateImageTestcode", handler.GetImageTestCodeHandler)

	r.HandleFunc("/label/api/getVideoIndexCurrentPage", handler.GetUserCurrentPageInfo)
	r.HandleFunc("/label/api/getImageIndexCurrentPage", handler.GetUserCurrentImagePage)

	r.HandleFunc("/label/api/serveImage", handler.ServeImage)

	r.HandleFunc("/label/api/getTestcodeWithTag", handler.GetTestCodeListHandler)
	r.HandleFunc("/label/api/getImageTestcodeWithTag", handler.GetImageTestCodeListHandler)

	r.HandleFunc("/label/api/getVideoListFromTag", handler.GetVideoListFromTagHandler)
	r.HandleFunc("/label/api/getImageListFromTag", handler.GetImageListFromTagHandler)

	r.HandleFunc("/label/api/getCSVFile", handler.MakeCSVFromTestHandler)

	r.HandleFunc("/label/api/getUserScore", handler.GetScoreDataFromUser)
	r.HandleFunc("/label/api/getUserImageScore", handler.GetImageScoreDataFromUser)

	r.HandleFunc("/label/api/imageNameList", handler.GetImageNameListHandler)

	// r.HandleFunc("", handler.exportImageDataHandler)
	// r.HandleFunc("", handler.exportVideoDataHandler)

	//labeling

	//Patch 이미지의 사이즈(총 개수, 가로, 세로)
	//r.HandleFunc("/patchsize", handler.GetPatchSizeHandler)

	return r
}
