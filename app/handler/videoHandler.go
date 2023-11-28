package handler

import (
	"GoogleProjectBackend/app/models"
	"GoogleProjectBackend/sql"
	"GoogleProjectBackend/util"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gorilla/mux"
)

func ServeOriginalVideosHandler(w http.ResponseWriter, r *http.Request) {
	// http.ServeFile(w, r, "./videos/video1.mp4")
	// 비디오 파일을 읽어서 클라이언트로 전송
	videoID := mux.Vars(r)["id"]

	if r.Method == http.MethodGet {
		videoFilePrefix := fmt.Sprintf("./originalVideos/originalVideo%s", videoID)
		var videoFilePath string
		var fileExtension string
		for ext := range models.MimeTypes {
			tempPath := videoFilePrefix + ext
			if _, err := os.Stat(tempPath); err == nil {
				videoFilePath = tempPath
				fileExtension = ext
				break
			}
		}
		if videoFilePath == "" {
			http.Error(w, "Video not found", http.StatusNotFound)
			return
		}
		mimeType, exists := models.MimeTypes[fileExtension]
		if !exists {
			http.Error(w, "Unsupported video format", http.StatusUnsupportedMediaType)
			return
		}

		w.Header().Set("Content-Type", mimeType)
		http.ServeFile(w, r, videoFilePath)
	}
}

func ServeArtifactVideosHandler(w http.ResponseWriter, r *http.Request) {
	// http.ServeFile(w, r, "./videos/video1.mp4")
	// 비디오 파일을 읽어서 클라이언트로 전송
	videoID := mux.Vars(r)["id"]

	if r.Method == http.MethodGet {
		videoFilePrefix := fmt.Sprintf("./artifactVideos/artifactVideo%s", videoID)
		var videoFilePath string
		var fileExtension string
		for ext := range models.MimeTypes {
			tempPath := videoFilePrefix + ext
			if _, err := os.Stat(tempPath); err == nil {
				videoFilePath = tempPath
				fileExtension = ext
				break
			}
		}
		if videoFilePath == "" {
			http.Error(w, "Video not found", http.StatusNotFound)
			return
		}
		mimeType, exists := models.MimeTypes[fileExtension]
		if !exists {
			http.Error(w, "Unsupported video format", http.StatusUnsupportedMediaType)
			return
		}
		w.Header().Set("Content-Type", mimeType)
		http.ServeFile(w, r, videoFilePath)
	}
}

func UploadVideoHandler(w http.ResponseWriter, r *http.Request) {
	util.EnableCors(&w)
	err := r.ParseMultipartForm(2000 << 20) //500MB 프론트에서 용량 계산이 가능..?
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	tags := r.MultipartForm.Value["tags"]
	fmt.Println(tags)
	var originalVideosName []string
	var oriVideos []string
	var originalVideosFileForm []string
	originalVideos := r.MultipartForm.File["original"]
	//TODO: 여기 originalVidoeileForm 이거 이름 같은거 두 개 쓰고 있음 수정이 요함
	for _, originalVideoHeader := range originalVideos {
		//들어온 비디오를 video+n 이름으로 바꿈
		originalVideosName = append(originalVideosName, filepath.Base(originalVideoHeader.Filename))
		originalVidoeFileForm := filepath.Ext(originalVideoHeader.Filename)
		originalVideosFileForm = append(originalVideosFileForm, originalVidoeFileForm)
		originalVideoPath := "./originalVideos/"
		count, err := util.CountFile(originalVideoPath)
		if err != nil {
			fmt.Print("CountFile error : ")
			fmt.Println(err)
		}
		originalVideoName := "originalVideo" + fmt.Sprint(count)
		oriVideos = append(oriVideos, originalVideoName)
		originalFilePath := originalVideoPath + originalVideoName + originalVidoeFileForm

		originalOutputFile, err := os.Create(originalFilePath)
		if err != nil {
			fmt.Println("orignial video fail")
			http.Error(w, "Unable to create the file for writing", http.StatusInternalServerError)
			return
		}
		defer originalOutputFile.Close()
		originalImage, err := originalVideoHeader.Open()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			fmt.Println(err.Error())
			return
		}
		_, err = io.Copy(originalOutputFile, originalImage) //originFile: multipart.File
		if err != nil {
			http.Error(w, "Unable to write the file", http.StatusInternalServerError)
			return
		}
	}
	var artifactVideosName []string
	var artiVideos []string
	var arfectVideosFileForm []string
	//artifactVideo란 이름으로 들어온 multiform (비디오)파일들 읽기
	artifacVideos := r.MultipartForm.File["artifact"]
	for _, artifactVideoHeader := range artifacVideos {
		//들어온 비디오를 video+n 이름으로 바꿈
		artifactVideosName = append(artifactVideosName, filepath.Base(artifactVideoHeader.Filename))
		arfectVideoFileForm := filepath.Ext(artifactVideoHeader.Filename)
		arfectVideosFileForm = append(arfectVideosFileForm, arfectVideoFileForm)

		artifactVideoPath := "./artifactVideos/"
		count, err := util.CountFile(artifactVideoPath)
		if err != nil {
			fmt.Print("CountFile error : ")
			fmt.Println(err)
		}
		artifactVideoName := "artifactVideo" + fmt.Sprint(count)
		artiVideos = append(artiVideos, artifactVideoName)
		artifactFilePath := artifactVideoPath + artifactVideoName + arfectVideoFileForm
		artifactOutputFile, err := os.Create(artifactFilePath)

		if err != nil {
			fmt.Println("artifact video fail")
			http.Error(w, "Unable to create the file for writing", http.StatusInternalServerError)
			return
		}
		defer artifactOutputFile.Close()

		artifactVideo, err := artifactVideoHeader.Open()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			fmt.Println(err.Error())
			return
		}
		_, err = io.Copy(artifactOutputFile, artifactVideo)
		if err != nil {
			http.Error(w, "Unable to write the file", http.StatusInternalServerError)
		}
	}
	//TODO: 파일 포멧 저장 확인
	tags = strings.Split(tags[0], ",")
	for i := 0; i < len(oriVideos) && i < len(artiVideos); i++ {
		for j := 0; j < len(tags); j++ {
			uuid := util.MakeUUID()
			originaVideoFPS := util.GetVideoFPS("./originalVideos/" + oriVideos[i] + originalVideosFileForm[i])
			artifactVideoFPS := util.GetVideoFPS("./artifactVideos/" + artiVideos[i] + arfectVideosFileForm[i])
			err := sql.InsertVideoId(uuid, originalVideosName[i], oriVideos[i], originaVideoFPS, artifactVideosName[i], artiVideos[i], artifactVideoFPS, tags[j])
			if err != nil {
				fmt.Println(err)
			}
		}
	}
	fmt.Fprintln(w, "Video uploaded successfully")
}
