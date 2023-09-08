package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/kevin19930919/youtube_video_summarizer/audio_translate"
	"github.com/kevin19930919/youtube_video_summarizer/config"
	"github.com/kevin19930919/youtube_video_summarizer/video"
)

const (
	inputFileName = "video.mp4"
)

func main() {
	vidPtr := flag.String("vid", "", "Video ID")
	flag.Parse()

	if *vidPtr == "" {
		fmt.Println("Usage: yvt -vid <video_id>")
		return
	}
	config, err := config.LoadConfig() // Load the API key from a config file
	if err != nil {
		fmt.Println("Error loading configuration:", err)
		return
	}

	videoID := *vidPtr

	// init youtube service
	youtubeDownloader := video.NewYoutubeDownloader()
	// init open ai service
	openAIServer := audio_translate.NewOpenAIServer(config.OpenAISecret)

	context := context.Background()

	youtubeDownloader.GetVideoStream(context, videoID, inputFileName)
	openAIServer.SpeechToText(context, inputFileName)
}
