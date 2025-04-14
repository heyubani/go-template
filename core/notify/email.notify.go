package notify

// just to effect change

import (
	"encoding/json"
	"os"

	"github.com/heyubani/go-template/core/base"
	"github.com/heyubani/go-template/interfaces"
)

type Medium string

var (
	Email Medium = "EMAIL"
	SMS   Medium = "SMS"
	PUSH  Medium = "PUSH"

	NOTIFY_URL = os.Getenv("NOTIFICATION_URL")
)

type INotificationRequest struct {
	Medium       Medium         `json:"medium"`
	Recipient    string         `json:"recipient"`
	TemplateCode string         `json:"templateCode"`
	Subject      string         `json:"subject,omitempty"`
	Body         string         `json:"body,omitempty"`
	Service      string         `json:"service,omitempty"`
	Status       bool           `json:"status,omitempty"`
	Data         map[string]any `json:"data,omitempty"`
	Retries      int            `json:"retries,omitempty"`
	Meta         map[string]any `json:"meta,omitempty"`
	IsCloudsania bool           `json:"isCloudsania"`
}

func SendEmail(message INotificationRequest, logger interfaces.ILogger) {

	reqBody, err := json.Marshal(message)

	if err != nil {
		logger.Error("Error marshalling request body")
		return
	}

	headers := map[string]string{
		"Accept": "application/json",
	}

	_, errReq := base.HttpPostRequest(NOTIFY_URL, reqBody, headers, logger)

	if errReq != nil {
		logger.Errorf("Error sending notification", errReq.Error())
		return
	}

	logger.Infof("Log notification response")

	logger.Infof("Notification send successfull")
}
