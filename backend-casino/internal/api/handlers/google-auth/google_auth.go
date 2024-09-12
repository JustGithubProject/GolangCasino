package google_auth

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"github.com/markbates/goth/providers/google"

	"github.com/JustGithubProject/GolangCasino/backend-casino/internal/models"
	"github.com/JustGithubProject/GolangCasino/backend-casino/internal/services"
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
		RedirectURL:  "http://localhost:8081/google/auth/callback/",
		ClientID:     GOOGLE_CLIENT_ID,
		ClientSecret: GOOGLE_CLIENT_SECRET,
		Scopes:       []string{"email", "profile"},
		Endpoint:     google.Endpoint,
	}

	oauthStateString = "random"
}

func HandleGoogleLogin(c *gin.Context) {
	url := googleOauthConfig.AuthCodeURL(oauthStateString, oauth2.AccessTypeOffline)
	log.Println("URL", url)
	c.Redirect(http.StatusTemporaryRedirect, url)
}


func HandleGoogleCallback(c *gin.Context) {
	tokenID := c.PostForm("id_token")
	if tokenID == "" { // Error is here. TODO: ...
		log.Println("Code is empty")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Code not found"})
		return
	}

	ctx := context.Background()
	token, err := googleOauthConfig.Exchange(ctx, tokenID)
	if err != nil {
		log.Printf("Failed to exchange token: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get google token"})
		return
	}

	client := googleOauthConfig.Client(ctx, token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")
	if err != nil {
		log.Printf("Failed to get user info: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not get user info"})
		return
	}
	defer resp.Body.Close()

	var jsonResponse services.GoogleAuthInput
	if err = json.NewDecoder(resp.Body).Decode(&jsonResponse); err != nil {
		log.Printf("Failed to decode user info: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failure to decode JSON"})
		return
	}

	userRepository, err := services.InitializeUserRepository()
	if err != nil {
		log.Printf("Failed to initialize user repository: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to initialize user repository"})
		return
	}

	isUserExist, err := userRepository.FindByGoogleID(jsonResponse.ID)
	if err == nil && isUserExist.ID != 0 {
		// User already exists
		c.JSON(http.StatusOK, gin.H{"message": "User already exists", "user": isUserExist})
		return
	}

	// User does not exist, create new one
	user := models.User{
		Name:        jsonResponse.Name,
		Email:       jsonResponse.Email,
		Password:    services.GeneratePasswordForGoogleUser(),
		Balance:     0.0,
		GoogleID:    jsonResponse.ID,
		Picture:     jsonResponse.Picture,
		GivenName:   jsonResponse.GivenName,
		FamilyName:  jsonResponse.FamilyName,
		Locale:      "en-EN",
	}

	if err := userRepository.CreateGoogleUser(&user); err != nil {
		log.Println("Failed to create Google user:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create Google user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User created successfully", "user": user, "token": token})
}

// Second way to do auth

const (
	key = "randomString"
	maxAge = 86400 * 30
	isProd = false;
)

func GoogleGetAuthCallbackFunction(c *gin.Context) {
	// Getting param from URL
	err := godotenv.Load()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load .env file"})
		log.Fatal("Failed to load .env file")
	}

	GOOGLE_CLIENT_ID := os.Getenv("GOOGLE_CLIENT_ID")
	GOOGLE_CLIENT_SECRET := os.Getenv("GOOGLE_CLIENT_SECRET")
	
	if GOOGLE_CLIENT_ID == "" || GOOGLE_CLIENT_SECRET == "" {
		log.Fatal("Google client ID or secret is not set")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Google client ID or secret is not set"})
		return
	}

	goth.UseProviders(
		google.New(GOOGLE_CLIENT_ID, GOOGLE_CLIENT_SECRET, "http://localhost:8081/google/auth/callback/"),
	)


	user, err := gothic.CompleteUserAuth(c.Writer, c.Request)
	if err != nil {
		log.Fatal("Failed to complte user auth")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Failed to complete user authentication"})
		return
	}

	// RawData           map[string]interface{}
	// Provider          string
	// Email             string
	// Name              string
	// FirstName         string
	// LastName          string
	// NickName          string
	// Description       string
	// UserID            string
	// AvatarURL         string
	// Location          string
	// AccessToken       string
	// AccessTokenSecret string
	// RefreshToken      string
	// ExpiresAt         time.Time
	// IDToken           string

	userRepository, err := services.InitializeUserRepository()
	if err != nil {
		log.Printf("Failed to initialize user repository: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to initialize user repository"})
		return
	}

	isUserExist, err := userRepository.FindByGoogleID(user.UserID)
	if err == nil && isUserExist.ID != 0 {
		// User already exists
		c.JSON(http.StatusOK, gin.H{"message": "User already exists", "user": isUserExist})
		return
	}

	// User does not exist, create new one
	newUser := models.User{
		Name:        user.Name,
		Email:       user.Email,
		Password:    services.GeneratePasswordForGoogleUser(),
		Balance:     0.0,
		GoogleID:    user.UserID,
		Picture:     user.AvatarURL,
		GivenName:   user.FirstName,
		FamilyName:  user.LastName,
		Locale:      "en-EN",
	}

	if err := userRepository.CreateGoogleUser(&newUser); err != nil {
		log.Println("Failed to create Google user: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create Google user"})
		return
	}

	c.Redirect(302, "http://localhost:5173")
}