package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/slack-go/slack"
)

func main() {

	token := os.Getenv("SLACK_BOT_TOKEN")
	if token == "" {
		fmt.Println("SLACK_BOT_TOKEN is required")
		return
	}
	channelID := os.Getenv("CHANNEL_ID")
	if channelID == "" {
		fmt.Println("CHANNEL_ID is required")
		return
	}
	api := slack.New(token)
	channelArr := []string{channelID}
	fileArr := []string{"./Faisal Ndottijo Yerima CV.pdf"}

	for i := 0; i < len(fileArr); i++ {
		fpath := fileArr[i]
		fi, err := os.Stat(fpath)
		if err != nil {
			fmt.Printf("stat error: %v\n", err)
			return
		}
		params := slack.UploadFileParameters{
			Channels: channelArr,
			File:     fpath,
			Filename: filepath.Base(fpath),
			FileSize: int(fi.Size()),
		}
		file, err := api.UploadFile(params)
		if err != nil {
			fmt.Printf("%v\n", err)
			return
		}
		fmt.Printf("Title: %s, ID: %s\n", file.Title, file.ID)
	}
}
