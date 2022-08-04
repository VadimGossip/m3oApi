package contact

import (
	"time"
)

type Engine struct {
	HttpClient *Client
}

type Address struct {
	Label    string `json:"label"`
	Location string `json:"location"`
}

type Email struct {
	Label   string `json:"label"`
	Address string `json:"address"`
}

type Link struct {
	Label string `json:"label"`
	Url   string `json:"url"`
}

type Phone struct {
	Label  string `json:"label"`
	Number string `json:"number"`
}

type SocialMedia struct {
	Label    string `json:"label"`
	Username string `json:"username"`
}

type CreateRequest struct {
	Name         string        `json:"name,omitempty"`
	Birthday     string        `json:"birthday,omitempty"`
	Note         string        `json:"note,omitempty"`
	Phones       []Phone       `json:"phones,omitempty"`
	Addresses    []Address     `json:"addresses,omitempty"`
	Emails       []Email       `json:"emails,omitempty"`
	SocialMedias []SocialMedia `json:"socialMedias,omitempty"`
	Links        []Link        `json:"links,omitempty"`
}

type Data struct {
	Id           string        `json:"id,omitempty"`
	Name         string        `json:"name,omitempty"`
	Birthday     string        `json:"birthday,omitempty"`
	Note         string        `json:"note,omitempty"`
	Phones       []Phone       `json:"phones,omitempty"`
	Addresses    []Address     `json:"addresses,omitempty"`
	Emails       []Email       `json:"emails,omitempty"`
	SocialMedias []SocialMedia `json:"socialMedias,omitempty"`
	Links        []Link        `json:"links,omitempty"`
	CreatedAt    string        `json:"created_at,omitempty"`
	UpdatedAt    string        `json:"updated_at,omitempty"`
}

type CreateResponse struct {
	Data `json:"contact"`
}

func NewEngine() *Engine {
	return &Engine{HttpClient: NewHttpClient(time.Minute)}
}

func (e *Engine) CreateContact(request *CreateRequest) (*CreateResponse, error) {
	response := &CreateResponse{}
	err := e.HttpClient.SendPostRequest("Create", request, response)
	return response, err
}
