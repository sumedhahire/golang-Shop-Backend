package user

import (
	"bytes"
	"context"
	"fmt"
	"gopkg.in/gomail.v2"
	"html/template"
	"inventory/ent/entgen"
	"os"
)

type IUserService interface {
	Get(ctx context.Context, id string) (RSUser, error)
	List(ctx context.Context) ([]RSUser, error)
	Add(ctx context.Context, rq *RQUser) error
}

type SService struct {
	storage IUserStorage
}

func NewService(client *entgen.Client) IUserService {
	return SService{storage: NewStorage(client)}
}

func (s SService) Add(ctx context.Context, rq *RQUser) error {
	user := rq.MapTO()

	err := s.storage.Add(ctx, user)
	if err != nil {
		return err
	}

	err = sendUserMail(user)
	if err != nil {
		return err
	}

	return nil
}

func (s SService) Get(ctx context.Context, id string) (RSUser, error) {
	user, err := s.storage.Get(ctx, id)
	if err != nil {
		return RSUser{}, err
	}

	var rs RSUser
	rs.MapFrom(user)
	return rs, nil
}

func (s SService) List(ctx context.Context) ([]RSUser, error) {
	user, err := s.storage.List(ctx)
	if err != nil {
		return nil, err
	}

	rs := make([]RSUser, len(user))
	for index, value := range user {
		rs[index].MapFrom(&value)
	}
	return rs, nil
}

func sendUserMail(user *User) error {
	appMail := os.Getenv("APP_MAIL")
	appMailPassword := os.Getenv("APP_MAIL_PASSWORD")

	tmpl, err := template.ParseFiles("accepted.gohtml")
	if err != nil {
		return err
	}

	// Data to fill in the template
	data := struct {
		UserName string
		Password string
	}{
		UserName: user.Email,
		Password: user.Password,
	}

	// Render template to HTML string
	var body bytes.Buffer
	if err := tmpl.Execute(&body, data); err != nil {
		return err
	}

	// Set up email message
	m := gomail.NewMessage()

	// Sender info
	m.SetHeader("From", appMail)

	// Recipient(s)
	m.SetHeader("To", user.Email)

	// Subject
	sub := fmt.Sprintf("Login Details for %s %s", user.FirstName, user.LastName)
	m.SetHeader("Subject", sub)

	// Set HTML body
	m.SetBody("text/html", body.String())

	// Attach PDF file (change path as needed)
	//m.Attach("Bifurcation.pdf")

	// SMTP dialer setup (change host, port, user, pass)
	d := gomail.NewDialer("smtp.gmail.com", 587, appMail, appMailPassword)

	// Send the email
	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}
