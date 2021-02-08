package utils

import (
	"net/smtp"
)

//SendEmail sends an email
func SendEmail(to []string, subject string, body string) error {

	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"

	auth := smtp.PlainAuth("", "ad3fe6ed2861cb", "51f3038b85118a", "smtp.mailtrap.io")

	err := smtp.SendMail("smtp.mailtrap.io:2525", auth, "no-reply@gmail.com", to, []byte(subject+mime+body))
	if err != nil {
		return err
	}

	return nil
}
