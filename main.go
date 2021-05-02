package main

import (
	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/checkout/session"
)

func main() {
	stripe.Key = "sk_test_51ImPFCK5S3kVHXO0QNDoeLJ2FwPBlK4T2X1wNhcMOKKuYNfSVCB8kNZ76z3SarK12P3PMdoGn0qDhYhsXl3dFZqL00OL8nZy25"

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.File("index.html")
	})
	r.POST("/checkout", func(c *gin.Context) {
		params := &stripe.CheckoutSessionParams{
			PaymentMethodTypes: stripe.StringSlice([]string{
				"card",
			}),
			Mode: stripe.String(string(stripe.CheckoutSessionModePayment)),
			LineItems: []*stripe.CheckoutSessionLineItemParams{
				{
					PriceData: &stripe.CheckoutSessionLineItemPriceDataParams{
						Currency: stripe.String("usd"),
						ProductData: &stripe.CheckoutSessionLineItemPriceDataProductDataParams{
							Name: stripe.String("Corn Puff Variety Pack"),
						},
						UnitAmount: stripe.Int64(3000),
					},
					Quantity: stripe.Int64(1),
				},
			},
			SuccessURL: stripe.String("https://cornpuffs.store/success"),
			CancelURL:  stripe.String("https://cornpuffs.store/cancel"),
		}

		session, err := session.New(params)

		if err != nil {
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(200, gin.H{
			"sessionID": session.ID,
		})
	})
	r.Run()
}
