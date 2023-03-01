package asaas

import (
	"encoding/json"
	"fmt"
)

type Webhook struct {
	URL         string `json:"url"`
	Email       string `json:"email"`
	Enabled     bool   `json:"enabled"`
	Interrupted bool   `json:"interrupted"`
	APIVersion  int64  `json:"apiVersion"`
	AuthToken   string `json:"authToken"`
}

func (asaas *AsaasClient) GetAll(mode string) (*Webhook, *Error, error) {
	var response *Webhook

	err, errAPI := asaas.Request(mode, "GET", fmt.Sprintf("webhook"), nil, &response)

	if err != nil {
		return nil, nil, err
	}

	if errAPI != nil {
		return nil, errAPI, nil
	}

	return response, nil, nil
}

func (asaas *AsaasClient) Post(mode string, req Webhook) (*Webhook, *Error, error) {
	webhook := Webhook{
		URL:         req.URL,
		Email:       req.Email,
		Interrupted: req.Interrupted,
		APIVersion:  req.APIVersion,
		AuthToken:   req.AuthToken,
	}

	data, _ := json.Marshal(webhook)

	var response *Webhook

	err, errAPI := asaas.Request(mode, "POST", fmt.Sprintf("webhook"), data, &response)

	if err != nil {
		return nil, nil, err
	}

	if errAPI != nil {
		return nil, errAPI, nil
	}

	return response, nil, nil
}
