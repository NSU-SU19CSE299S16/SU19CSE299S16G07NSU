package chat_controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/pusher/pusher-http-go"
)

var client = pusher.Client{
	AppID: "838001",
	Key: "3dc4abf3a5014bea3dd7",
	Secret: "8ca16014ad0bbeeac5ab",
	Cluster: "ap2",
	Secure: true,
}

type user struct {
	Name  string `json:"name" xml:"name" form:"name" query:"name"`
	Email string `json:"email" xml:"email" form:"email" query:"email"`
}

func RegisterNewUser(rw http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}

	var newUser user

	err = json.Unmarshal(body, &newUser)
	if err != nil {
		panic(err)
	}

	client.Trigger("update", "new-user", newUser)

	json.NewEncoder(rw).Encode(newUser)
}

func PusherAuth(res http.ResponseWriter, req *http.Request) {
	params, _ := ioutil.ReadAll(req.Body)
	response, err := client.AuthenticatePrivateChannel(params)
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(res, string(response))
}
