package stripe_handlers

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/JustGithubProject/GolangCasino/backend-casino/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/stripe/stripe-go"
    "github.com/stripe/stripe-go/paymentintent"
)

func CreatePaymentIntent(c *gin.Context) {
	err := godotenv.Load()
    if err != nil{
        log.Fatal("Error loading .env file")
    }
	stripe.Key = os.Getenv("STRIPE_SECRET_KEY")
	
	var stripeCreatePaymentIntentInput services.StripeCreatePaymentIntentInput
    if err := services.PBindJSONData(c, &stripeCreatePaymentIntentInput); err != nil {
        return
    }

	// currency := stripeCreatePaymentIntentInput.Currency
	amount := stripeCreatePaymentIntentInput.Amount

	amountINT64, err := strconv.ParseInt(amount, 10, 64)
	if err != nil {
		log.Println("Failed to convert string to integer64")
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to convert string to integer64"})
		return
	}

	params := &stripe.PaymentIntentParams{
		Amount: stripe.Int64(amountINT64),
		Currency: stripe.String(string(stripe.CurrencyUSD)),
	}

	pi, err := paymentintent.New(params)
	if err != nil {
		log.Println("Error when generating a payment plan:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error when generating a payment plan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"client_secret": pi.ClientSecret})
}
