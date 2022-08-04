package contact

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

// создание клиента
// Универсальный вызов

const (
	baseURL = "https://api.facest.io/v1/contact/"
	token   = "YWJkZGJkYzYtNTAzNS00ZGU2LTlmMDgtNzQ2YTQ3NzM1YTBl"
)

type Client struct {
	baseURL    string
	token      string
	HttpClient *http.Client
}

func NewHttpClient(time time.Duration) *Client {
	return &Client{
		baseURL: baseURL,
		token:   token,
		HttpClient: &http.Client{
			Timeout: time,
		},
	}
}

func (c *Client) SendPostRequest(endpoint string, request, response interface{}) error {
	url := c.baseURL + "/" + endpoint

	reqBytes, err := json.Marshal(&request)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(reqBytes))

	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.token)
	req.Close = true

	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, &response)
	if err != nil {
		return err
	}
	return nil
}
