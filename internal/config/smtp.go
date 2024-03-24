package config

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/smtp"
	"os"
)

type SmtpClient struct {
	Host     string
	Port     string
	Mail     string
	password string
	Addr     string
}

const (
	smtpHostName = "smtp_hostname"
	smtpPort     = "smtp_port"
	smtpPassword = "smtp_password"
	smtpMail     = "smtp_mail"
	pathTemplate = "internal/template/"
)

func NewSmtpClient() *SmtpClient {
	host := os.Getenv(smtpHostName)
	port := os.Getenv(smtpPort)
	addr := fmt.Sprintf("%s:%v", host, port)
	return &SmtpClient{
		Host:     host,
		Port:     port,
		Mail:     os.Getenv(smtpMail),
		password: os.Getenv(smtpPassword),
		Addr:     addr,
	}
}

func (s SmtpClient) auth() smtp.Auth {
	return smtp.PlainAuth("", s.Mail, s.password, s.Host)
}

func (s SmtpClient) SendMail(addressee string, body string) {
	mime := "MIME-version: 1.0;\nContent-Type:  text/html; charset=\"UTF-8\";\n\n"
	subject := "Subject: Reporte \n"
	msg := []byte(subject + mime + "\n" + body)
	addressees := []string{addressee}
	err := smtp.SendMail(s.Addr, s.auth(), s.Mail, addressees, msg)
	if err != nil {
		log.Fatal(err)
	}
}

func (s SmtpClient) ParseTemplate(templateFileName string, data interface{}) string {
	t, err := template.ParseFiles(pathTemplate + templateFileName)
	if err != nil {
		fmt.Printf("error read template html %s", err)
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		fmt.Printf("error parser template and data %s", err)
	}
	return buf.String()
}
