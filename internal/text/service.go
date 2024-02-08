package text

import (
	"fmt"
	"net/smtp"
)

var providerToSMSGatewayMapping = map[string]string{
	"att":          "txt.att.net",
	"boostmobile":  "smsmyboostmobile.com",
	"cricket":      "sms.cricketwireless.net",
	"sprint":       "messaging.sprintpcs.com",
	"tmobile":      "tmomail.net",
	"uscellular":   "email.uscc.net",
	"verison":      "vtext.com",
	"virginmobile": "vmobl.com",
}

type TextServicer interface {
	SendOverSMSGateway(toNumber string, msg string, provider string) error
}

type Text struct {
	Auth smtp.Auth
}

func (t Text) SendOverSMSGateway(toNumber string, msg string, provider string) error {
	gw, found := providerToSMSGatewayMapping[provider]
	if !found {
		return fmt.Errorf("no matching sms gateway for \"%s\"", provider)
	}

	err := smtp.SendMail("smtp.gmail.com:587",
		t.Auth,
		"", []string{fmt.Sprintf("%s@%s", toNumber, gw)}, []byte(msg))

	return err
}
