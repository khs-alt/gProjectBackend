package util

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"gopkg.in/vansante/go-ffprobe.v2"
)

// r_frame_rate로부터 FPS를 계산합니다.
func getFPSFromRFrameRate(rFrameRate string) float64 {
	var num, den int
	_, err := fmt.Sscanf(rFrameRate, "%d/%d", &num, &den)
	if err != nil {
		log.Fatalf("Error parsing r_frame_rate: %v", err)
	}

	return float64(num) / float64(den)
}

func GetVideoFPS(videoPath string) float32 {
	ctx, cancelFn := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFn()

	// ffprobe로 정보를 가져옵니다.
	data, err := ffprobe.ProbeURL(ctx, videoPath)
	if err != nil {
		log.Println("Error probing video: %v", err)
	}

	// 비디오 스트림에서 프레임 속도를 찾습니다.
	for _, stream := range data.Streams {
		if stream.CodecType == "video" {
			fps := getFPSFromRFrameRate(stream.RFrameRate)
			roundedFPS, err := strconv.ParseFloat(fmt.Sprintf("%.2f", fps), 32)
			if err != nil {
				log.Println(err)
			}
			return float32(roundedFPS)
		}
	}
	return 0
}
