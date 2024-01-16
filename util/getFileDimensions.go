package util

import (
	"fmt"
	"os/exec"
	"regexp"
)

// GetFileDimensions returns the width and height of a video or image file using ffprobe.
func GetFileDimensions(filePath string) (width, height int, err error) {
	// Construct the ffprobe command to get width and height.
	cmd := exec.Command("ffprobe",
		"-v", "error",
		"-select_streams", "v:0",
		"-show_entries", "stream=width,height",
		"-of", "csv=p=0",
		filePath)

	// Run the command and capture the output.
	output, err := cmd.Output()
	if err != nil {
		return 0, 0, err
	}

	// Parse the output to extract width and height.
	re := regexp.MustCompile(`(\d+),(\d+)`)
	matches := re.FindStringSubmatch(string(output))
	if len(matches) != 3 {
		return 0, 0, fmt.Errorf("could not parse output")
	}

	fmt.Sscanf(matches[1], "%d", &width)
	fmt.Sscanf(matches[2], "%d", &height)

	return width, height, nil
}
