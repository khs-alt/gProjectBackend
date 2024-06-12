package app

import (
	"backend/app/handler"

	"github.com/gin-gonic/gin"
)

// 이미지와 비디오에 관련된 데이터를 다른 DB에서 각각 처리합니다
// 그렇기에 대부분의 Handler는 이미지와 비디오로 나뉩니다
func Routes(r *gin.Engine) {
	// Main route
	r.GET("/", handler.MainHandler)

	// Grouping routes under /label/api
	// TODO:reverse proxy, cors
	api := r.Group("/label/api")
	{
		// 유저에게 받은 비디오 및 이미지 평가데이터를 DB에 저장
		api.POST("/postdata", handler.GetVideoScoringData)
		api.POST("/postimagedata", handler.GetImageScoreData)

		// 회원기입 및 일반 유저, 관리자 로그인 핸들러
		api.POST("/login", handler.RequestLoginHandler)
		api.POST("/admin/login", handler.AdminLoginHandler)
		api.POST("/signUp", handler.SignupHandler)

		// 이미지와 비디오를 각 태그와 함께 업로드 하는 핸들러
		api.POST("/upload/video", handler.UploadVideoHandler)
		api.POST("/upload/image", handler.UploadImageHandler)

		// 이미지나 비디오에 추가할 태그를 추가하고 삭제하는 핸들러
		api.POST("/addTag", handler.ReceivedVideoTagHandler)
		api.POST("/addImageTag", handler.ReceivedImageTagHandler)
		api.POST("/deleteTag", handler.DeleteVideotagHandler)
		api.POST("/deleteImageTag", handler.DeleteImagetagHandler)

		// 평가할 이미지 혹은 비디오 리스트가 담긴 테스트 코드 생성
		// 선택된 태그를 바탕으로 테스트 코드 생성
		api.POST("/generateTestcode", handler.GetVideoTestCodeHandler)
		api.POST("/generateImageTestcode", handler.GetImageTestCodeHandler)

		// 유저가 특정 테스트 코드에 작성된
		api.POST("/getVideoIndexCurrentPage", handler.GetUserCurrentPageInfo)
		api.POST("/getImageIndexCurrentPage", handler.GetUserCurrentImagePage)

		api.POST("/exportVideo", handler.ExportVideoDataHandler)
		api.POST("/exportImage", handler.ExportImageDataHandler)

		// 비디오에서 선택된 프레임을 받아서 이미지로 자르는 핸들러
		// 주요 병목 구간 중 한 곳
		api.POST("admin/postVideoFrameTime", handler.PostVideoFrameTimeHandler)

		// 업로드된 비디오와 이미지를 저장하는 곳
		// 비디오와 이미지 수십 개가 한 번에 업로드 되기에 느림
		api.GET("/postvideo/original/:id", handler.ServeOriginalVideosHandler)
		api.GET("/postvideo/artifact/:id", handler.ServeArtifactVideosHandler)
		api.GET("/postvideo/diff/:id", handler.ServeDiffVideosHandler)
		api.GET("/postimage/original/:id", handler.ServeOriginalImagesHandler)
		api.GET("/postimage/artifact/:id", handler.ServeArtifactImagesHandler)
		api.GET("/postimage/difference/:id", handler.ServeDiffImagesHandler)

		// 저장된 태그들의 보여주는 곳
		api.GET("/getTag", handler.GetVideoTagHandler)
		api.GET("/getImageTag", handler.GetImageTagHandler)

		// 웹 페이지 로고 이미지 보여주는 곳
		api.GET("/serveImage", handler.ServeImage)

		// 태그의 리스트를 보내면 해당하는 테스트 코드를 생성
		api.GET("/getTestcodeWithTag", handler.GetTestCodeListHandler)
		api.GET("/getImageTestcodeWithTag", handler.GetImageTestCodeListHandler)

		api.GET("/getVideoListFromTag", handler.GetVideoListFromTagHandler)
		api.GET("/getImageListFromTag", handler.GetImageListFromTagHandler)

		api.GET("/getuserScoringList", handler.GetUserScoringListHandler)
		api.POST("getUserLabelingList", handler.GetUserLabelingListHandler)

		api.POST("/getUserImageScore", handler.GetImageScoreDataFromUser)

		api.POST("/imageNameList", handler.GetImageNameListHandler)

		api.GET("/admin/getVideoIndex", handler.GetVideoListFromTestCodeHandler)

		api.GET("/admin/getSelectedFrameList", handler.GetSelectedFrameListHandler)

		api.POST("/getScoreCnt", handler.GetScoreCntHandler)
	}
}
