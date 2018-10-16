package main

import (
	"flag"
	"fmt"
	"log"
)

var USERNAME = flag.String("username", "", "Facebook username")
var TOKEN = flag.String("token", "", "Facebook token")

func main() {
	flag.Parse()

	if len(*USERNAME) == 0 {
		log.Fatalln("You need to provide a username (--username=<username>)")
	}

	if len(*TOKEN) == 0 {
		log.Fatalln("You need to provide a token (--token=<token>)")
	}

	fmt.Println("Downloading pictures for", *USERNAME)
}
