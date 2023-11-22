package handler

import (
	"GoogleProjectBackend/sql"
	"GoogleProjectBackend/util"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
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
	util.EnableCors(&w)
	ImageID := mux.Vars(r)["id"]
	fmt.Println("serveVideosHandler : " + r.Method)
	fmt.Println(ImageID)
	if r.Method == http.MethodGet {
		imageFilePath := fmt.Sprintf("./originalImages/originalImage%s.jpg", ImageID)
		//각 URL에 알맞는 비디오 지정
		http.ServeFile(w, r, imageFilePath)
	}
}

func ServeArtifactImagesHandler(w http.ResponseWriter, r *http.Request) {
	// http.ServeFile(w, r, "./videos/video1.mp4")
	// 비디오 파일을 읽어서 클라이언트로 전송
	util.EnableCors(&w)
	ImageID := mux.Vars(r)["id"]
	fmt.Println("serveVideosHandler : " + r.Method)
	fmt.Println(ImageID)
	if r.Method == http.MethodGet {
		imageFilePath := fmt.Sprintf("./artifactImages/artifactImage%s.jpg", ImageID)
		//각 URL에 알맞는 비디오 지정
		http.ServeFile(w, r, imageFilePath)
	}
}

func UploadImageHandler(w http.ResponseWriter, r *http.Request) {
	util.EnableCors(&w)
	err := r.ParseMultipartForm(500 << 20) //50MB 프론트에서 용량 계산이 가능..?
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	tags := r.MultipartForm.Value["tags"]
	fmt.Println(tags)
	var originalImagesName []string
	var oriImages []string
	var originalVideosFileForm []string
	originalImages := r.MultipartForm.File["original"]
	for _, originalImageHeader := range originalImages {
		//들어온 비디오를 video+n 이름으로 바꿈
		originalImagesName = append(originalImagesName, filepath.Base(originalImageHeader.Filename))
		originalImageFileForm := filepath.Ext(originalImageHeader.Filename)
		originalVideosFileForm = append(originalVideosFileForm, originalImageFileForm)
		originalVideoPath := "./originalImages/"
		count, err := util.CountFile(originalVideoPath)
		if err != nil {
			fmt.Print("CountFile error : ")
			fmt.Println(err)
		}
		originalVideoName := "originalImage" + fmt.Sprint(count)
		oriImages = append(oriImages, originalVideoName)
		originalFilePath := originalVideoPath + originalVideoName + originalImageFileForm

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
	var artifactImagesName []string
	var artiImages []string
	var arfectImagesFileForm []string
	artifacVideos := r.MultipartForm.File["artifact"]
	for _, artifactVideoHeader := range artifacVideos {
		//들어온 비디오를 video+n 이름으로 바꿈
		artifactImagesName = append(artifactImagesName, filepath.Base(artifactVideoHeader.Filename))
		artifactVideoFileForm := filepath.Ext(artifactVideoHeader.Filename)
		arfectImagesFileForm = append(arfectImagesFileForm, artifactVideoFileForm)

		artifactVideoPath := "./artifactImages/"
		count, err := util.CountFile(artifactVideoPath)
		if err != nil {
			fmt.Print("CountFile error : ")
			fmt.Println(err)
		}
		artifactVideoName := "artifactImage" + fmt.Sprint(count)
		artiImages = append(artiImages, artifactVideoName)
		artifactFilePath := artifactVideoPath + artifactVideoName + artifactVideoFileForm
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
	tags = strings.Split(tags[0], ",")
	for i := 0; i < len(oriImages) && i < len(artiImages); i++ {
		for j := 0; j < len(tags); j++ {
			uuid := util.MakeUUID()
			err := sql.InsertImageId(uuid, originalImagesName[j], oriImages[i], artifactImagesName[i], artiImages[i], tags[i])
			if err != nil {
				log.Println(err)
			}
		}
	}

	fmt.Fprintln(w, "Image uploaded successfully")
}
