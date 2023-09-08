package audio_translate

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/sashabaranov/go-openai"
)

const (
	inputFileName  = "video.mp4"
	outputBaseName = "split_part"
	tempDir        = "temp"
)

type OpenAIServer struct {
	Client *openai.Client
}

func NewOpenAIServer(apiKey string) *OpenAIServer {
	return &OpenAIServer{
		Client: openai.NewClient(apiKey),
	}
}

func (s *OpenAIServer) SpeechToText(ctx context.Context, inputFileName string) {
	// Create a temporary directory for storing split files
	if err := os.MkdirAll(tempDir, 0755); err != nil {
		fmt.Println("Error creating temporary directory:", err)
		return
	}
	defer os.RemoveAll(tempDir)

	splitCmd := exec.Command("ffmpeg", "-i", inputFileName, "-f", "segment", "-segment_time", "360", "-c", "copy", filepath.Join(tempDir, outputBaseName+"%03d.mp4"))

	splitCmd.Stdout = os.Stdout
	splitCmd.Stderr = os.Stderr

	if err := splitCmd.Run(); err != nil {
		fmt.Println("Error splitting video:", err)
		return
	}

	fmt.Println("Video split into parts successfully")
	splitFileNames := getFileNamesInFolder(tempDir)
	// Process each split file
	for _, splitFilePath := range splitFileNames {

		req := openai.AudioRequest{
			Model:    openai.Whisper1,
			FilePath: splitFilePath,
		}
		resp, err := s.Client.CreateTranscription(ctx, req)
		if err != nil {
			fmt.Printf("Transcription error: %v\n", err)
			return
		}
		fmt.Println(resp.Text)
	}
}

func getFileNamesInFolder(folderPath string) []string {
	files, err := ioutil.ReadDir(folderPath)
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}

	var fileNames []string
	for _, file := range files {
		if !file.IsDir() {
			fullOutputPath := filepath.Join(tempDir, file.Name())
			fileNames = append(fileNames, fullOutputPath)
		}
	}

	return fileNames
}
