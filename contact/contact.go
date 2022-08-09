package contact

import (
	"fmt"
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
	ContactData *Data `json:"contact,omitempty"`
}

func (c *Data) PrintContactData() {
	fmt.Printf("Conatct ID = %s, Name = %s, Birthday = %s, Note = %s,  Phones = %+v, Addresses = %+v, Emails = %+v, SocialMedias = %+v, Links = %+v, CreatedAt = %s UpdatedAt = %s\n",
		c.Id,
		c.Name,
		c.Birthday,
		c.Note,
		c.Phones,
		c.Addresses,
		c.Emails,
		c.SocialMedias,
		c.Links,
		c.CreatedAt,
		c.UpdatedAt)

}

type ReadRequest struct {
	Id string `json:"id,omitempty"`
}

type ReadResponse struct {
	ContactData *Data `json:"contact,omitempty"`
}

type UpdateRequest struct {
	Id           string        `json:"id,omitempty"`
	Name         string        `json:"name,omitempty"`
	Birthday     string        `json:"birthday,omitempty"`
	Note         string        `json:"note,omitempty"`
	Phones       []Phone       `json:"phones,omitempty"`
	Addresses    []Address     `json:"addresses,omitempty"`
	Emails       []Email       `json:"emails,omitempty"`
	SocialMedias []SocialMedia `json:"socialMedias,omitempty"`
	Links        []Link        `json:"links,omitempty"`
}

type UpdateResponse struct {
	ContactData *Data `json:"contact,omitempty"`
}

type ListRequest struct {
	Limit  int32 `json:"limit,omitempty"`
	Offset int32 `json:"offset,omitempty"`
}

type ListResponse struct {
	ContactDataList []Data `json:"contacts,omitempty"`
}

type DeleteRequest struct {
	Id string `json:"id,omitempty"`
}

type DeleteResponse struct {
}

func NewEngine() *Engine {
	return &Engine{HttpClient: NewHttpClient(time.Minute)}
}

func (e *Engine) CreateContact(request *CreateRequest) (*CreateResponse, error) {
	response := &CreateResponse{}
	err := e.HttpClient.SendPostRequest("Create", request, response)
	return response, err
}

func (e *Engine) ReadContact(request *ReadRequest) (*ReadResponse, error) {
	response := &ReadResponse{}
	err := e.HttpClient.SendPostRequest("Read", request, response)
	return response, err
}

func (e *Engine) UpdateContact(request *UpdateRequest) (*UpdateResponse, error) {
	response := &UpdateResponse{}
	err := e.HttpClient.SendPostRequest("Update", request, response)
	return response, err
}

func (e *Engine) ListContacts(request *ListRequest) (*ListResponse, error) {
	response := &ListResponse{}
	err := e.HttpClient.SendPostRequest("List", request, response)
	return response, err
}

func (e *Engine) DeleteContact(request *DeleteRequest) (*DeleteResponse, error) {
	response := &DeleteResponse{}
	err := e.HttpClient.SendPostRequest("Delete", request, response)
	return response, err
}
