package google_auth

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"


	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"

	"github.com/JustGithubProject/GolangCasino/backend-casino/internal/services"
	"github.com/JustGithubProject/GolangCasino/backend-casino/internal/models"
)

var googleOauthConfig *oauth2.Config
var oauthStateString string

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to load .env file")
	}

	GOOGLE_CLIENT_ID := os.Getenv("GOOGLE_CLIENT_ID")
	GOOGLE_CLIENT_SECRET := os.Getenv("GOOGLE_CLIENT_SECRET")
	
	if GOOGLE_CLIENT_ID == "" || GOOGLE_CLIENT_SECRET == "" {
		log.Fatal("Google client ID or secret is not set")
	}

	googleOauthConfig = &oauth2.Config{
		RedirectURL:  "http://localhost:8080/auth/callback",
		ClientID:     GOOGLE_CLIENT_ID,
		ClientSecret: GOOGLE_CLIENT_SECRET,
		Scopes:       []string{"email", "profile"},
		Endpoint:     google.Endpoint,
	}

	oauthStateString = "random"
}

func HandleGoogleLogin(c *gin.Context) {
	url := googleOauthConfig.AuthCodeURL(oauthStateString, oauth2.AccessTypeOffline)
	c.Redirect(http.StatusTemporaryRedirect, url)
}

type GoogleAuthInput struct {
	ID            string `json:"id"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Picture       string `json:"picture"`
	Locale        string `json:"locale"`
}

func HandleGoogleCallback(c *gin.Context) {
	code := c.Query("code")
	if code == "" {
		c.JSON(http.StatusBadRequest, "Code not found")
		return
	}

	ctx := context.Background()
	token, err := googleOauthConfig.Exchange(ctx, code)
	if err != nil {
		log.Printf("Failed to exchange token: %v\n", err)
		c.JSON(http.StatusInternalServerError, "Failure to get google token")
		return
	}

	client := googleOauthConfig.Client(ctx, token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		log.Printf("Failed to get user info: %v\n", err)
		c.JSON(http.StatusInternalServerError, "Could not get user info")
		return
	}
	defer resp.Body.Close()

	var jsonResponse GoogleAuthInput
	if err = json.NewDecoder(resp.Body).Decode(&jsonResponse); err != nil {
		log.Printf("Failed to decode user info: %v\n", err)
		c.JSON(http.StatusInternalServerError, "Failure to decode json")
		return
	}

	user_repository, err := services.InitializeUserRepository()
	user := models.User{
		Name: jsonResponse.Name,
		Email: jsonResponse.Email,
		Password: 
	}



	c.JSON(http.StatusOK, jsonResponse)
	// c.Redirect(http.StatusTemporaryRedirect, "/dashboard")
}

