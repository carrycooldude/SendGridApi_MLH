package main

import (
	"fmt"
	"log"
	"os"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func main() {
	from := mail.NewEmail("me", "carrycooldude@gmails.com")
	to := mail.NewEmail("Kartikey", "rawatkari554@gmail.com")
	subject := "MLH Sendgrid API..."
	plainTextContent := "and easy to do anywhere , even with Go"
	htmlContent := "<strong>and yaas this is Kartikey aka Carrycooldude"
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	response, err := client.Send(message)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}
