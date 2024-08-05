package services

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"os"
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

    err := godotenv.Load()
    if err != nil{
        log.Fatal("Error loading .env file")
    }

    // Getting vars from .env file
    var clientID string = os.Getenv("PAYPAL_CLIENT_ID")
    var paypalSecret string = os.Getenv("PAYPAL_SECRET")
    
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


func PostRequestUsingPaypalMethod(c *gin.Context, paymentURL string, accessToken string, paymentJSON []byte) ([]byte, error) {
	req, err := http.NewRequest(http.MethodPost, paymentURL, strings.NewReader(string(paymentJSON)))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Failed to create request"})
		return nil, err
	}

    // Set needed headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+accessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
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
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to read response body"})
		return nil, err
	}
	return body, nil
}
