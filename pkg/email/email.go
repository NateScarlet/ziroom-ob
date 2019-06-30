package email

import (
	"errors"
	"os"
	"strconv"

	"gopkg.in/gomail.v2"
)

// Send mail to configured user
func Send(subject string, text string) (err error) {

	host := os.Getenv("EMAIL_HOST")
	portStr := os.Getenv("EMAIL_PORT")
	user := os.Getenv("EMAIL_USER")
	password := os.Getenv("EMAIL_PASSWORD")
	fromAddr := os.Getenv("EMAIL_FROM")
	toAddr := os.Getenv("EMAIL_TO")

	if fromAddr == "" {
		return errors.New("Not configured environment variable: `EMAIL_FROM`")
	}
	if host == "" {
		return errors.New("Not configured environment variable: `EMAIL_HOST`")
	}
	if portStr == "" {
		return errors.New("Not configured environment variable: `EMAIL_PORT`")
	}
	if password == "" {
		return errors.New("Not configured environment variable: `EMAIL_PASSWORD`")
	}
	if user == "" {
		user = fromAddr
	}
	if toAddr == "" {
		toAddr = fromAddr
	}
	port, err := strconv.Atoi(portStr)
	if err != nil {
		return
	}

	msg := gomail.NewMessage()
	msg.SetAddressHeader("From", fromAddr, "自如抢房监控")
	msg.SetHeader("To", toAddr)
	msg.SetHeader("Subject", "抢房监控: "+subject)
	msg.SetBody("text/plain", text)

	dialer := gomail.NewDialer(host, port, user, password)
	err = dialer.DialAndSend(msg)
	return
}
