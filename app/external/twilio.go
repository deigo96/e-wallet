package external

import (
	"encoding/json"
	"fmt"

	"github.com/deigo96/e-wallet.git/config"
	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
)

type TwilioService interface {
	SendOTP(phoneNumber, message string) error
}

type twilioService struct {
	config *config.Configuration
}

func NewTwilioService(config *config.Configuration) TwilioService {
	return &twilioService{
		config: config,
	}
}

func (t *twilioService) SendOTP(phoneNumber, message string) error {
	from := t.config.TwilioConfig.PhoneNumber
	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: t.config.TwilioConfig.AccountSID,
		Password: t.config.TwilioConfig.AuthToken,
	})

	params := &twilioApi.CreateMessageParams{}
	params.SetTo(phoneNumber)
	params.SetFrom(from)
	params.SetBody("Hello from Go!")

	resp, err := client.Api.CreateMessage(params)
	if err != nil {
		fmt.Println("Error sending SMS message: " + err.Error())
		return err
	} else {
		response, _ := json.Marshal(*resp)
		fmt.Println("Response: " + string(response))
	}

	return nil
}
