package data

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/mangatmodi/k8s-loadtest/slave/util"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

/**
* The package abstracts the calls to Master Data
**/

type OfferPublisherLink struct {
	Id          string `json:"id"`
	OfferId     string `json:"offer_id"`
	PublisherId string `json:"publisher_id"`
}

type ID struct {
	Value string `json:"id"`
}

func putEntity(url string, d []byte) (*ID, error) {
	client := &http.Client{}
	client.Timeout = time.Second * 15

	body := bytes.NewBuffer(d)
	req, err := http.NewRequest(http.MethodPut, url, body)
	if err != nil {
		log.Fatalf("http.NewRequest() failed with '%s'\n", err)
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("client.Do() failed with '%s'\n", err)
		return nil, err
	}
	s, err := strconv.Atoi(resp.Status)
	if s > 399 {
		return nil, errors.New("Unable to create entity because: " + resp.Status)
	}

	defer resp.Body.Close()
	d, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var id ID
	err = json.Unmarshal(d, &id)
	if err != nil {
		return nil, err
	}
	return &id, err
}

func putPublisher() (*ID, error) {
	p, err := RandomPublisher()
	if err != nil {
		log.Fatalf("create random publisher failed with '%s'\n", err)
		return nil, err
	}
	d, err := json.Marshal(*p)
	if err != nil {
		log.Fatalf("json.Marshal() failed with '%s'\n", err)
		return nil, err
	}
	masterDataURL, err := util.GetEnv("MASTER_DATA_URL")
	pubRoute := masterDataURL + "/api/v1/publishers/"
	if err != nil {
		log.Fatalf("getting MASTER_DATA_URL  failed with '%s'\n", err)
		return nil, err
	}
	return putEntity(pubRoute, d)
}

func putAdvertiser() (*ID, error) {
	a, err := RandomAdvertiser()
	if err != nil {
		log.Fatalf("create random advertiser failed with '%s'\n", err)
		return nil, err
	}
	d, err := json.Marshal(*a)
	if err != nil {
		log.Fatalf("json.Marshal() failed with '%s'\n", err)
		return nil, err
	}
	masterDataURL, err := util.GetEnv("MASTER_DATA_URL")
	advRoute := masterDataURL + "/api/v1/advertisers/"
	if err != nil {
		log.Fatalf("getting MASTER_DATA_URL  failed with '%s'\n", err)
		return nil, err
	}
	return putEntity(advRoute, d)
}

func putOffer(advId *ID) (*ID, error) {
	o, err := RandomOffer()
	o.AdvertiserID = advId.Value
	if err != nil {
		log.Fatalf("create random offer failed with '%s'\n", err)
		return nil, err
	}
	d, err := json.Marshal(*o)
	if err != nil {
		log.Fatalf("json.Marshal() failed with '%s'\n", err)
		return nil, err
	}
	masterDataURL, err := util.GetEnv("MASTER_DATA_URL")
	offerRoute := masterDataURL + "/api/v1/offers/"
	if err != nil {
		log.Fatalf("getting MASTER_DATA_URL failed with '%s'\n", err)
		return nil, err
	}
	return putEntity(offerRoute, d)
}

func putOfferPublisherLink(pubId, offerId *ID) (*ID, error) {
	op, err := OfferPublisherInputLink(pubId, offerId)
	if err != nil {
		log.Fatalf("create random offer failed with '%s'\n", err)
		return nil, err
	}
	d, err := json.Marshal(*op)
	if err != nil {
		log.Fatalf("json.Marshal() failed with '%s'\n", err)
		return nil, err
	}
	masterDataURL, err := util.GetEnv("MASTER_DATA_URL")
	opRoute := masterDataURL + "/api/v1/offer-publisher/"
	if err != nil {
		log.Fatalf("getting MASTER_DATA_URL failed with '%s'\n", err)
		return nil, err
	}
	return putEntity(opRoute, d)
}

func GetOfferPublisherLink() ([]*OfferPublisherLink, error) {
	adv, err := putAdvertiser()
	if err != nil {
		log.Println("Error in creating advertiser: %+v\n", err)
		return []*OfferPublisherLink{}, err
	}

	offerId, err := putOffer(adv)
	if err != nil {
		log.Println("Error in creating offer: %+v\n", err)
		return []*OfferPublisherLink{}, err
	}

	pubId, err := putPublisher()
	if err != nil {
		log.Println("Error in creating publisher: %+v\n", err)
		return []*OfferPublisherLink{}, err
	}

	linkId, err := putOfferPublisherLink(pubId, offerId)
	if err != nil {
		log.Println("Error in creating offer-publisher link: %+v\n", err)
		return []*OfferPublisherLink{}, err
	}

	op := &OfferPublisherLink{linkId.Value, offerId.Value, pubId.Value}
	links := []*OfferPublisherLink{op}

	return links, nil
}
