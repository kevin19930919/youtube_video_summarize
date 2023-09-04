package main

import (
	"fmt"

	"github.com/kevin19930919/youtube_video_summarizer/audio_translate"
	"github.com/kevin19930919/youtube_video_summarizer/config"
	"github.com/kevin19930919/youtube_video_summarizer/video"
)

func main() {
	config, err := config.LoadConfig() // Load the API key from a config file
	if err != nil {
		fmt.Println("Error loading configuration:", err)
		return
	}

	videoID := "kllyM3xWmek"
	video.GetVideoStream(videoID)
	audio_translate.GetSummarize(config.OpenAISecret)
}
