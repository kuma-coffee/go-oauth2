package controllers

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	ssoGolang    *oauth2.Config
	randomString = "random-string"
)

func init() {
	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatal("error loading .env file")
	}

	ssoGolang = &oauth2.Config{
		RedirectURL:  os.Getenv("REDIRECT_URL"),
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}
}

func Signin(w http.ResponseWriter, r *http.Request) {
	url := ssoGolang.AuthCodeURL(randomString)
	fmt.Println(url)

	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}
