package services

import (
	"context"
	"fmt"
	"github.com/mailgun/mailgun-go/v4"
	"log"
	"time"
)

type Service struct{}

// SEND EMAIL METHOD THAT  WILL BE USED TO SEND EMAILS TO USERS
func (s *Service) SendMail(subject, body, recipient, Private, Domain string) bool {
	privateAPIKey := Private
	yourDomain := Domain

	mg := mailgun.NewMailgun(yourDomain, privateAPIKey)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	// Create a new template
	err := mg.CreateTemplate(ctx, &mailgun.Template{
		Name: "template!",
		Version: mailgun.TemplateVersion{
			Template: `'<div class="entry"> <h1>{{.title}}</h1> <div class="body"> {{.body}} </div> </div>'`,
			Engine:   mailgun.TemplateEngineGo,
			Tag:      "v1",
		},
	})
	if err != nil {
		return false
	}

	// Create a new message with template
	m := mg.NewMessage("Oja Ecommerce <Oja@Decadev.gon>", subject, "")
	m.SetTemplate("template!")

	// Add recipients
	_ = m.AddRecipient(recipient)

	// Add the variables recipient be used by the template
	_ = m.AddVariable("title", subject)
	_ = m.AddVariable("body", body)

	// Send the message with a 10 second timeout
	resp, id, err := mg.Send(ctx, m)
	if err != nil {
		log.Fatal(err)
		return false
	}

	fmt.Printf("ID: %s Resp: %s\n", id, resp)
	return true
}
