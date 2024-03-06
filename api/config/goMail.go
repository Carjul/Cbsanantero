package config

import (
	"fmt"

	"gopkg.in/mail.v2"
)

func GoMail() {
	email := mail.NewMessage()
	email.SetHeader("From", "emugameplay.tv@gmail.com")
	email.SetHeader("To", "carlosrgonzales88@gmail.com")
	email.SetHeader("Subject", "mi primer correo en go")
	htmlContent := `
	<html>
	<head>
		<style>
			body {
				font-family: Arial, sans-serif;
			}
			h1 {
				color: blue;
			}
		</style>
	</head>
	<body>
		<h1>Hola Mundo!</h1>
		<p>Este es un correo electrónico con <strong>estilos</strong> CSS incrustados.</p>
	</body>
	</html>
`

	// Establecer el contenido del correo electrónico como HTML
	email.SetBody("text/html", htmlContent)

	// Configuración del servidor SMTP
	smtpServer := mail.NewDialer("smtp.gmail.com", 587, "emugameplay.tv@gmail.com", "jwve tmdz qkgn krxa")

	// Intentar enviar el correo electrónico
	err := smtpServer.DialAndSend(email)
	if err != nil {
		fmt.Println("Error al enviar el correo electrónico:", err)
		return
	}

	fmt.Println("Correo electrónico enviado correctamente")
}
