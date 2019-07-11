package data

import (
	"github.com/google/uuid"
)

type Advertiser struct {
	DisplayName string `json:"display_name"`
	Status      string `json:"status"`
}

func RandomAdvertiser() (*Advertiser, error) {
	var err error
	a := &Advertiser{}
	r, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	a.DisplayName = r.String()
	a.Status = "ACTIVE"
	return a, nil
}
