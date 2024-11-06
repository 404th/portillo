package email

import (
	"log"

	"github.com/404th/portillo/config/data"
	"github.com/mailjet/mailjet-apiv3-go/v4"
	"github.com/sirupsen/logrus"
)

func NewMailGen(pub, pri, toEmail, toName, fromEmail, fromName string) {
	mailjetClient := mailjet.NewMailjetClient(pub, pri)
	messagesInfo := []mailjet.InfoMessagesV31{
		{
			From: &mailjet.RecipientV31{
				Email: fromEmail,
				Name:  fromName,
			},
			To: &mailjet.RecipientsV31{
				mailjet.RecipientV31{
					Email: toEmail,
					Name:  toName,
				},
			},
			Subject:  data.RandomEmailSubject,
			TextPart: data.RandomEmailIntro,
			HTMLPart: data.RandomEmailBody,
		},
	}
	messages := mailjet.MessagesV31{Info: messagesInfo}
	res, err := mailjetClient.SendMailV31(&messages)
	if err != nil {
		log.Fatal(err)
	}
	logrus.Printf("Data: %+v\n", res.ResultsV31)
}
