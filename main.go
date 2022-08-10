package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ijasmoopan/sendsms_twilio/handlers"
)

func main() {
	http.HandleFunc("/", handlers.App)
	http.HandleFunc("/sendsms", handlers.SendSMS)
	fmt.Println("Server starting on the port 8092")
	log.Fatal(http.ListenAndServe(":8092", nil))
}

