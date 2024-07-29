package paypal_handlers

import (
	"bytes"
    "log"
	"encoding/json"
    "encoding/base64"
    "io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

// Getting vars from .env file
var clientID string = os.Getenv("PAYPAL_CLIENT_ID")
var paypalSecret string = os.Getenv("PAYPAL_SECRET")


func PaypalGetAccessToken() (string, error) {
	paypalURL := "https://api.sandbox.paypal.com/v1/oauth2/token"
	data := url.Values{}
	auth := base64.StdEncoding.EncodeToString([]byte(clientID + ":" + paypalSecret))
	data.Set("grant_type", "client_credentials")

	req, err := http.NewRequest(http.MethodPost, paypalURL, strings.NewReader(data.Encode()))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", "Basic "+auth)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return "", err
	}

	return result["access_token"].(string), nil
}


func CreatePaypalPaymentHandler(c *gin.Context) {
    // Getting accessToken to do API request to PAYPAL
	accessToken, err := PaypalGetAccessToken()
    log.Println("AcessToken: ", accessToken)
	if err != nil {
        log.Println("AccessToken is empty or invalid")
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to get access token"})
		return
	}
    log.Println("AccessToken is okay")
	paymentURL := "https://api.sandbox.paypal.com/v1/payments/payment"

    // Getting amount of money and currency to execute payment using paypal method
    total := c.PostForm("Total")
    currency := c.PostForm("Currency")

    log.Println("Total: ", total)
    log.Println("Currency: ", currency)

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
	req, err := http.NewRequest(http.MethodPost, paymentURL, strings.NewReader(string(paymentJSON)))
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


    // Convert JSON to golang map
	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to unmarshal JSON"})
		return
	}

	c.JSON(http.StatusOK, result)
}



func CreateCreditCardPaymentHandler(c *gin.Context) {

    // Getting accessToken to do API request to PAYPAL
    accessToken, err := PaypalGetAccessToken()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to get access token"})
        return
    }

    paymentURL := "https://api.sandbox.paypal.com/v1/payments/payment"

    // Getting needed data to do payment
    numberCard := c.PostForm("numberCard")
    typeCard := c.PostForm("typeCard")
    expireMonthCard := c.PostForm("expireMonthCard")
    expireYearCard := c.PostForm("expireYearCard")
    cvv2 := c.PostForm("cvv2")
    firstName := c.PostForm("firstName")
    lastName := c.PostForm("lastName")
    total := c.PostForm("total")
    currency := c.PostForm("currency")

    paymentData := map[string]interface{}{
        "intent": "sale",
        "payer": map[string]interface{}{
            "payment_method": "credit_card",
            "funding_instruments": []map[string]interface{}{
                {
                    "credit_card": map[string]interface{}{
                        "number":       numberCard,
                        "type":        typeCard,
                        "expire_month": expireMonthCard,
                        "expire_year":  expireYearCard,
                        "cvv2":         cvv2,
                        "first_name":   firstName,
                        "last_name":    lastName,
                        "billing_address": map[string]string{
                            "line1":       "52 N Main ST",
                            "city":        "Johnstown",
                            "state":       "OH",
                            "postal_code": "43210",
                            "country_code": "US",
                        },
                    },
                },
            },
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

    c.JSON(http.StatusOK, result)
}