package paypal_handlers

import (
    "log"
	"encoding/json"
	"net/http"
    "strconv"

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

//     paypalOrderURL := "https://api-m.sandbox.paypal.com/v2/checkout/orders"

func CreatePaymentOrder(c *gin.Context){
    // Getting paypal AccessToken
    accessToken, err := services.PGetAccessToken(c)
    if err != nil{
        return
    }
    log.Println("accessToken=", accessToken)
    paypalOrderURL := "https://api-m.paypal.com/v2/checkout/orders"
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
    moneyValueString := strconv.FormatFloat(moneyValue, 'f', -1, 64)
    orderData := services.GetOrderPaymentData(currencyCode, moneyValueString)

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
        Amount: moneyValue,
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

    paypalOrderURL := "https://api-m.paypal.com/v2/checkout/orders/" + orderID

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



func UpdatePaymentStatusToApproved(c *gin.Context){
    orderID := c.Query("token")
    if orderID == "" {
        c.JSON(http.StatusBadRequest, gin.H{"message": "Order ID is required"})
        return
    }

    // Getting access token
    accessToken, err := services.PGetAccessToken(c)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to get paypal access token"})
        return
    }

    paypalOrderURL := "https://api-m.paypal.com/v2/checkout/orders/" + orderID

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

    // Getting Status
    status := result["status"].(string)
    log.Println("OrderStatus", status)

    // If status was approved. We will update balance and status
    if status == "APPROVED"{
        log.Println("Status was approved!!!")
        services.UpdatePaymentStatus(c, orderID, status)
        c.JSON(http.StatusOK, gin.H{"message": "Updated successfully"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Failure to update balance and status"})

}

func PickUpMoneyAndStatusToSuccess(c *gin.Context) {
    var paypalPickUpMoneyInput services.PaypalPickUpMoneyInput
    if err := services.PBindJSONData(c, &paypalPickUpMoneyInput); err != nil {
        return
    }
    currentStatus := paypalPickUpMoneyInput.Status
    amountOfMoney := paypalPickUpMoneyInput.Total
    orderID := paypalPickUpMoneyInput.OrderID

    amountOfMoneyString := strconv.FormatFloat(amountOfMoney, 'f', -1, 64)

    // If status is approved. We're updating balance and status to success 
    if currentStatus == "APPROVED"{
        services.UpdateUserBalance(c, amountOfMoneyString)
        services.UpdatePaymentStatus(c, orderID, "Success")
    }
}


func WithdrawFundsPaypal(c *gin.Context) {
    // Getting paypal access token
    accessToken, err := services.PGetAccessToken(c)
    if err != nil{
        log.Println("Failed to get paypal access token")
        c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to get paypal access token"})
        return
    }
    // paypal endpoint URL to withdraw funds 
    paypalWithdrawFundsURL := "https://api-m.paypal.com/v1/payments/payouts"

    var paypalWithdrawInput services.PaypalWithdrawFundsInput
    if err := services.PBindJSONData(c, &paypalWithdrawInput); err != nil{
        return
    }

    // Getting the required data from JSON
    total := paypalWithdrawInput.Total
    currency := paypalWithdrawInput.Currency
    receiverEmail := paypalWithdrawInput.ReceiverEmail

    // Minimum withdrawl amount
    if total < 100.00 {
        c.JSON(http.StatusInternalServerError, gin.H{"message": "Minimum withdrawal amount is 100$"})
    }
    
    // Converting float64 to string
    totalString := strconv.FormatFloat(total, 'f', -1, 64)

    // Logging recieved data
    log.Println("Total: ", total)
    log.Println("Currency: ", currency)


    // Getting needed data for POST request
    withdrawFundsData := services.GetWithdrawFundsData(currency, totalString, receiverEmail)

    // Converting data for POST request to JSON
    withdrawFundsDataJSON, err := json.Marshal(withdrawFundsData)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to marshal withdraw funds JSON data "})
        return
    }

    // POST request to withdraw funds
    response, err := services.PostRequestWithdrawFunds(c, paypalWithdrawFundsURL, accessToken, withdrawFundsDataJSON)
    if err != nil{
        c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to complete the POST request for withdrawal of funds"})
        return
    }

    // Handling response and getting golang object to work with it
    result, err := services.PHandlePaypalResponse(c, response)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to handle paypal response"})
        return
    }

    // Take the balance from the site
    
    services.UpdateNegativeUserBalance(c, totalString)

    c.JSON(http.StatusOK, result)
}