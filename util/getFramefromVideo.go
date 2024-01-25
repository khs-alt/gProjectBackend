package util

import (
	"os/exec"
)

func ExtractFrame(videoFilePath string, videoCurrentTime string, outputImage string) error {
	cmd := exec.Command("ffmpeg", "-i", videoFilePath, "-ss", videoCurrentTime, "-vframes", "1", outputImage)
	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}
