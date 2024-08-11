package paypal_handlers

import (
    "log"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"

    "github.com/JustGithubProject/GolangCasino/backend-casino/internal/services"
)

/*
To do payment:
    After creating an order using the /v2/checkout/orders endpoint,
    you will receive an order ID and a confirmation link.
    This link is the URL that the customer must go to to complete the payment process.
*/


func CreatePaypalPaymentHandler(c *gin.Context) {

    // Getting paypal AccessToken
    accessToken, err := services.PGetAccessToken(c)
    if err != nil {
        return
    }
    log.Println("AccessToken: ", accessToken)

    /*
        Sandbox. https://api-m.sandbox.paypal.com
        Live. https://api-m.paypal.com
    */
    paymentURL := "https://api.sandbox.paypal.com/v1/payments/payment"

    var paypalInput services.PaypalPaymentInput
    if err := services.PBindJSONData(c, &paypalInput); err != nil {
        return
    }

    log.Println("Received data: ", paypalInput)

    total, currency := services.ExtractPaypalPaymentData(&paypalInput)

    if total == "" || currency == "" {
        log.Println("Total and Currency are required")
        c.JSON(http.StatusBadRequest, gin.H{"message": "Total and Currency are required"})
        return
    }

    paymentData := services.GetPaypalPaymentData(total, currency)

    paymentJSON, err := json.Marshal(paymentData)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to marshal payment data"})
        return
    }

    body, err := services.PPostPaypalRequest(c, paymentURL, accessToken, paymentJSON)
    if err != nil {
        return
    }

    result, err := services.PHandlePaypalResponse(c, body)
    if err != nil {
        return
    }

    services.UpdateUserBalance(c, total)
    c.JSON(http.StatusOK, result)
}

func CreateCreditCardPaymentHandler(c *gin.Context) {
    // Getting paypal AccessToken
    accessToken, err := services.PGetAccessToken(c)
    if err != nil {
        return
    }

    paymentURL := "https://api.sandbox.paypal.com/v1/payments/payment"

    var paypalCardInput services.PaypalPaymentCardInput
    if err := services.PBindJSONData(c, &paypalCardInput); err != nil {
        return
    }

    log.Println("Received data: ", paypalCardInput)

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

    paymentJSON, err := json.Marshal(paymentData)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to marshal payment data"})
        return
    }

    body, err := services.PPostPaypalRequest(c, paymentURL, accessToken, paymentJSON)
    if err != nil {
        return
    }

    result, err := services.PHandlePaypalResponse(c, body)
    if err != nil {
        return
    }

    services.UpdateUserBalance(c, total)
    c.JSON(http.StatusOK, result)
}


func CreatePaymentOrder(c *gin.Context){
    // Getting paypal AccessToken
    accessToken, err := services.PGetAccessToken(c)
    if err != nil{
        return
    }
    log.Println("accessToken=", accessToken)
    paypalOrderURL := "https://api-m.sandbox.paypal.com/v2/checkout/orders"
    var paypalCreateOrderInput services.PaypalCreateOrderInput
    if err := services.PBindJSONData(c, &paypalCreateOrderInput); err != nil {
        return
    }

    // Getting currency and amount of money from JSON
    currencyCode := paypalCreateOrderInput.CurrencyCode
    moneyValue := paypalCreateOrderInput.MoneyValue

    log.Println("CurrencyCode=", currencyCode)
    log.Println("MonetValue=", moneyValue)

    // Needed data to do request
    orderData := services.GetOrderPaymentData(currencyCode, moneyValue)

    log.Println("OrderData", orderData["intent"])
    orderJSON, err := json.Marshal(orderData)
    if err != nil{
        c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to marshal order JSON data"})
        return
    }
    
    log.Println("JSON: ", orderJSON)
    
    body, err := services.PPostPaypalCreateOrderRequest(c, paypalOrderURL, accessToken, orderJSON)
    if err != nil{
        // Dropped here
        log.Println("Error occurred after PPostPaypalCreateOrderRequest")
        c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to create paypal order"})
        return
    }

    result, err := services.PHandlePaypalResponse(c, body)
    if err != nil {
        log.Println("Error occurred after PHandlePaypalResponse")
        return
    }

    log.Println("order_id", result["id"])

    services.UpdateUserBalance(c, moneyValue)
    c.JSON(http.StatusOK, result)
}


func GetOrderDetailByID(c *gin.Context) {
    // Getting order id
    orderID := c.Query("id")
    if orderID == "" {
        c.JSON(http.StatusBadRequest, gin.H{"message": "Order ID is required"})
        return
    }

    // Getting access token
    accessToken, err := services.PGetAccessToken(c)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to get access token"})
        return
    }

    paypalOrderURL := "https://api-m.sandbox.paypal.com/v2/checkout/orders/" + orderID

    // Do get request to PayPal API and get response
    response, err := services.PGetPaypalOrderDetails(c, paypalOrderURL, accessToken)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to get PayPal order details"})
        return
    }

    result, err := services.PHandlePaypalResponse(c, response)
    if err != nil {
        return
    }
    c.JSON(http.StatusOK, result)
}