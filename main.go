package main

import (
	"fmt"
	"log"
	"net/smtp"
	"strings"
)

const SMTP_HOST = "yoursmtp@tld.com"
const SMTP_PORT = 0
const SENDER_NAME = "yourname <youremail@tld.com"
const EMAIL = "youremail@tld.com"
const PASSWORD = "yourpassword"

func main(){
	to := []string{"recipient@tld.com", "recipient1@tld.com"}
	cc := []string{"recipientcc@tld.com"}
	subject := "your subject"
	message := "Hello"

	err := sendMail(to, cc, subject, message)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println("Mail sent")
}

func sendMail(to []string, cc []string, subject, message string) error {
	body := "From: " + SENDER_NAME + "\n" +
			"To: " + strings.Join(to, ",") + "\n" +
			"Subject: " + subject + "\n\n" +
			message

	auth := smtp.PlainAuth("", EMAIL, PASSWORD, SMTP_HOST)
	smtpAddr := fmt.Sprintf("%s:%d", SMTP_HOST, SMTP_PORT)

	err := smtp.SendMail(smtpAddr, auth, EMAIL, append(to, cc...), []byte(body))
	if err != nil {
		return err
	}
	return nil
}