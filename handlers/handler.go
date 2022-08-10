package handlers

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
)

func App(w http.ResponseWriter, r *http.Request) {
	template := `<html><body><a href="/sendsms">Send Message</a></body></html>`
	fmt.Fprintf(w, template)
}
func SendSMS(w http.ResponseWriter, r *http.Request) {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Can't access env file")
	}

	accountSid := os.Getenv("ACCOUNT_SID")
	authToken := os.Getenv("AUTH_TOKEN")
	phoneNumber := os.Getenv("TWILIO_PHONE_NUMBER")

	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSid,
		Password: authToken,
	})

	params := &openapi.CreateMessageParams{}
	params.SetTo("+917034464400")
	params.SetFrom(phoneNumber)
	params.SetBody("Hello from Go using Twilio")

	resp, err := client.Api.CreateMessage(params)
	if err != nil {
		fmt.Println("Error:", err.Error())
		err = nil
	} else {
		fmt.Println("Message sent")
		fmt.Println("Message Sid:", *resp.Sid)
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
