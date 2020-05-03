package NotificationServices

import "fmt"

type SmsSender struct {
	SmsGateway string
}

func (s SmsSender) SendSms(number int, message string) {
	fmt.Println(message)
}
