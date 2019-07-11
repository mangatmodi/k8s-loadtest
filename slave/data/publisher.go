package data

import (
	"github.com/google/uuid"
)

type Publisher struct {
	DisplayName     string `json:"display_name"`
	PostbackUrl     string `json:"postback_url"`
	IntegrationType string `json:"integration_type"`
	AccountType     string `json:"account_type"`
	Status          string `json:"status"`
}

func RandomPublisher() (*Publisher, error) {
	var err error
	p := &Publisher{}
	r, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	p.DisplayName = r.String()
	p.PostbackUrl = "https://publisher.nieeoed.com"
	p.IntegrationType = "API"
	p.AccountType = "DIRECT"
	p.Status = "ACTIVE"
	return p, nil
}
