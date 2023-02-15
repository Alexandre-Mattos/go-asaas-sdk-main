package asaas

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	///"os"
	"time"
)

type AsaasClient struct {
	client *http.Client
	Token  string
}

type Error struct {
	Errors []ErrorItem `json:"errors"`
	Body   string      `json:"body"`
}

type ErrorItem struct {
	Code        string `json:"code"`
	Description string `json:"description"`
}

func NewAsaasClient(token string) *AsaasClient {
	return &AsaasClient{
		client: &http.Client{Timeout: 60 * time.Second},
		Token:  token,
	}
}

func (asaas *AsaasClient) Request(mode, method, action string, body []byte, out interface{}) (error, *Error) {
	if asaas.client == nil {
		asaas.client = &http.Client{Timeout: 60 * time.Second}
	}

	endpoint := fmt.Sprintf("%s/%s", asaas.devProd(mode), action)
	req, err := http.NewRequest(method, endpoint, bytes.NewBuffer(body))
	if err != nil {
		return err, nil
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("access_token", asaas.Token)
	res, err := asaas.client.Do(req)
	if err != nil {
		return err, nil
	}
	bodyResponse, err := ioutil.ReadAll(res.Body)
	if res.StatusCode > 201 {
		var errAPI Error
		err = json.Unmarshal(bodyResponse, &errAPI)
		if err != nil {
			return err, nil
		}
		errAPI.Body = string(bodyResponse)
		return nil, &errAPI
	}
	err = json.Unmarshal(bodyResponse, out)
	if err != nil {
		return err, nil
	}
	return nil, nil
}

func (asaas *AsaasClient) devProd(mode string) string {
	/**if os.Getenv("ENV") == "develop" {
		return "https://private-anon-ba248353a0-asaasv3.apiary-mock.com/api/v3"
	}*/
    
    if mode == "sandbox" {
        return "https://sandbox.asaas.com/api/v3"
    }
    
	return "https://www.asaas.com/api/v3"
}
