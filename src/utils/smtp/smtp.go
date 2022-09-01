package smtpMail

import (
	"errors"
	"net/smtp"
	"os"
	"strconv"
)

type loginAuth struct {
	username, password string
}

func LoginAuth(username, password string) smtp.Auth {
	return &loginAuth{username, password}
}

func (a *loginAuth) Start(server *smtp.ServerInfo) (string, []byte, error) {
	return "LOGIN", []byte{}, nil
}

func (a *loginAuth) Next(fromServer []byte, more bool) ([]byte, error) {

	if more {

		switch string(fromServer) {

		case "Username:":
			return []byte(a.username), nil
		case "Password:":
			return []byte(a.password), nil
		default:
			return nil, errors.New("Unkown fromServer")

		}
	}

	return nil, nil

}

func SendMail(secret int, url string, emailAddress string) error {

	auth := LoginAuth(os.Getenv("SMTP_EMAIL"), os.Getenv("SMTP_PASSWORD"))

	to := []string{emailAddress}

	subject := "Subject: Confirn your email\n"
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body := `<html lang="en">
	<head>
	  <meta charset="utf-8" />
	</head>
	<body>` +
		"confirmation link : " + os.Getenv("URL_SERVER_FRONT") + "?code=" + url + "\r\n" + "<br>" +
		"your secret code : " + strconv.Itoa(secret) + "\r\n" + "<br>" +
		`</body>
  	</html>`

	msg := []byte(subject + mime + body)

	err := smtp.SendMail("smtp.gmail.com:587", auth, os.Getenv("SMTP_EMAIL"), to, msg)

	if err != nil {

		return err

	}

	return nil

}
