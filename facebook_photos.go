package main

import (
	"flag"
	"fmt"
	"log"

	fb "github.com/huandu/facebook"
)

var USERNAME = flag.String("username", "", "Facebook username")
var TOKEN = flag.String("token", "", "Facebook token")

func getFacebookUser() (queryResult interface{}) {
	res, err := fb.Get("/me/feed", fb.Params{
		"access_token": *TOKEN,
	})
	if err != nil {
		log.Fatalln("Could not access your feed")
	}
	return res
}

func main() {
	flag.Parse()

	if len(*USERNAME) == 0 {
		log.Fatalln("You need to provide a username (--username=<username>)")
	}

	if len(*TOKEN) == 0 {
		log.Fatalln("You need to provide a token (--token=<token>)")
	}

	fmt.Println("Downloading pictures for", *USERNAME)

	res := getFacebookUser()
	println(res)
}
