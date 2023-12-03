package handler

import (
	"GoogleProjectBackend/sql"
	"GoogleProjectBackend/util"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
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
	// 이미지 파일을 읽어서 클라이언트로 전송
	session, _ := util.Store.Get(r, "survaySession")
	session.Options = &sessions.Options{
		MaxAge: 1800, // 초 단위
	}
	session.Save(r, w)
	util.EnableCors(&w)
	ImageID := mux.Vars(r)["id"]
	fmt.Println("serveVideosHandler : " + r.Method)
	fmt.Println(ImageID)
	if r.Method == http.MethodGet {
		imageFilePath := fmt.Sprintf("./originalImages/originalImage%s.png", ImageID)
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
		imageFilePath := fmt.Sprintf("./artifactImages/artifactImage%s.png", ImageID)
		//각 URL에 알맞는 비디오 지정
		http.ServeFile(w, r, imageFilePath)
	}
}

func ServeDiffImagesHandler(w http.ResponseWriter, r *http.Request) {
	// http.ServeFile(w, r, "./videos/video1.mp4")
	// 비디오 파일을 읽어서 클라이언트로 전송
	util.EnableCors(&w)
	ImageID := mux.Vars(r)["id"]
	fmt.Println("serveVideosHandler : " + r.Method)
	fmt.Println(ImageID)
	if r.Method == http.MethodGet {
		imageFilePath := fmt.Sprintf("./diffImages/diffImage%s.png", ImageID)
		//각 URL에 알맞는 비디오 지정
		http.ServeFile(w, r, imageFilePath)
	}
}

func UploadImageHandler(w http.ResponseWriter, r *http.Request) {
	util.EnableCors(&w)
	err := r.ParseMultipartForm(5000 << 20) //50MB 프론트에서 용량 계산이 가능..?
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
	//artifactImage란 이름으로 들어온 multiform (비디오)파일들 읽기
	var artifactImagesName []string
	var artiImages []string
	var arfectImagesFileForm []string
	artifacImages := r.MultipartForm.File["artifact"]
	for _, artifactImageHeader := range artifacImages {
		//들어온 비디오를 video+n 이름으로 바꿈
		artifactImagesName = append(artifactImagesName, filepath.Base(artifactImageHeader.Filename))
		artifactImageFileForm := filepath.Ext(artifactImageHeader.Filename)
		arfectImagesFileForm = append(arfectImagesFileForm, artifactImageFileForm)

		artifactImagePath := "./artifactImages/"
		count, err := util.CountFile(artifactImagePath)
		if err != nil {
			fmt.Print("CountFile error : ")
			fmt.Println(err)
		}
		artifactVideoName := "artifactImage" + fmt.Sprint(count)
		artiImages = append(artiImages, artifactVideoName)
		artifactFilePath := artifactImagePath + artifactVideoName + artifactImageFileForm
		artifactOutputFile, err := os.Create(artifactFilePath)

		if err != nil {
			fmt.Println("artifact image fail")
			http.Error(w, "Unable to create the file for writing", http.StatusInternalServerError)
			return
		}
		defer artifactOutputFile.Close()

		artifactImage, err := artifactImageHeader.Open()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			fmt.Println(err.Error())
			return
		}
		_, err = io.Copy(artifactOutputFile, artifactImage)
		if err != nil {
			http.Error(w, "Unable to write the file", http.StatusInternalServerError)
		}
	}

	var diffImagesName []string
	var diffs []string
	var diffImagesFileForm []string
	diffImages := r.MultipartForm.File["diff"]
	for _, diffImageHeader := range diffImages {
		//들어온 이미지를 video+n 이름으로 바꿈
		diffImagesName = append(diffImagesName, filepath.Base(diffImageHeader.Filename))
		diffImageFileForm := filepath.Ext(diffImageHeader.Filename)
		diffImagesFileForm = append(diffImagesFileForm, diffImageFileForm)

		diffImagePath := "./diffImages/"
		count, err := util.CountFile(diffImagePath)
		if err != nil {
			fmt.Print("CountFile error : ")
			fmt.Println(err)
		}
		diffImageName := "diffImage" + fmt.Sprint(count)
		diffs = append(diffs, diffImageName)
		diffFilePath := diffImagePath + diffImageName + diffImageFileForm
		diffOutputFile, err := os.Create(diffFilePath)

		if err != nil {
			fmt.Println("diff image fail")
			http.Error(w, "Unable to create the file for writing", http.StatusInternalServerError)
			return
		}
		defer diffOutputFile.Close()

		diffImage, err := diffImageHeader.Open()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			fmt.Println(err.Error())
			return
		}
		_, err = io.Copy(diffOutputFile, diffImage)
		if err != nil {
			http.Error(w, "Unable to write the file", http.StatusInternalServerError)
		}
	}
	fmt.Println("Here is the problem-------------")
	fmt.Println(originalImagesName)
	fmt.Println(oriImages)
	fmt.Println(artifactImagesName)
	fmt.Println(artiImages)
	fmt.Println(tags)
	tags = strings.Split(tags[0], ",")
	for i := 0; i < len(oriImages) && i < len(artiImages); i++ {
		for j := 0; j < len(tags); j++ {
			uuid := util.MakeUUID()
			err := sql.InsertImageId(uuid, originalImagesName[i], oriImages[i], artifactImagesName[i], artiImages[i], diffImagesName[i], diffs[i], tags[j])
			if err != nil {
				log.Println(err)
			}
		}
	}

	fmt.Fprintln(w, "Image uploaded successfully")
}

func GetImageNameListHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		util.EnableCorsResponse(&w)
	}
	if r.Method == http.MethodPost {
		util.EnableCors(&w)
		body, _ := util.ProcessRequest(w, r)

		var data map[string]interface{}
		err := json.Unmarshal(body, &data)
		if err != nil {
			http.Error(w, "Error decoding JSON data", http.StatusBadRequest)
			return
		}
		testcode := data["testcode"].(string)
		imageOriginalCSVList, _ := sql.GetImageListFromTestCode(testcode)

		imageList := util.MakeCSVToStringList(imageOriginalCSVList)
		var indexList []int
		fmt.Println(imageList)
		for _, image := range imageList {
			id := strings.TrimLeft(image, "originalImage")
			num, _ := strconv.Atoi(id)
			indexList = append(indexList, num)
		}

		imageOriginalList, imageArtifactList := sql.GetImageNameListFromVideoList(imageList)

		listData := struct {
			ImageList         []int    `json:"image_list"`
			ImageOriginalList []string `json:"original_list"`
			ImageArtifactList []string `json:"artifact_list"`
		}{
			ImageList:         indexList,
			ImageOriginalList: imageOriginalList,
			ImageArtifactList: imageArtifactList,
		}
		jsonData, err := json.Marshal(listData)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			fmt.Println("json marshal error")
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonData)
	}
}
