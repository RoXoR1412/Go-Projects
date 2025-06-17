package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/slack-go/slack"
)

func main(){
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	channelArr := os.Getenv("CHANNEL_ID")
	fileArr := []string{"Backend Assignment.pdf"}

	for i :=0;i<len(fileArr);i++{
		fileInfo, err := os.Stat(fileArr[i])
    	if err != nil {
        	log.Fatal(err)
    	}
		params := slack.UploadFileV2Parameters{
			Channel : channelArr,
			File: fileArr[i],
			Filename : string(fileInfo.Name()),
			FileSize: int(fileInfo.Size()),
		}
		file, err := api.UploadFileV2(params)
		if err != nil{
			fmt.Printf("%s\n",err)
			return
		}
		fmt.Printf("Name : %s, URL : %s\n", file.Title, file.ID)
	}
}