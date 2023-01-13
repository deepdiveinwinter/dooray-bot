package clients

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/sirupsen/logrus"
)

var (
	HttpHeader = "application/json"
)

type DoorayHookClient struct {
	url             string
	botName         string
	botIconImageUrl string
}

type DoorayMessage struct {
	BotName         string             `json:"botName"`
	BotIconImageUrl string             `json:"botIconImage"`
	Text            string             `json:"text"`
	Attachments     []DoorayAttachment `json:"attachments"`
}

type DoorayAttachment struct {
	Title     string `json:"title"`
	TitleLink string `json:"titleLink"`
	Text      string `json:"text"`
	Color     string `json:"color"`
}

func NewHookClient(url string, botName string, botIconImageUrl string) *DoorayHookClient {
	return &DoorayHookClient{
		url:             url,
		botName:         botName,
		botIconImageUrl: botIconImageUrl,
	}
}

func (d *DoorayHookClient) SendMessage(message DoorayMessage) error {
	// Set Dooraybot Information
	message.BotName = d.botName
	message.BotIconImageUrl = d.botIconImageUrl

	// Encoding Dooray Message
	reqBodyBuff := new(bytes.Buffer)
	err := json.NewEncoder(reqBodyBuff).Encode(message)
	if err != nil {
		logrus.Errorf("fail to encode dooray message, error=%s", err.Error())
		return err
	}

	// Request POST
	resp, err := http.Post(d.url, HttpHeader, reqBodyBuff)
	if err != nil {
		logrus.Errorf("fail to request POST to dooray service, error=%s", err.Error())
		return err
	}
	respBody, err := io.ReadAll(resp.Body)
	logrus.Infof("response message from dooray service, message=%s", string(respBody))
	return nil
}
