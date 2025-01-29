package service

import (
	"encoding/json"
	"net/http"
)

type ExternalService struct {
	baseURL string
}

func NewExternalService(baseURL string) *ExternalService {
	return &ExternalService{baseURL: baseURL}
}

func (e *ExternalService) FetchData(endpoint string) (map[string]interface{}, error) {
	resp, err := http.Get(e.baseURL + endpoint)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	return data, nil
}
