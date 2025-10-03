package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"trader-backend/models"
)

type FinamClient interface {
	GetPortfolio() (*models.Portfolio, error)
	GetMarketData(ticker string) (interface{}, error)
	CreateOrder(payload map[string]interface{}) (interface{}, error)
}

type finamClientImpl struct {
	BaseURL string
	Token   string
	Client  *http.Client
}

func NewFinamClient(token string) FinamClient {
	return &finamClientImpl{
		BaseURL: "https://demo-api.finam.ru", // демо URL
		Token:   token,
		Client:  &http.Client{},
	}
}

// Вспомогательная функция для создания запроса с токеном
func (f *finamClientImpl) newRequest(method, endpoint string, body interface{}) (*http.Request, error) {
	var buf bytes.Buffer
	if body != nil {
		if err := json.NewEncoder(&buf).Encode(body); err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, f.BaseURL+endpoint, &buf)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+f.Token)
	return req, nil
}

// Получение портфеля
func (f *finamClientImpl) GetPortfolio() (*models.Portfolio, error) {
	req, err := f.newRequest("GET", "/portfolio", nil)
	if err != nil {
		return nil, err
	}

	resp, err := f.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, errors.New("failed to get portfolio")
	}

	var pf models.Portfolio
	body, _ := ioutil.ReadAll(resp.Body)
	if err := json.Unmarshal(body, &pf); err != nil {
		return nil, err
	}

	return &pf, nil
}

// Получение рыночных данных
func (f *finamClientImpl) GetMarketData(ticker string) (interface{}, error) {
	req, err := f.newRequest("GET", "/market-data/"+ticker, nil)
	if err != nil {
		return nil, err
	}

	resp, err := f.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, errors.New("failed to get market data")
	}

	var data interface{}
	body, _ := ioutil.ReadAll(resp.Body)
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	return data, nil
}

// Создание ордера
func (f *finamClientImpl) CreateOrder(payload map[string]interface{}) (interface{}, error) {
	req, err := f.newRequest("POST", "/order", payload)
	if err != nil {
		return nil, err
	}

	resp, err := f.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, errors.New("failed to create order")
	}

	var res interface{}
	body, _ := ioutil.ReadAll(resp.Body)
	if err := json.Unmarshal(body, &res); err != nil {
		return nil, err
	}

	return res, nil
}
