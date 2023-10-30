package util

import (
	"fmt"
	"os"
	"path/filepath"
)

func DeleteAllFilesInFolder(folderPath string) error {
	err := filepath.Walk(folderPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 디렉토리는 건너뜁니다.
		if info.IsDir() {
			return nil
		}

		// 파일 삭제
		err = os.Remove(path)
		if err != nil {
			return err
		}

		fmt.Printf("파일 삭제: %s\n", path)
		return nil
	})

	return err
}
