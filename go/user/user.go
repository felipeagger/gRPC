package user

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"golang.org/x/net/context"
)

type Server struct {
}

type user struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Username  string `json:"login"`
	AvatarURL string `json:"avatar_url"`
	Location  string `json:"location"`
	Repos     int64  `json:"public_repos"`
}

// felipeagger

//GetUser is get user on github
func (s *Server) GetUser(ctx context.Context, in *UserRequest) (*UserResponse, error) {
	log.Printf("Receive message from client: %s", in.Username)

	res, err := http.Get(fmt.Sprintf("https://api.github.com/users/%v", in.Username))
	if err != nil {
		log.Fatal(err)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	usr := user{}
	jsonErr := json.Unmarshal(body, &usr)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	return &UserResponse{
		Id:        usr.ID,
		Name:      usr.Name,
		Username:  usr.Username,
		Avatarurl: usr.AvatarURL,
		Location:  usr.Location,
		Repos:     usr.Repos}, nil
}
