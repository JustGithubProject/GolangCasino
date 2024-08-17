package google_auth

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var googleOauthConfig *oauth2.Config
var oauthStateString string

func init() {
	// Loading .env file to parse variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to load .env file")
	}

	// Getting needed variables for google
	GOOGLE_CLIENT_ID := os.Getenv("GOOGLE_CLIENT_ID")
	GOOGLE_CLIENT_SECRET := os.Getenv("GOOGLE_CLIENT_SECRET")
	
	log.Println("GOOCLE_CLIENT_ID length: ", GOOGLE_CLIENT_ID)
	log.Println("Google_CLIENT_SECRET length: ", GOOGLE_CLIENT_SECRET)

	// Init googleOauthConfig and oauthStateString
	googleOauthConfig = &oauth2.Config{
		RedirectURL:  "http://localhost:8080/auth/callback",
		ClientID:     GOOGLE_CLIENT_ID,
		ClientSecret: GOOGLE_CLIENT_SECRET,
		Scopes:       []string{"email", "profile"},
		Endpoint:     google.Endpoint,
	}

	oauthStateString = "random" 
}


func HandleGoogleLogin(c* gin.Context){
	url := googleOauthConfig.AuthCodeURL(oauthStateString, oauth2.AccessTypeOffline)
	c.Redirect(http.StatusTemporaryRedirect, url)
}

type GoogleAuthInput struct {
	ID string `json:"id"`
	//...
}


func HandleGoogleCallback(c* gin.Context) {
	code := c.Query("code")
	token, err := googleOauthConfig.Exchange(c, code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Failure to get google token")
		return 
	}

	client := googleOauthConfig.Client(c, token)
    userInfo, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
    if err != nil {
        c.JSON(http.StatusInternalServerError, "could not get user info")
        return
    }
    defer userInfo.Body.Close()

    var user map[string]interface{}
    if err := json.NewDecoder(userInfo.Body).Decode(&user); err != nil {
        c.JSON(http.StatusInternalServerError, "could not decode user info")
        return
    }

	c.JSON(http.StatusOK, token)
}