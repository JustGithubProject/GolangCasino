package paypal_handlers

import (
	"bytes"
	"encoding/json"
    "encoding/base64"
    "io/ioutil"
    "log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

// Getting vars from .env file
var clientID string = os.Getenv("PAYPAL_CLIENT_ID")
var paypalSecret string = os.Getenv("PAYPAL_SECRET")


func PaypalGetAccessTokenHandler(c *gin.Context) {
    paypalURL := "https://api.sandbox.paypal.com/v1/oauth2/token"
    data := url.Values{}
    auth := base64.StdEncoding.EncodeToString([]byte(clientID + ":" + paypalSecret))
    data.Set("grant_type", "client_credentials")

    req, err := http.NewRequest(http.MethodPost, paypalURL, strings.NewReader(data.Encode()))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"message": "failure to create connection"})
        return
    }
    req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
    req.Header.Set("Authorization", "Basic "+ auth)

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"message": "Something went wrong"})
        return
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        c.JSON(resp.StatusCode, gin.H{"message": "Failed to get access token", "status": resp.Status})
        return
    }

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to read response body"})
        return
    }

    var result map[string]interface{}
    err = json.Unmarshal(body, &result)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to unmarshal JSON"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"access_token": result["access_token"]})
}