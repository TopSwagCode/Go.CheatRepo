package notificationservices

import "fmt"

// SmsSender is used to setup a SmsSender to Send Sms.
type SmsSender struct {
	SmsGateway string
}

// SendSms sends an Sms to the number
func (s SmsSender) SendSms(number int, message string) {
	fmt.Println(message)
}
