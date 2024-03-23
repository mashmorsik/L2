package main

import "fmt"

// EmailNotification представляет функциональность отправки уведомлений по электронной почте
type EmailNotification struct{}

func (en *EmailNotification) SendEmailNotification(message string) {
	fmt.Println("Sending email notification:", message)
}

// SMSNotification представляет функциональность отправки уведомлений по SMS
type SMSNotification struct{}

func (sn *SMSNotification) SendSMSNotification(message string) {
	fmt.Println("Sending SMS notification:", message)
}

// MessengerNotification представляет функциональность отправки уведомлений через мессенджеры
type MessengerNotification struct{}

func (mn *MessengerNotification) SendMessengerNotification(message string) {
	fmt.Println("Sending messenger notification:", message)
}
