package handler

import (
	"GoogleProjectBackend/sql"
	"GoogleProjectBackend/util"
	"net/http"

	"github.com/joho/sqltocsv"
)

func ExportImageDataHandler(w http.ResponseWriter, r *http.Request) {
	util.EnableCors(&w)
	if r.Method == http.MethodGet {

		rows, err := sql.ExportImageData()
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
	util.EnableCors(&w)
	if r.Method == http.MethodGet {

		rows, err := sql.ExportVideoData()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-type", "text/csv")
		w.Header().Set("Content-Disposition", "attachment; filename=\"userVideoData.csv\"")

		sqltocsv.Write(w, rows)
	}
}
