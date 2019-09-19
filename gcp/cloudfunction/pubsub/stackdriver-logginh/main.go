package main


import (
	"context"
	"encoding/json"
	"google.golang.org/genproto/googleapis/logging/v2"
	"os"
	"golang.org/x/xerrors"
	"github.com/nlopes/slack"
	"github.com/pkg/errors"
)

var (
	SlackWebhookURL string
)

// PubSubMessage is the payload of a Pub/Sub event.
type PubSubMessage struct {
	Data []byte `json:"data"`
}


type JsonPayload struct {
	field map[string]interface{}
}

func init() {
	SlackWebhookURL = os.Getenv("WEBHOOK_URL")
}

// この変数が呼ばれる
func Subscribe(ctx context.Context, m PubSubMessage) error {
	entry, err := convert2LogEntry(m.Data)
	if err != nil {
		return err
	}
	msg := buildMessage(entry)
	err = postWebhook(SlackWebhookURL, msg)
	if err != nil {
		return errors.Errorf("Failed to send a message to Slack: %s", err)
	}

	return nil
}

func convert2LogEntry(data []byte) (*logging.LogEntry, error) {
	entry := new(logging.LogEntry)
	err := json.Unmarshal(data, entry)
	if err != nil {
		return nil, xerrors.Errorf("failed to unmarshal: %w", err)
	}
	err = json.Unmarshal(entry.JsonPayload
	return entry, err
}

func buildMessage(entry *logging.LogEntry) *slack.WebhookMessage {
	return &slack.WebhookMessage{
		Text:        entry.InsertId,
		Attachments: []slack.Attachment{},
	}
}

func postWebhook(url string, msg *slack.WebhookMessage) error {
	return slack.PostWebhook(url, msg)
}
