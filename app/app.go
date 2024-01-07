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
		api.POST("/postdata", handler.GetScoringData)
		api.POST("/postimagedata", handler.GetImageScoreData)
		api.POST("/login", handler.ReqeustLoginHandler)
		api.POST("/admin/login", handler.AdminLoginHandler)
		api.POST("/signUp", handler.SignupHandler)
		api.POST("/upload/video", handler.UploadVideoHandler)
		api.POST("/upload/image", handler.UploadImageHandler)
		api.POST("/addTag", handler.ReceivedTagHandler)
		api.POST("/addImageTag", handler.ReceivedImageTagHandler)
		api.POST("/deleteTag", handler.DeletetagHandler)
		api.POST("/deleteImageTag", handler.DeleteImagetagHandler)
		api.POST("/generateTestcode", handler.GetTestCodeHandler)
		api.POST("/generateImageTestcode", handler.GetImageTestCodeHandler)
		api.POST("/getVideoIndexCurrentPage", handler.GetUserCurrentPageInfo)
		api.POST("/getImageIndexCurrentPage", handler.GetUserCurrentImagePage)
		api.POST("/getCSVFile", handler.ExportVideoDataHandler)
		api.POST("/getImageCSVFile", handler.ExportImageDataHandler)
		// api.POST("/exportImage", handler.ExportImageDataHandler)
		// api.POST("/exportVideo", handler.ExportVideoDataHandler)

		// GET routes
		api.GET("/postvideo/original/:id", handler.ServeOriginalVideosHandler)
		api.GET("/postvideo/artifact/:id", handler.ServeArtifactVideosHandler)
		api.GET("/postimage/original/:id", handler.ServeOriginalImagesHandler)
		api.GET("/postimage/artifact/:id", handler.ServeArtifactImagesHandler)
		api.GET("/postimage/difference/:id", handler.ServeDiffImagesHandler)
		api.GET("/getTag", handler.GetTagHandler)
		api.GET("/getImageTag", handler.GetImageTagHandler)
		api.GET("/serveImage", handler.ServeImage)
		api.GET("/getTestcodeWithTag", handler.GetTestCodeListHandler)
		api.GET("/getImageTestcodeWithTag", handler.GetImageTestCodeListHandler)
		api.GET("/getVideoListFromTag", handler.GetVideoListFromTagHandler)
		api.GET("/getImageListFromTag", handler.GetImageListFromTagHandler)
		api.GET("/getUserScore", handler.GetScoreDataFromUser)
		api.GET("/getUserImageScore", handler.GetImageScoreDataFromUser)
		api.GET("/imageNameList", handler.GetImageNameListHandler)
	}

	// You can also add more route groups if needed for different URL prefixes
}
