package app

import (
	"backend/app/handler"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	// Main route
	r.GET("/", handler.MainHandler)

	// Grouping routes under /label/api
	api := r.Group("/label/api")
	{

		// POST routes
		api.POST("/postdata", handler.GetVideoScoringData)
		api.POST("/postimagedata", handler.GetImageScoreData)

		api.POST("/login", handler.RequestLoginHandler)
		api.POST("/admin/login", handler.AdminLoginHandler)
		api.POST("/signUp", handler.SignupHandler)

		api.POST("/upload/video", handler.UploadVideoHandler)
		api.POST("/upload/image", handler.UploadImageHandler)

		api.POST("/addTag", handler.ReceivedVideoTagHandler)
		api.POST("/addImageTag", handler.ReceivedImageTagHandler)

		api.POST("/deleteTag", handler.DeleteVideotagHandler)
		api.POST("/deleteImageTag", handler.DeleteImagetagHandler)

		api.POST("/generateTestcode", handler.GetVideoTestCodeHandler)
		api.POST("/generateImageTestcode", handler.GetImageTestCodeHandler)

		api.POST("/getVideoIndexCurrentPage", handler.GetUserCurrentPageInfo)
		api.POST("/getImageIndexCurrentPage", handler.GetUserCurrentImagePage)

		api.POST("/exportVideo", handler.ExportVideoDataHandler)
		api.POST("/exportImage", handler.ExportImageDataHandler)

		// select frame
		api.POST("admin/postVideoFrameTime", handler.PostVideoFrameTimeHandler)

		// GET routes
		api.GET("/postvideo/original/:id", handler.ServeOriginalVideosHandler)
		api.GET("/postvideo/artifact/:id", handler.ServeArtifactVideosHandler)
		api.GET("/postvideo/diff/:id", handler.ServeDiffVideosHandler)
		api.GET("/postimage/original/:id", handler.ServeOriginalImagesHandler)
		api.GET("/postimage/artifact/:id", handler.ServeArtifactImagesHandler)
		api.GET("/postimage/difference/:id", handler.ServeDiffImagesHandler)

		api.GET("/getTag", handler.GetVideoTagHandler)
		api.GET("/getImageTag", handler.GetImageTagHandler)

		api.GET("/serveImage", handler.ServeImage)

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
	// You can also add more route groups if needed for different URL prefixes
}
