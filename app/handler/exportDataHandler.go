package handler

import (
	"GoogleProjectBackend/sql"
	"GoogleProjectBackend/util"
	"encoding/json"
	"net/http"

	"github.com/joho/sqltocsv"
)

func ExportImageDataHandler(w http.ResponseWriter, r *http.Request) {
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
		rows, err := sql.ExportImageData(testcode)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-type", "text/csv")
		w.Header().Set("Content-Disposition", "attachment; filename=\"userImageData.csv\"")

		sqltocsv.Write(w, rows)
	}
}

func ExportVideoDataHandler(w http.ResponseWriter, r *http.Request) {
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
		rows, err := sql.ExportVideoData(testcode)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// w.Header().Set("Content-type", "text/csv")
		// w.Header().Set("Content-Disposition", "attachment; filename=\"userVideoData.csv\"")

		sqltocsv.Write(w, rows)
	}
}
