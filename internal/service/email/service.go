package sendEmail

import (
	"io"
	"strconv"

	"github.com/itelman/doodocs-rest/config"
	"gopkg.in/gomail.v2"
)

type SendEmailService struct {
	config config.Config
}

func NewSendEmailService(config config.Config) *SendEmailService {
	return &SendEmailService{
		config: config,
	}
}

func (s *SendEmailService) SendFileToEmail(emails []string, file io.Reader, fileName string) error {
	for _, email := range emails {
		msg := gomail.NewMessage()
		msg.SetHeader("From", s.config.SMTPUsername)
		msg.SetHeader("To", email)
		msg.SetHeader("Subject", "File from Golang")
		msg.SetBody("text/plain", "This is the file you requested.")

		msg.Attach(fileName, gomail.SetCopyFunc(func(w io.Writer) error {
			_, err := io.Copy(w, file)
			return err
		}))

		port, err := strconv.Atoi(s.config.SMTPPort)
		if err != nil {
			return err
		}

		d := gomail.NewDialer(s.config.SMTPServer, port, s.config.SMTPUsername, s.config.SMTPPassword)
		if err := d.DialAndSend(msg); err != nil {
			return err
		}
	}
	return nil
}
