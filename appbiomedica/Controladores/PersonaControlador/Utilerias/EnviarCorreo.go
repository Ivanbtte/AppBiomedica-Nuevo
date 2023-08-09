package Utilerias

import (
	"fmt"
	"log"
	"net/smtp"
)

func EnviarCorreo(mensaje string, correo string, asunto string) error {
	fmt.Println("enviando correo")
	from := "rogelio.hdzv2610@gmail.com"
	pass := "rogelioSCC261017"
	to := correo

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: " + asunto + "\n\n" +
		mensaje + "\n\n\n\n" +
		"Departamento de Biomédica"

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{to}, []byte(msg))

	if err != nil {
		log.Printf("smtp error: %s", err)
		return err
	}

	log.Print("sent, visit http://foobarbazz.mailinator.com")
	return nil
}
