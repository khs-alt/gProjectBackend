package util

import (
	"fmt"
	"os/exec"
)

func ExtractFrame(videoFilePath string, videoCurrentTime float32, outputImage string) error {
	vt := fmt.Sprintln(videoCurrentTime)
	cmd := exec.Command("ffmpeg", "-i", videoFilePath, "-ss", vt, "-vframes", "1", "-vf", "scale=320:-1", outputImage)
	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}
