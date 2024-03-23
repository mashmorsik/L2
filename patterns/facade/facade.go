package main

// NotificationFacade предоставляет простой интерфейс для отправки уведомлений
type NotificationFacade struct {
	emailNotification     *EmailNotification
	smsNotification       *SMSNotification
	messengerNotification *MessengerNotification
}

func NewNotificationFacade() *NotificationFacade {
	return &NotificationFacade{
		emailNotification:     &EmailNotification{},
		smsNotification:       &SMSNotification{},
		messengerNotification: &MessengerNotification{},
	}
}

func (nf *NotificationFacade) SendNotification(message string) {
	nf.emailNotification.SendEmailNotification(message)
	nf.smsNotification.SendSMSNotification(message)
	nf.messengerNotification.SendMessengerNotification(message)
}
