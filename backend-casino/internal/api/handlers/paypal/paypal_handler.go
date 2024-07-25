package paypal_handlers

import (
    "github.com/gin-gonic/gin"
    "github.com/plutov/paypal/v4"
    "log"
    "net/http"
    "os"
)

var clientID string = os.Getenv("PAYPAL_CLIENT_ID")
var paypalSecret string = os.Getenv("PAYPAL_SECRET")


func CreatePayment(c *gin.Context) {

	// Getting vars from .env file
	clientID := os.Getenv("PAYPAL_CLIENT_ID")
	paypalSecret := os.Getenv("PAYPAL_SECRET")

    client, err := paypal.NewClient(clientID, paypalSecret, paypal.APIBaseSandBox)
    if err != nil {
        log.Fatal(err)
    }
    client.SetLog(os.Stdout)

    order, err := client.CreateOrder(paypal.Order{
        Intent: "CAPTURE",
        PurchaseUnits: []paypal.PurchaseUnit{
            {
                Amount: &paypal.Amount{
                    Currency: "USD",
                    Total:    "10.00", 
                },
            },
        },
        ApplicationContext: &paypal.ApplicationContext{
            ReturnUrl: "http://localhost:8080/execute-payment",
            CancelUrl: "http://localhost:8080",
        },
    })
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    for _, link := range order.Links {
        if link.Rel == "approve" {
            c.Redirect(http.StatusFound, link.Href)
            return
        }
    }
    c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to create payment"})
}

func ExecutePayment(c *gin.Context) {
    client, err := paypal.NewClient("YOUR_CLIENT_ID", "YOUR_CLIENT_SECRET", paypal.APIBaseSandBox)
    if err != nil {
        log.Fatal(err)
    }
    client.SetLog(os.Stdout)

    orderID := c.Query("token")
    if orderID == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Order ID not provided"})
        return
    }

    capture, err := client.CaptureOrder(orderID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

	// TODO: update balance of user (model)
    c.JSON(http.StatusOK, gin.H{"message": "Payment successful", "capture": capture})
}
