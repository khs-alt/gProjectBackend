package handler

import (
	"GoogleProjectBackend/sql"
	"GoogleProjectBackend/util"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
)

func ServeOriginalVideosHandler(w http.ResponseWriter, r *http.Request) {
	// http.ServeFile(w, r, "./videos/video1.mp4")
	// 비디오 파일을 읽어서 클라이언트로 전송
	videoID := mux.Vars(r)["id"]
	fmt.Println("serveVideosHandler : " + r.Method)
	fmt.Println(videoID)
	if r.Method == http.MethodGet {
		videoFilePath := fmt.Sprintf("./originalVideos/originalVideo%s.mp4", videoID)
		//각 URL에 알맞는 비디오 지정
		http.ServeFile(w, r, videoFilePath)
	}
}

func ServeArtifactVideosHandler(w http.ResponseWriter, r *http.Request) {
	// http.ServeFile(w, r, "./videos/video1.mp4")
	// 비디오 파일을 읽어서 클라이언트로 전송
	videoID := mux.Vars(r)["id"]
	fmt.Println("serveVideosHandler : " + r.Method)
	fmt.Println(videoID)
	if r.Method == http.MethodGet {
		videoFilePath := fmt.Sprintf("./artifactVideos/artifactVideo%s.mp4", videoID)
		//각 URL에 알맞는 비디오 지정
		http.ServeFile(w, r, videoFilePath)
	}
}

func UploadVideoHandler(w http.ResponseWriter, r *http.Request) {
	util.EnableCors(&w)
	err := r.ParseMultipartForm(50 << 20) //50MB 프론트에서 용량 계산이 가능..?
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	tags := r.MultipartForm.Value["tags"]
	fmt.Println(tags)
	var oriVideos []string
	originalVideos := r.MultipartForm.File["original"]
	for _, originalVideoHeader := range originalVideos {
		//들어온 비디오를 video+n 이름으로 바꿈
		s := strings.Split(originalVideoHeader.Filename, ".")
		lastindex := len(s) - 1
		fileForm := s[lastindex]
		originalVideoPath := "./originalVideos/"
		count, err := util.CountFile(originalVideoPath)
		if err != nil {
			fmt.Print("CountFile error : ")
			fmt.Println(err)
		}
		originalVideoName := "originalVideo" + fmt.Sprint(count)
		oriVideos = append(oriVideos, originalVideoName)
		originalFilePath := originalVideoPath + originalVideoName + "." + fileForm

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
	var artiVideos []string
	//artifactVideo란 이름으로 들어온 multiform (비디오)파일들 읽기
	artifacVideos := r.MultipartForm.File["artifact"]
	for _, artifactVideoHeader := range artifacVideos {
		//들어온 비디오를 video+n 이름으로 바꿈
		s := strings.Split(artifactVideoHeader.Filename, ".")
		lastindex := len(s) - 1
		fileForm := s[lastindex]
		artifactVideoPath := "./artifactVideos/"
		count, err := util.CountFile(artifactVideoPath)

		if err != nil {
			fmt.Print("CountFile error : ")
			fmt.Println(err)
		}
		artifactVideoName := "artifactVideo" + fmt.Sprint(count)
		artiVideos = append(artiVideos, artifactVideoName)
		artifactFilePath := artifactVideoPath + artifactVideoName + "." + fileForm
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
	tags = strings.Split(tags[0], ",")
	for i := 0; i < len(oriVideos) && i < len(artiVideos); i++ {
		for j := 0; j < len(tags); j++ {
			uuid := util.MakeUUID()
			err := sql.InsertVideoId(uuid, oriVideos[i], artiVideos[i], tags[j])
			if err != nil {
				fmt.Println(err)
			}
		}
	}
	fmt.Fprintln(w, "Video uploaded successfully")
}
