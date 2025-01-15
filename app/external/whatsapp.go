package external

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/deigo96/e-wallet.git/app/constant"
	"github.com/deigo96/e-wallet.git/config"
)

type Whatsapp struct {
	URL         string
	AccessToken string
	AccountID   string
	PhoneNumber string
	APIVersion  string
}

func NewWhatsappService(config *config.Configuration) *Whatsapp {
	return &Whatsapp{
		URL:         config.WAConfig.BaseURL,
		AccessToken: config.WAConfig.AccessToken,
		AccountID:   config.WAConfig.AccountID,
		PhoneNumber: config.WAConfig.PhoneNumber,
		APIVersion:  config.WAConfig.APIVersion,
	}
}

func (w *Whatsapp) SendMessage(to string, message string) (any, error) {

	request := fmt.Sprintf(string(constant.WhatsappOTPText), to, message)

	resp, err := w.sendMessage([]byte(request))
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (w *Whatsapp) sendMessage(payload []byte) (any, error) {
	url := w.URL + w.APIVersion + "/" + w.AccountID + "/messages"

	body := bytes.NewBuffer(payload)

	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+w.AccessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	log.Println("Response: " + string(bodyBytes))

	if resp.StatusCode != 200 && resp.StatusCode != 201 {
		log.Println("Error sending SMS message: " + resp.Status)
		return nil, constant.ErrInternalServerError
	}

	log.Println("Response: " + resp.Status)

	return resp, nil
}
