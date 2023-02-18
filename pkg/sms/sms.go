package sms

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/L1LSunflower/auction/pkg/logger"
	"github.com/L1LSunflower/auction/pkg/logger/message"
)

const (
	urlAddress   = "https://api.mobizon.kz/service/Message/SendSmsMessage?output=json&api=v1"
	apiKey       = "&apiKey=kz8060053a17d8c500c6daeafd2511b2991c734acb041d5f1937e6b29e4bd068920f6f"
	contentType  = "application/x-www-form-urlencoded"
	cacheControl = "cache-control: no-cache"
)

var requestHeader = map[string]string{
	"content-type":  contentType,
	"cache-control": cacheControl,
}

func SendSMS(recipient, code string) error {
	data := map[string]any{}

	bodyBytes := strings.NewReader(fmt.Sprintf("recipient=%s&text=%s", recipient, code) + "&params%5Bvalidity%5D=1440")
	respData, err := doRequest(bodyBytes, urlAddress+apiKey)
	if err != nil {
		return err
	}

	err = json.Unmarshal(respData, &data)
	if err != nil {
		return err
	}

	logger.Log.Info(message.NewMessage(fmt.Sprintf("response data: %v", data)))

	return nil
}

func doRequest(requestBody *strings.Reader, url string) ([]byte, error) {
	client := http.Client{
		Timeout: 10 * time.Second,
	}

	req, err := http.NewRequest("POST", url, requestBody)
	if err != nil {
		return nil, err
	}

	for name, value := range requestHeader {
		req.Header.Add(name, value)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	respData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return respData, nil
}
