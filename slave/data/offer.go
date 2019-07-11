package data

import (
	"github.com/erggo/datafiller"
	"github.com/google/uuid"
	"time"
)

type Offer struct {
	AdvertiserID       string    `json:"advertiser_id"`
	DisplayName        string    `json:"display_name"`
	State              string    `json:"state"`
	StartTime          time.Time `json:"start_time"`
	EndTime            time.Time `json:"end_time"`
	DeviceIDRequired   bool      `json:"device_id_required"`
	RevenueAmount      float64   `json:"revenue_amount"`
	RevenueCurrency    string    `json:"revenue_currency"`
	PayoutAmount       float64   `json:"payout_amount"`
	PayoutCurrency     string    `json:"payout_currency"`
	OfferURL           string    `json:"offer_url"`
	PreviewURL         string    `json:"preview_url"`
	DailyConversionCap int       `json:"daily_conversion_cap"`
	Category           string    `json:"category"`
	GeoTargetingRules  []struct {
		CountryCode string `json:"country_code"`
	} `json:"geo_targeting_rules"`
	DeviceTargetingRules []struct {
		OsName    string `json:"os_name"`
		OsVersion string `json:"os_version"`
	} `json:"device_targeting_rules"`
	AttributionWindow struct {
		Start    int `json:"start"`
		Duration int `json:"duration"`
	} `json:"attribution_window"`
}

func RandomOffer() (*Offer, error) {
	var err error
	o := &Offer{}
	datafiller.Fill(o)
	o, err = clean(o)
	return o, err
}

func clean(o *Offer) (*Offer, error) {
	var err error
	r, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	o.AdvertiserID = r.String()
	o.PreviewURL = "http://applift-test.com"
	o.OfferURL = "http://example.com?click_id={click_id}"
	o.PayoutCurrency = "USD"
	o.RevenueCurrency = "USD"
	o.State = "ACTIVE"
	o.Category = "UNKNOWN"
	o.DeviceIDRequired = false
	o.EndTime = o.StartTime.Add(time.Hour * 24 * 30 * 12) //1 year from start
	o = cleanGeoRules(o)
	o = cleanDeviceRules(o)
	return o, nil
}

func cleanGeoRules(o *Offer) *Offer {
	geo := o.GeoTargetingRules[0:2]
	geo[0].CountryCode = "DE"
	geo[1].CountryCode = "US"
	o.GeoTargetingRules = geo
	return o
}

func cleanDeviceRules(o *Offer) *Offer {
	dev := o.DeviceTargetingRules[0:1]
	dev[0].OsName = "ANDROID"
	dev[0].OsVersion = "5+"
	o.DeviceTargetingRules = dev
	return o
}
