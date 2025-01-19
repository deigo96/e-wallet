package external

import (
	"github.com/deigo96/e-wallet.git/config"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
)

type Midtrans struct {
	ServerKey  string
	ClientKey  string
	MerchantID string
	BaseURL    string
	APIVersion string
	Client     coreapi.Client
}

func NewMidtrans(config *config.Configuration) *Midtrans {
	c := coreapi.Client{}
	c.New(config.Midtrans.ServerKey, midtrans.Sandbox)

	return &Midtrans{
		ServerKey:  config.Midtrans.ServerKey,
		ClientKey:  config.Midtrans.ClientKey,
		MerchantID: config.Midtrans.MerchantID,
		BaseURL:    config.Midtrans.BaseURL,
		APIVersion: config.Midtrans.APIVersion,
		Client:     c,
	}
}

func (m *Midtrans) Url(endpoint string) string {
	return m.BaseURL + m.APIVersion + "/" + endpoint
}

// func (m *Midtrans) NewClient() coreapi.Client {
// 	c := coreapi.Client{}

// 	c.New(m.ServerKey, midtrans.Sandbox)

// 	return c
// }

// func (m *Midtrans) Client() *coreapi.Client {
// 	return &coreapi.Client{
// 		ServerKey: m.ServerKey,
// 		ClientKey: m.ClientKey,
// 		Env:       midtrans.Sandbox,
// 	}
// }

// func (m *Midtrans) Call(endpoint string, payload []byte) (any, error) {
// 	m.Client().ChargeTransaction(&coreapi.ChargeReq{})
// }
