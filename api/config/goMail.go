package config

import (
	"fmt"
	"os"
	"time"

	"gopkg.in/mail.v2"
)

func GoMail(correo string, asunto string, htmlContent string) {
	email_password := os.Getenv("EMAIL_PASSWORD")

	email := mail.NewMessage()
	email.SetHeader("From", "costabrisaweb@gmail.com")
	email.SetHeader("To", correo)
	email.SetHeader("Subject", asunto)
	email.SetBody("text/html", htmlContent)

	// Configuración del servidor SMTP
	smtpServer := mail.NewDialer("smtp.gmail.com", 587, "costabrisaweb@gmail.com", email_password)

	// Intentar enviar el correo electrónico
	err := smtpServer.DialAndSend(email)
	if err != nil {
		fmt.Println("Error:", err.Error())
		return
	}

	fmt.Println("Correo electrónico enviado correctamente")
}

func GetDate() string {
	now := time.Now()

	// Formatear la fecha a YYYY-MM-DD
	formattedDate := now.Format("2006-01-02")

	return formattedDate
}
func GetHour() string {
	now := time.Now()

	// Formatear la hora a HH:MM:SS
	formattedHour := now.Format("15:04:05")

	return formattedHour

}
