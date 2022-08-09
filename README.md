# m3oApi
client api for m3o contacts app

Example
````
package main

import (
	"fmt"
	"github.com/VadimGossip/m3oApi/contact"
)


func main()  {
	engine := contact.NewEngine()

	createReq := contact.CreateRequest{
		Name:         "Vadim",
		Birthday:     "17.04.1982",
		Note:         "Hahaha",
	}
	createRes, err := engine.CreateContact(&createReq)
	if err != nil {
		fmt.Println(err)
	}
	createRes.ContactData.PrintContactData()


	readReq := contact.ReadRequest{Id: createRes.ContactData.Id}
	readRes, err := engine.ReadContact(&readReq)
	if err != nil {
		fmt.Println(err)
	}
	readRes.ContactData.PrintContactData()
	
	updateReq := contact.UpdateRequest{
		Id:           createRes.ContactData.Id,
		Name:         "Vaddington",
		Birthday:     "17.04.1992",
		Note:         "Nice",
		Phones:       []contact.Phone{
			{
			Label:  "home",
			Number: "+792198785432",
			},
		},
	}

	updateRes, err := engine.UpdateContact(&updateReq)
	if err != nil {
		fmt.Println(err)
	}
	updateRes.ContactData.PrintContactData()
   
	listReq := contact.ListRequest{
		Limit:  30,
	}

	listRes, err := engine.ListContacts(&listReq)
	if err != nil {
		fmt.Println(err)
	}
	for _, val := range listRes.ContactDataList {
		val.PrintContactData()
	}

	for _, val := range listRes.ContactDataList {
		delReq := contact.DeleteRequest{Id: val.Id}
		_, err := engine.DeleteContact(&delReq)
		if err != nil {
			fmt.Println(err)
		}
	}


	listResAfterDelete, err := engine.ListContacts(&listReq)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("List After Delete all")
	if len(listResAfterDelete.ContactDataList) == 0 {
		fmt.Println("Empty list")
	}
	for _, val := range listResAfterDelete.ContactDataList {
		val.PrintContactData()
	}
}

```
