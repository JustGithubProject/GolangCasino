package paypal_handlers

import (
	"bytes"
    "log"
	"encoding/json"
    "io/ioutil"
	"net/http"
	"strings"
    "strconv"

	"github.com/gin-gonic/gin"

    "github.com/JustGithubProject/GolangCasino/backend-casino/internal/services"
)



func CreatePaypalPaymentHandler(c *gin.Context) {
    // Getting accessToken to do API request to PAYPAL
	accessToken, err := services.PaypalGetAccessToken()
    log.Println("AcessToken: ", accessToken)
	if err != nil {
        log.Println("AccessToken is empty or invalid")
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to get access token"})
		return
	}
    log.Println("AccessToken is okay")
	paymentURL := "https://api.sandbox.paypal.com/v1/payments/payment"

    // Trying to get data from json
    var paypalInput services.PaypalPaymentInput
    if err := c.BindJSON(&paypalInput); err != nil {
        log.Println("Failed to bind JSON:", err)
        c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid paypal input"})
        return
    }

    log.Println("Received data: ", paypalInput)

    // Getting amount of money and currency to execute payment using paypal method
    total, currency := services.ExtractPaypalPaymentData(&paypalInput)

    log.Println("Total: ", total)
    log.Println("Currency: ", currency)

    if total == "" || currency == ""{
        log.Println("Total and Currency are required")
        c.JSON(http.StatusBadRequest, gin.H{"message": "Total and Currency are required"})
        return
    }

	paymentData := map[string]interface{}{
		"intent": "sale",
		"redirect_urls": map[string]string{
			"return_url": "http://127.0.0.1:5173/",
			"cancel_url": "http://127.0.0.1:5173/",
		},
		"payer": map[string]string{
			"payment_method": "paypal",
		},
		"transactions": []map[string]interface{}{
			{
				"amount": map[string]string{
					"total":    total,
					"currency": currency,
				},
				"description": "This is the payment transaction description.",
			},
		},
	}

    // Convert to json paymentData
	paymentJSON, err := json.Marshal(paymentData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to marshal payment data"})
		return
	}
    
     // POST Request to endpoint of PayPal API to create payment using paypal method
	body, err := services.PostRequestUsingPaypalMethod(c, paymentURL, accessToken, paymentJSON)


    // Convert JSON to golang map
	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to unmarshal JSON"})
		return
	}

    // Getting username by token
    username, err := services.ValidateToken(c)
    if err != nil{
        log.Println("Issues with token")
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to validate token"})
        return
    }
    
    // Init user repository to manage user
    user_repository, err := services.InitializeUserRepository()
    if err != nil{
        log.Println("Failed to init repository")
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to init db and repository"})
        return
    }

    // Get user by username
    user, err := user_repository.GetUserByUsername(username)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user"})
        return
    }

    // Converting string to float64
    convertedToFloatTotal, err := strconv.ParseFloat(total, 64)
    if err != nil{
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to convert string to float"})
        return
    }

    user.Balance += convertedToFloatTotal 
    err = user_repository.UpdateBalanceUser(user)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user balance"})
        return
    }

	c.JSON(http.StatusOK, result)
}


func CreateCreditCardPaymentHandler(c *gin.Context) {

    // Getting accessToken to do API request to PAYPAL
    accessToken, err := services.PaypalGetAccessToken()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to get access token"})
        return
    }

    paymentURL := "https://api.sandbox.paypal.com/v1/payments/payment"
    var paypalCardInput services.PaypalPaymentCardInput
    if err := c.BindJSON(&paypalCardInput); err != nil{
        log.Println("Failed to bind JSON for PaypalPaymentCardInput:", err)
        c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to bind JSON for PaypalPaymentCardInput"})
    }

    log.Println("Received data: ", paypalCardInput)

    // Getting needed data to do payment
    total, currency,
    numberCard, typeCard,
    expireMonthCard, expireYearCard,
    cvv2, firstName,
    lastName := services.ExtractCreditCardPaymentData(&paypalCardInput)

    paymentData := services.GetCreditCardPaymentData(
        total, currency,
        numberCard, typeCard,
        expireMonthCard, expireYearCard,
        cvv2, firstName, lastName,
    )

    // Convert to JSON
    paymentJSON, err := json.Marshal(paymentData)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to marshal payment data"})
        return
    }

    // POST Request to endpoint of PayPal API to create payment using credit card
    req, err := http.NewRequest(http.MethodPost, paymentURL, bytes.NewBuffer(paymentJSON))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"message": "Failed to create request"})
        return
    }

    // Set needed headers
    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("Authorization", "Bearer "+accessToken)

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"message": "Failed to send request"})
        return
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusCreated {
        body, _ := ioutil.ReadAll(resp.Body)
        c.JSON(resp.StatusCode, gin.H{"message": "Failed to create payment", "details": string(body)})
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

    // Getting username by token
    username, err := services.ValidateToken(c)
    if err != nil{
        log.Println("Issues with token")
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to validate token"})
        return
    }
    
    // Init user repository to manage user
    user_repository, err := services.InitializeUserRepository()
    if err != nil{
        log.Println("Failed to init repository")
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to init db and repository"})
        return
    }

    // Get user by username
    user, err := user_repository.GetUserByUsername(username)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user"})
        return
    }

    // Converting string to float64
    convertedToFloatTotal, err := strconv.ParseFloat(total, 64)
    if err != nil{
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to convert string to float"})
        return
    }

    user.Balance += convertedToFloatTotal 
    err = user_repository.UpdateBalanceUser(user)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user balance"})
        return
    }

    c.JSON(http.StatusOK, result)
}