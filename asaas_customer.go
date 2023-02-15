package asaas

import (
	"encoding/json"
	"fmt"
)

type Customer struct {
	Object               string `json:"object,omitempty"`
	ID                   string `json:"id"`
	DateCreated          string `json:"dateCreated,omitempty"`
	Name                 string `json:"name"`
	Email                string `json:"email"`
	Phone                string `json:"phone"`
	MobilePhone          string `json:"mobilePhone"`
	Address              string `json:"address"`
	AddressNumber        string `json:"addressNumber"`
	Complement           string `json:"complement"`
	Province             string `json:"province"`
	PostalCode           string `json:"postalCode"`
	CpfCnpj              string `json:"cpfCnpj"`
	PersonType           string `json:"personType,omitempty"`
	Deleted              bool   `json:"deleted"`
	AdditionalEmails     string `json:"additionalEmails"`
	ExternalReference    string `json:"externalReference"`
	NotificationDisabled bool   `json:"notificationDisabled"`
	City                 int    `json:"city"`
	State                string `json:"state"`
	Country              string `json:"country,omitempty"`
}

func (asaas *AsaasClient) NewCustomer(mode string, req Customer) (*Customer, *Error, error) {
	data, _ := json.Marshal(req)
	var response *Customer
	err, errAPI := asaas.Request(mode, "POST", fmt.Sprintf("customers"), data, &response)
	if err != nil {
		return nil, nil, err
	}
	if errAPI != nil {
		return nil, errAPI, nil
	}
	return response, nil, nil
}
