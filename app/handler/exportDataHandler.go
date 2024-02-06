package handler

import (
	"backend/app/models"
	"backend/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/sqltocsv"
)

func ExportImageDataHandler(c *gin.Context) {
	var requestData models.TestCodeData

	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	testcode := requestData.TestCode
	rows, err := sql.ExportImageData(testcode)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()
	w := c.Writer
	w.Header().Set("Content-type", "text/csv")
	w.Header().Set("Content-Disposition", "attachment; filename=\"report.csv\"")

	sqltocsv.Write(w, rows)
}

func ExportVideoDataHandler(c *gin.Context) {
	w := c.Writer
	var requestData models.TestCodeData

	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	testcode := requestData.TestCode
	rows, err := sql.ExportVideoData(testcode)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	c.Header("Content-Disposition", "attachment; filename=video_scoring_data_"+testcode+".csv")
	sqltocsv.Write(w, rows)
}

func ExportVideoFrameDataHandler(c *gin.Context) {
	w := c.Writer
	rows, err := sql.ExportVideoFrameData()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	sqltocsv.Write(w, rows)
}
