package open_exchange_client

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"go.uber.org/multierr"
	"io/ioutil"
	"net/http"
)

type Currency struct {
	From string
	To   string
	Well float64
}

type Config struct {
	Url   string
	AppId string
}

type Client struct {
	config Config
}

type Response struct {
	Status  int
	Message string
	Data    []Currency
}

type oeAnswer struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Rates   map[string]float64
}

func New(config Config) *Client {
	return &Client{
		config: config,
	}
}

func (c *Client) url() string {
	return fmt.Sprint(c.config.Url + c.config.AppId)
}

func (c *Client) GetRates(currencies []Currency) (*Response, error) {
	resp, err := http.Get(c.url())
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer multierr.AppendInto(&err, resp.Body.Close())

	var oeAnswer oeAnswer
	err = json.Unmarshal(body, &oeAnswer)
	if err != nil {
		return nil, err
	}

	response := &Response{
		Status:  oeAnswer.Status,
		Message: oeAnswer.Message,
		Data:    make([]Currency, 0, len(currencies)),
	}

	for _, curr := range currencies {
		oeCurrFromUSDWell, ok := oeAnswer.Rates[curr.From]
		if !ok {
			logrus.Errorf("unfamiliar currency in DB: %s", curr.From)
			continue
		}
		oeCurrToUSDWell, ok := oeAnswer.Rates[curr.To]
		if !ok {
			logrus.Errorf("unfamiliar currency in DB: %s", curr.To)
			continue
		}
		response.Data = append(response.Data, Currency{
			From: curr.From,
			To:   curr.To,
			Well: oeCurrToUSDWell / oeCurrFromUSDWell,
		})
	}

	return response, nil
}
