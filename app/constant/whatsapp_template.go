package constant

import "fmt"

var (
	WhatsappOTPTemplate = `{
		"messaging_product": "whatsapp",
		"recipient_type": "individual",
		"to": "%s",
		"type": "template",
		"template": {
			"name": "otp",
			"language": {
				"code": "en_US"
			},
			"components": [
				{
					"type": "body",
					"parameters": [
						{
							"type": "text",
							"text": "%s"
						}
					]
				}
			]
		}
	}`

	WhatsappOTPText = `{
		"messaging_product": "whatsapp",
		"recipient_type": "individual",
		"to": "%s",
		"type": "text",
		"text": {
			"body": "%s"
		}
	}`

	WhatsappOTPTemplateDefault = `{
		"messaging_product": "whatsapp",
		"to": "%s",
		"type": "template",
		"template": {
			"name": "hello_world",
			"language": {
				"code": "en_US"
			}
		}
	}`
)

func GetWhatsappOTPTemplate(phoneNumber string, otp string) []byte {
	return []byte(fmt.Sprintf(string(WhatsappOTPTemplate), phoneNumber, otp))
}
