package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	razorpay "github.com/razorpay/razorpay-go"
)

type Order struct {
	ID interface{}
	Entity interface{}
	Amount interface{}
	AmountPaid interface{}
	AmountDue interface{}
	Currency interface{}
	Receipt interface{}
	OfferID interface{}
	Status interface{}
	Attempts interface{}
	Notes interface{}
	CreatedAt interface{}
}

func main(){
	http.HandleFunc("/", App)
	log.Fatal(http.ListenAndServe(":8090", nil))
}

func App(w http.ResponseWriter, r *http.Request){

	err := godotenv.Load("./.env")
	if err != nil {
		log.Println("Can't open env file")
	}

	// Adding razorpay credentials
	client := razorpay.NewClient(os.Getenv("RAZORPAY_KEY_ID"), os.Getenv("RAZORPAY_SECRET"))

	data := map[string]interface{}{
		"amount": 99880,
		"currency": "INR",
		"receipt": "some_receipt_id",
	}
	body, err := client.Order.Create(data, nil)
	if err != nil {
		log.Println("Can't create order")
		os.Exit(1)
	}
	log.Println("Order created")

	var order Order
	order.ID = body["id"]
	order.Entity = body["entity"]
	order.Amount = body["amount"]
	order.AmountPaid = body["amount_paid"]
	order.AmountDue = body["amount_due"]
	order.Currency = body["currency"]
	order.Receipt = body["receipt"]
	order.OfferID = body["offer_id"]
	order.Status = body["status"]
	order.Attempts = body["attempts"]
	order.Notes = body["notes"]
	order.CreatedAt = body["created_at"]

	fmt.Println("Order Details:", order)

	t, err := template.ParseFiles("app.html")
	if err != nil {
		log.Println("Can't open html file")
	}
	err = t.Execute(w, order)
	if err != nil {
		log.Println("Can't open html file")
	}
}