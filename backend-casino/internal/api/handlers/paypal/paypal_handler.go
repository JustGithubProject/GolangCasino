package paypal_handlers

import (
    "log"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"

    "github.com/JustGithubProject/GolangCasino/backend-casino/internal/services"
)



func CreatePaypalPaymentHandler(c *gin.Context) {
    accessToken, err := services.PGetAccessToken(c)
    if err != nil {
        return
    }

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