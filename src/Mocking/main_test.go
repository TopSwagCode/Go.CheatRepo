package main

import (
	"./notificationservices"
	"strconv"
	"strings"
	"testing"
)

// Internal Mocking
func TestWelcomeEmail(t *testing.T) {

	// Arrange
	var sb strings.Builder

	sender := MockMailSender{
		SendFunc: func(to, subject, body string) error {
			sb.WriteString(to)
			sb.WriteString(subject)
			sb.WriteString(body)

			return nil
		},
		SendFromFunc: nil,
	}

	// Act
	SendWelcomeEmail(sender, "josh@topswagcode.com ", "Go Is Awesome ", "Rant about why Go is awesome")
	msg := sb.String()

	// Assert
	if msg != "josh@topswagcode.com Go Is Awesome Rant about why Go is awesome" {
		t.Error("SendWelcomeEmail:", msg)
	}
}

func SendWelcomeEmail(sender MailSender, to string, subject string, body string) {
	// Bla bla bla do some stuff.
	sender.Send(to, subject, body)
}

type MailSender interface {
	Send(to, subject, body string) error
	SendFrom(from, to, subject, body string) error
}

type MockMailSender struct {
	SendFunc     func(to, subject, body string) error
	SendFromFunc func(from, to, subject, body string) error
}

func (m MockMailSender) Send(to, subject, body string) error {
	return m.SendFunc(to, subject, body)
}

func (m MockMailSender) SendFrom(from, to, subject, body string) error {
	return m.SendFromFunc(from, to, subject, body)
}

// External mocking
func TestWelcomeSms(t *testing.T) {

	// Creating both Mock and Real to show that we can create interface for other packages
	smsSender := notificationservices.SmsSender{SmsGateway: "Some gateway far far away"}

	var sb strings.Builder

	mockSender := MockSmsSender{
		SendSmsFunc: func(number int, message string) {
			sb.WriteString(strconv.Itoa(number))
			sb.WriteString(" ")
			sb.WriteString(message)
		},
	}

	// SendWelcomeSms works with both our Mock and the external package without a interface.
	SendWelcomeSms(smsSender, 122, "122")
	SendWelcomeSms(mockSender, 122, "122")
	msg := sb.String()

	mockSender.SendSms(122, "122")

	if msg != "122 122" {
		t.Error("SendSms:", msg)
	}
}

// Our SendWelcomeSms function.
func SendWelcomeSms(sender SmsSender, number int, message string) {
	// Bla bla bla do some stuff.
	sender.SendSms(number, message)
}

type SmsSender interface {
	SendSms(number int, message string)
}

type MockSmsSender struct {
	SendSmsFunc func(number int, message string)
}

func (m MockSmsSender) SendSms(number int, message string) {
	m.SendSmsFunc(number, message)
}
