package handler

import (
	"GoogleProjectBackend/util"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
)

func ServeImage(w http.ResponseWriter, r *http.Request) {
	util.EnableCors(&w)
	if r.Method == http.MethodGet {
		videoFilePath := fmt.Sprintf("./image/google_logo.png")
		//각 URL에 알맞는 비디오 지정
		http.ServeFile(w, r, videoFilePath)
	}
}

func ServeOriginalImagesHandler(w http.ResponseWriter, r *http.Request) {
	// http.ServeFile(w, r, "./videos/video1.mp4")
	// 비디오 파일을 읽어서 클라이언트로 전송
	videoID := mux.Vars(r)["id"]
	fmt.Println("serveVideosHandler : " + r.Method)
	fmt.Println(videoID)
	if r.Method == http.MethodGet {
		videoFilePath := fmt.Sprintf("./artifactImages/artifactVideo%s.mp4", videoID)
		//각 URL에 알맞는 비디오 지정
		http.ServeFile(w, r, videoFilePath)
	}
}
func ServeArtifactImagesHandler(w http.ResponseWriter, r *http.Request) {
	// http.ServeFile(w, r, "./videos/video1.mp4")
	// 비디오 파일을 읽어서 클라이언트로 전송
	videoID := mux.Vars(r)["id"]
	fmt.Println("serveVideosHandler : " + r.Method)
	fmt.Println(videoID)
	if r.Method == http.MethodGet {
		videoFilePath := fmt.Sprintf("./artifactImages/artifactVideo%s.mp4", videoID)
		//각 URL에 알맞는 비디오 지정
		http.ServeFile(w, r, videoFilePath)
	}
}

func UploadImageHandler(w http.ResponseWriter, r *http.Request) {
	util.EnableCors(&w)
	err := r.ParseMultipartForm(50 << 20) //50MB 프론트에서 용량 계산이 가능..?
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	tags := r.MultipartForm.Value["tags"]
	fmt.Println(tags)
	var oriImages []string
	originalImages := r.MultipartForm.File["original"]
	for _, originalImageHeader := range originalImages {
		//들어온 비디오를 video+n 이름으로 바꿈
		s := strings.Split(originalImageHeader.Filename, ".")
		lastindex := len(s) - 1
		fileForm := s[lastindex]
		originalVideoPath := "./originalImages/"
		count, err := util.CountFile(originalVideoPath)
		if err != nil {
			fmt.Print("CountFile error : ")
			fmt.Println(err)
		}
		originalVideoName := "originalImage" + fmt.Sprint(count)
		oriImages = append(oriImages, originalVideoName)
		originalFilePath := originalVideoPath + originalVideoName + "." + fileForm

		originalOutputFile, err := os.Create(originalFilePath)
		if err != nil {
			fmt.Println("orignial image fail")
			http.Error(w, "Unable to create the file for writing", http.StatusInternalServerError)
			return
		}
		defer originalOutputFile.Close()

		originalImage, err := originalImageHeader.Open()
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
	//artifactVideo란 이름으로 들어온 multiform (비디오)파일들 읽기
	var artiImages []string
	artifacVideos := r.MultipartForm.File["artifact"]
	for _, artifactVideoHeader := range artifacVideos {
		//들어온 비디오를 video+n 이름으로 바꿈
		s := strings.Split(artifactVideoHeader.Filename, ".")
		lastindex := len(s) - 1
		fileForm := s[lastindex]
		artifactVideoPath := "./artifactImages/"
		count, err := util.CountFile(artifactVideoPath)
		if err != nil {
			fmt.Print("CountFile error : ")
			fmt.Println(err)
		}
		artifactVideoName := "artifactImage" + fmt.Sprint(count)
		artiImages = append(artiImages, artifactVideoName)
		artifactFilePath := artifactVideoPath + artifactVideoName + "." + fileForm

		artifactOutputFile, err := os.Create(artifactFilePath)
		if err != nil {
			fmt.Println("artifact image fail")
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

	fmt.Fprintln(w, "Video uploaded successfully")
}
