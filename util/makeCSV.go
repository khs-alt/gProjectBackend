package util

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

func SaveToCSV(originalVideoList []string, artifactVideoList []string, score []int, tag string) (string, error) {
	// 파일 생성
	filepath := fmt.Sprintf("./csv/%s.csv", tag)
	file, err := os.Create(filepath)
	if err != nil {
		return "", fmt.Errorf("cannot create file: %v", err)
	}
	defer file.Close()

	// CSV 작성기 생성
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// 데이터 쓰기
	for i := 0; i < len(originalVideoList); i++ {
		record := []string{originalVideoList[i], strconv.Itoa(score[i]), artifactVideoList[i]}
		if err := writer.Write(record); err != nil {
			return "", fmt.Errorf("cannot write to file: %v", err)
		}
	}

	return filepath, nil
}
