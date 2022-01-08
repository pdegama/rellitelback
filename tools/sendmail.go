package tools

import (
	"fmt"
	"net/smtp"
)

func SendMail(email string, subject string, message string) {

	from := "com.the.noob.team@gmail.com"
	password := "Hello@1234"

	to := []string{
		email,
	}

	//fmt.Println(to)

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	sub := "Subject: " + subject + "!\n"
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	
	msg := []byte(sub + mime + message)

	auth := smtp.PlainAuth("", from, password, smtpHost)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, msg)

	if err != nil {
		fmt.Println(err)
		return
	}

}
