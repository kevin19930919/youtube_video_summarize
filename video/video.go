package video

import (
	"context"
	"io"
	"os"

	"github.com/kkdai/youtube/v2"
)

type YoutubeDownloader struct {
	Client *youtube.Client
}

func NewYoutubeDownloader() *YoutubeDownloader {
	return &YoutubeDownloader{
		Client: &youtube.Client{},
	}
}

func (yd *YoutubeDownloader) GetVideoStream(_ context.Context, videoID string, videoFileName string) {

	video, err := yd.Client.GetVideo(videoID)
	if err != nil {
		panic(err)
	}

	formats := video.Formats.WithAudioChannels() // only get videos with audio
	stream, _, err := yd.Client.GetStream(video, &formats[0])
	if err != nil {
		panic(err)
	}

	file, err := os.Create(videoFileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = io.Copy(file, stream)
	if err != nil {
		panic(err)
	}

}
