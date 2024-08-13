package paypal_handlers

import (
    "log"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"

    "github.com/JustGithubProject/GolangCasino/backend-casino/internal/services"
    "github.com/JustGithubProject/GolangCasino/backend-casino/internal/models"
)

/*
To do payment:
    After creating an order using the /v2/checkout/orders endpoint,
    you will receive an order ID and a confirmation link.
    This link is the URL that the customer must go to to complete the payment process.
*/

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

    // TODO: fix it. You need to somehow update the balance only after the order is approved
    // services.UpdateUserBalance(c, moneyValue)

    /*
        BLOCK WITH GETTING OBJECT OF USER (WILL BE MOVE SOON)
    */
    username, err := services.ValidateToken(c)
    if err != nil{
        log.Println("Issues with casino auth token")
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
    /*
        END BLOCK
    */
    paymentRepository, err := services.InitializePaymentRepository()
    if err != nil{
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to init db and repository"})
        return
    }
    payment := models.Payment{
        OrderID: result["id"].(string),
        UserID: user.ID,
        Status: "Pending",
    }

    err = paymentRepository.CreatePayment(&payment)
    if err != nil{
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create payment"})
        return 
    }


    c.JSON(http.StatusOK, result)
}

func InitializePaymentRepository() {
	panic("unimplemented")
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

func GetListPaypalPayments(c *gin.Context){
    /*
        BLOCK WITH GETTING OBJECT OF USER (WILL BE MOVE SOON)
    */
    username, err := services.ValidateToken(c)
    if err != nil{
        log.Println("Issues with casino auth token")
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
    /*
        END BLOCK
    */

    userWithPayments, err := user_repository.GetUserPayments(user.ID)
    if err != nil{
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get payments of user"})
    }
    
    c.JSON(http.StatusOK, userWithPayments.Payments)
}