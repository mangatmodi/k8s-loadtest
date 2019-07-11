package data

type OfferPublisher struct {
	OfferID        string  `json:"offer_id"`
	PublisherID    string  `json:"publisher_id"`
	PayoutAmount   float64 `json:"payout_amount"`
	PayoutCurrency string  `json:"payout_currency"`
	Status         string  `json:"status"`
}

func OfferPublisherInputLink(pubId, offerId *ID) (*OfferPublisher, error) {
	op := &OfferPublisher{}
	op.OfferID = offerId.Value
	op.PublisherID = pubId.Value
	op.PayoutCurrency = "USD"
	op.PayoutAmount = 0.9
	op.Status = "ACTIVE"
	return op, nil
}
