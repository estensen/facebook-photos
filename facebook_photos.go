package main

import (
	"flag"
	"fmt"
	"log"

	fb "github.com/huandu/facebook"
)

type User struct {
	FirstName string
}

var USERNAME = flag.String("username", "", "Facebook username")
var TOKEN = flag.String("token", "", "Facebook token")

func runFacebookQuery(query string) (result fb.Result) {
	res, err := fb.Get(query, fb.Params{
		"fields": "first_name",
		"access_token": *TOKEN,
	})
	if err != nil {
		log.Fatalln("Could not access your feed")
	}
	return res
}

func main() {
	flag.Parse()

	if len(*TOKEN) == 0 {
		log.Fatalln("You need to provide a token (--token=<token>)")
	}

	userQuery := "/me"
	userResult := runFacebookQuery(userQuery)
	var user User
	userResult.Decode(&user)
	fmt.Println("Downloading pictures for", user.FirstName)

	albumQuery := userQuery + "/albums"
	albumResult := runFacebookQuery(albumQuery)
	var albums []fb.Result
	albumResult.DecodeField("data", &albums)
	fmt.Println("Album ids:")
	for _, album := range albums {
		fmt.Println(album["id"])
	}
}
