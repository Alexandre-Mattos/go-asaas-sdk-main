package asaas

import (
	"encoding/json"
	"fmt"
)

type TransferResponse struct {
	Object        string      `json:"object"`
	DateCreated   string      `json:"dateCreated"`
	EffectiveDate string      `json:"effectiveDate"`
	Status        string      `json:"status"`
	ID            string      `json:"id"`
	OperationType string      `json:"operationType"`
	Type          string      `json:"type"`
	Value         float32     `json:"value"`
	FailReason    string      `json:"failReason"`
	BankAccount   BankAccount `json:"bankAccount"`
}

type Transfer struct {
	Value         float64     `json:"value"`
	BankAccount   BankAccount `json:"bankAccount"`
	OperationType string      `json:"operationType"`
}

type BankAccount struct {
	Bank            Bank   `json:"bank"`
	AccountName     string `json:"accountName"`
	OwnerName       string `json:"ownerName"`
	cpfCnpj         string `json:"cpfCnpj"`
	agency          string `json:"agency"`
	account         string `json:"account"`
	accountDigit    string `json:"accountDigit"`
	bankAccountType string `json:"bankAccountType"`
}

type Bank struct {
	Code string `json:"code"`
}

type Pix struct {
	OperationType     string  `json:"operationType"`
	PixAddressKey     string  `json:"pixAddressKey"`
	PixAddressKeyType string  `json:"pixAddressKeyType"`
	Value             float32 `json:"value"`
}

func (asaas *AsaasClient) transfer(mode string, req Transfer) (*TransferResponse, *Error, error) {
	transfer := Transfer{
		OperationType: "TED",
		Value:         req.Value,
		BankAccount:   req.BankAccount,
	}
	data, _ := json.Marshal(transfer)
	var response *TransferResponse
	err, errAPI := asaas.Request(mode, "POST", fmt.Sprintf("transfers"), data, &response)
	if err != nil {
		return nil, nil, err
	}
	if errAPI != nil {
		return nil, errAPI, nil
	}

	return response, nil, nil
}

func (asaas *AsaasClient) transferPix(mode string, req Pix) (*TransferResponse, *Error, error) {
	pix := Pix{
		OperationType:     "PIX",
		PixAddressKey:     req.PixAddressKey,
		PixAddressKeyType: req.PixAddressKeyType,
		Value:             req.Value,
	}
	data, _ := json.Marshal(pix)
	var response *TransferResponse
	err, errAPI := asaas.Request(mode, "POST", fmt.Sprintf("transfers"), data, &response)
	if err != nil {
		return nil, nil, err
	}
	if errAPI != nil {
		return nil, errAPI, nil
	}

	return response, nil, nil
}
