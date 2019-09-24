package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/nlopes/slack"
	"golang.org/x/xerrors"
)

var (
	SlackWebhookURL string
	// エラーレベルに応じたSlackメッセージの色
	Color = map[string]string{
		"DEBUG":    "#4175e1",
		"INFO":     "#76a9fa",
		"WARNING":  "warning",
		"ERROR":    "danger",
		"CRITICAL": "#ff0000",
	}
)

// PubSubMessage はPubsubから取得したデータ.
type PubSubMessage struct {
	Data []byte `json:"data"`
}

func init() {
	SlackWebhookURL = os.Getenv("WEBHOOK_URL")
}

// CloudFunctionの呼び出し関数にSubscribeを設定する
func Subscribe(ctx context.Context, m PubSubMessage) error {
	stdMeg, err := unmarshal(m.Data)
	if err != nil {
		return xerrors.Errorf("can't unmarshal stackdriver message: %w", err)
	}

	msg := buildMessage(stdMeg)

	err = postWebhook(SlackWebhookURL, msg)
	if err != nil {
		return xerrors.Errorf("Failed to send a message to Slack: %v", err)
	}

	return nil
}

// stackdriverのログをunmarshal
func unmarshal(data []byte) (Message, error) {
	msg := Message{}
	err := json.Unmarshal(data, &msg)
	if err != nil {
		return msg, xerrors.Errorf("failed to unmarshal: %v, ", err)
	}
	return msg, err
}

//slackへpostするメッセージを組み立てる
func buildMessage(msg Message) *slack.WebhookMessage {
	return &slack.WebhookMessage{
		Attachments: []slack.Attachment{
			{
				Title: fmt.Sprintf("%sでエラーが発生しました", msg.Resource.Labels.ContainerName),
				Color: Color[msg.Severity],
				Text:  msg.JsonPayload.Msg, // 実際のエラー内容等が入る
				Ts:    json.Number(msg.Timestamp),
			},
		},
	}
}

func postWebhook(url string, msg *slack.WebhookMessage) error {
	return slack.PostWebhook(url, msg)
}
