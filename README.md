# YouTube Video Summarizer

[![License Badge](https://img.shields.io/github/license/kevin19930919/youtube_video_summarize)](https://github.com/kevin19930919/youtube_video_summarize/blob/main/LICENSE)
[![Issues Badge](https://img.shields.io/github/issues/kevin19930919/youtube_video_summarize)](https://github.com/kevin19930919/youtube_video_summarize/issues)
[![Pull Requests Badge](https://img.shields.io/github/issues-pr/kevin19930919/youtube_video_summarize)](https://github.com/kevin19930919/youtube_video_summarize/pulls)
[![Contributors Badge](https://img.shields.io/github/contributors/kevin19930919/youtube_video_summarize)](https://github.com/kevin19930919/youtube_video_summarize/graphs/contributors)
[![contributions welcome](https://img.shields.io/badge/contributions-welcome-brightgreen.svg?style=flat)](https://github.com/dwyl/esta/issues)

The YouTube Video Summarizer is a codebase that allows users to summarize the audio content of YouTube videos. It utilizes the OpenAI API to transcribe the audio and provide a summary of the video's content. The codebase consists of several modules and functions that handle video downloading, audio transcription, and splitting the video into smaller parts for processing.

## Functionalities

- Download YouTube videos and extract the audio stream
- Split the video into smaller parts for efficient processing
- Transcribe the audio using the OpenAI API
- Generate a summary of the video's content

## Installation

To install the YouTube Video Summarizer, follow these steps:

1. Clone the repository: `git clone https://github.com/kevin19930919/youtube_video_summarize.git`
2. Navigate to the project directory: `cd youtube_video_summarize`
3. Install the required dependencies: `go get -d ./...`

## Dependencies

The YouTube Video Summarizer has the following dependencies:

- [go-openai](https://github.com/sashabaranov/go-openai): Go client library for the OpenAI API
- [kkdai/youtube](https://github.com/kkdai/youtube): Go client library for the YouTube API

## Usage

To use the YouTube Video Summarizer, follow these steps:

1. Obtain an API key from OpenAI and set it as an environment variable: `export OPENAI_API_KEY=<your_api_key>`
2. Run the main program with the video ID as a command-line argument: `go run main.go -vid <video_id>`

The program will download the YouTube video, split it into smaller parts, transcribe the audio using the OpenAI API, and generate a summary of the video's content.

## Authors

The YouTube Video Summarizer codebase is maintained by Kevin (kevin19930919).

## Contributing

Contributions to the YouTube Video Summarizer are welcome! If you encounter any issues or have suggestions for improvements, please open an issue on the GitHub repository.

To contribute code changes, follow these steps:

1. Fork the repository
2. Create a new branch for your feature or bug fix: `git checkout -b my-feature`
3. Make your changes and commit them: `git commit -am 'Add new feature'`
4. Push the branch to your forked repository: `git push origin my-feature`
5. Open a pull request on the main repository

## Support

If you need support or have any questions, feel free to reach out to the maintainer at kevin@example.com.

## Commercial Support

For commercial support inquiries, please contact the maintainer at kevin@example.com.

## License

The YouTube Video Summarizer is licensed under the [MIT License](https://github.com/kevin19930919/youtube_video_summarize/blob/main/LICENSE).