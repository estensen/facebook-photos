package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/user"

	fb "github.com/huandu/facebook"
)

var USERNAME = flag.String("username", "", "Facebook username")
var TOKEN = flag.String("token", "", "Facebook token")

func runFacebookQuery(query string) (result fb.Result) {
	fmt.Println(query)
	res, err := fb.Get(query, fb.Params{
		"access_token": *TOKEN,
	})
	if err != nil {
		log.Fatalln("Could not access your feed")
	}
	return res
}

func downloadPhoto(query string, baseDir string) {
	// Save photo to disk
}

func downloadAlbumPhotos(name string, albumId string, baseDir string) {
	dir := fmt.Sprintf("%v/%v/%v", baseDir, name, albumId)
	os.MkdirAll(dir, 0755)

	albumQuery := "/" + albumId + "/photos"
	albumResult := runFacebookQuery(albumQuery)
	var photos []fb.Result
	albumResult.DecodeField("data", &photos)

	for _, photo := range photos {
		photoId := photo["id"].(string)
		downloadPhoto(photoId, dir)
	}
}

func main() {
	flag.Parse()

	if len(*TOKEN) == 0 {
		log.Fatalln("You need to provide a token (--token=<token>)")
	}

	usr, _ := user.Current()
	baseDir := fmt.Sprintf("%v/Pictures/Facebook", usr.HomeDir)
	fmt.Println("Saving files to", baseDir)

	userQuery := "/me?fields=first_name"
	userResult := runFacebookQuery(userQuery)
	var user User
	userResult.Decode(&user)
	fmt.Println("Downloading pictures for", user.FirstName)

	albumsQuery := "/me/albums"
	albumsResult := runFacebookQuery(albumsQuery)
	var albums []fb.Result
	albumsResult.DecodeField("data", &albums)

	for _, album := range albums {
		albumId := album["id"].(string)
		downloadAlbumPhotos(user.FirstName, albumId, baseDir)
	}
}
