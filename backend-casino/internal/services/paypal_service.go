package services

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

///////////////////////////////////////
//	BLOCK WITH PAYPAL STRUCTS	     //
///////////////////////////////////////


type PaypalPaymentInput struct {
    Total    string `json:"Total"`
    Currency string `json:"Currency"`
}

type PaypalPaymentCardInput struct {
    NumberCard string `json:"NumberCard"`
    TypeCard string `json:"TypeCard"`
    ExpireMonthCard string `json:"ExpireMonthCard"`
    ExpireYearCard string `json:"ExpireYearCard"`
    CVV2 string `json:"CVV2"`
    FirstName string `json:"FirstName"`
    LastName string `json:"LastName"`
    Total string `json:"Total"`
    Currency string `json:"Currency"`
}

//////////////////////////////////////
/////////////////////////////////////
////////////////////////////////////


func PaypalGetAccessToken() (string, error) {
    /*
    An access token is typically used to perform operations on behalf of the application,
    such as creating payments, checking the status of transactions, etc.
    */

    err := godotenv.Load()
    if err != nil{
        log.Fatal("Error loading .env file")
    }

    // Getting vars from .env file
    var clientID string = os.Getenv("PAYPAL_CLIENT_ID")
    var paypalSecret string = os.Getenv("PAYPAL_SECRET")
    
	paypalURL := "https://api-m.sandbox.paypal.com/v1/oauth2/token"
	data := url.Values{}
	auth := base64.StdEncoding.EncodeToString([]byte(clientID + ":" + paypalSecret))


    /*
    grant_type=???
    For example, when access to the API is required on behalf of a specific user,
    other authorization methods such as auth_code or password are used,
    where the user directly interacts with the authorization system.
    */
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


func ExtractPaypalPaymentData(data *PaypalPaymentInput) (string, string){
	return data.Total, data.Currency
}


func ExtractCreditCardPaymentData(data *PaypalPaymentCardInput) (string, string, string, string, string, string, string, string, string){
	return data.Total, data.Currency, data.NumberCard, data.TypeCard, data.ExpireMonthCard, data.ExpireYearCard, data.CVV2, data.FirstName, data.LastName
}


func GetCreditCardPaymentData(
    total, currency, numberCard, typeCard, expireMonthCard, expireYearCard, cvv2, firstName, lastName string,
) map[string]interface{} {
    paymentData := map[string]interface{}{
        "intent": "sale",
        "payer": map[string]interface{}{
            "payment_method": "credit_card",
            "funding_instruments": []map[string]interface{}{
                {
                    "credit_card": map[string]interface{}{
                        "number":        numberCard,
                        "type":          typeCard,
                        "expire_month":  expireMonthCard,
                        "expire_year":   expireYearCard,
                        "cvv2":          cvv2,
                        "first_name":    firstName,
                        "last_name":     lastName,
                        "billing_address": map[string]string{
                            "line1":        "52 N Main ST",
                            "city":         "Johnstown",
                            "state":        "OH",
                            "postal_code":  "43210",
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
    return paymentData
}


func GetPaypalPaymentData(total, currency string) map[string]interface{}{
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
	return paymentData
}


func PostRequestUsingPaypalMethod(c *gin.Context, paymentURL string, accessToken string, paymentJSON []byte) ([]byte, error) {
	req, err := http.NewRequest(http.MethodPost, paymentURL, strings.NewReader(string(paymentJSON)))
	if err != nil {
        log.Println("Failed to create request")
		c.JSON(http.StatusBadRequest, gin.H{"message": "Failed to create request"})
		return nil, err
	}

    // Set needed headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+accessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
        log.Println("Failed to send request")
		c.JSON(http.StatusBadRequest, gin.H{"message": "Failed to send request"})
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		body, _ := ioutil.ReadAll(resp.Body)
		c.JSON(resp.StatusCode, gin.H{"message": "Failed to create payment", "details": string(body)})
		return nil, errors.New("failed to create payment")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
        log.Println("Failed to read response body")
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to read response body"})
		return nil, err
	}
	return body, nil
}



func UpdateUserBalance(c *gin.Context, total string){
	// Getting username by token
    username, err := ValidateToken(c)
    if err != nil{
        log.Println("Issues with token")
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to validate token"})
        return
    }
    
    // Init user repository to manage user
    user_repository, err := InitializeUserRepository()
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
}


func PGetAccessToken(c *gin.Context) (string, error) {
    accessToken, err := PaypalGetAccessToken()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to get access token"})
        return "", err
    }
    return accessToken, nil
}

func PBindJSONData(c *gin.Context, input interface{}) error {
    if err := c.BindJSON(input); err != nil {
        log.Println("Failed to bind JSON:", err)
        c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input"})
        return err
    }
    return nil
}

func PPostPaypalRequest(c *gin.Context, paymentURL, accessToken string, paymentJSON []byte) ([]byte, error) {
    body, err := PostRequestUsingPaypalMethod(c, paymentURL, accessToken, paymentJSON)
    if err != nil {
        log.Println("Post request using paypal method failed!!!")
        c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to do POST request"})
        return nil, err
    }
    return body, nil
}

func PHandlePaypalResponse(c *gin.Context, body []byte) (map[string]interface{}, error) {
    var result map[string]interface{}
    err := json.Unmarshal(body, &result)
    if err != nil {
        log.Println("Failed to unmarshal JSON")
        c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to unmarshal JSON"})
        return nil, err
    }
    return result, nil
}