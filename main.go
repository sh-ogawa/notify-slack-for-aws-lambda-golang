package main

import (
	"github.com/bluele/slack"
	"github.com/aws/aws-lambda-go/lambda"
	"context"
	"os"
	"github.com/aws/aws-lambda-go/events"
	"fmt"
)

func HandleRequest(ctx context.Context, s3Event events.S3Event) (string, error) {

	token := os.Getenv("slackToken")
	channel := os.Getenv("slackChannel")

	for _, record := range s3Event.Records {
		s3 := record.S3
		text := fmt.Sprintf("[%s - %s] Bucket = %s, Key = %s \n", record.EventSource, record.EventTime, s3.Bucket.Name, s3.Object.Key)
		chatOpt := slack.ChatPostMessageOpt{true, "sh-ogawa", "", "", nil, "", "", "", ""}

		api := slack.New(token)
		err := api.ChatPostMessage(channel, text, &chatOpt)
		if err != nil {
			panic(err)
		}
	}

	return "send slack", nil
}


func main() {
	lambda.Start(HandleRequest)
}