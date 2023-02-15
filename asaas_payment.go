package asaas

import (
	"encoding/json"
	"fmt"
)

type Payment struct {
	Object                  string          `json:"object"`
	ID                      string          `json:"id"`
	DateCreated             string          `json:"dateCreated"`
	Customer                string          `json:"customer"`
	DueDate                 string          `json:"dueDate"`
	Value                   float32         `json:"value"`
	InstallmentCount        float32         `json:"installmentCount"`
	InstallmentValue        float32         `json:"installmentValue"`
	NetValue                float32         `json:"netValue"`
	BillingType             string          `json:"billingType"`
	Status                  string          `json:"status"`
	Description             string          `json:"description"`
	ExternalReference       string          `json:"externalReference"`
	OriginalDueDate         string          `json:"originalDueDate"`
	PaymentDate             string          `json:"paymentDate"`
	ClientPaymentDate       string          `json:"clientPaymentDate"`
	InvoiceURL              string          `json:"invoiceUrl"`
	BankSlipURL             string          `json:"bankSlipUrl"`
	InvoiceNumber           string          `json:"invoiceNumber"`
	Discount                PaymentDiscount `json:"discount"`
	Fine                    PaymentFine     `json:"fine"`
	Interest                PaymentInterest `json:"interest"`
	CreditCard              PaymentCreditCard `json:"creditCard"`
	CreditCardHolderInfo    PaymentCreditCardHolderInfo `json:"creditCardHolderInfo"`
	Deleted                 bool            `json:"deleted"`
	PostalService           bool            `json:"postalService"`
	Anticipated             bool            `json:"anticipated"`
}

type PaymentDiscount struct {
	Value            float32 `json:"value"`
	DueDateLimitDays int32   `json:"dueDateLimitDays"`
}
type PaymentFine struct {
	Value float32 `json:"value"`
}

type PaymentInterest struct {
	Value float32 `json:"value"`
}

type PaymentCreditCard struct {
	HolderName      string `json:"holderName"`
    Number          string `json:"number"`
    ExpiryMonth     string `json:"expiryMonth"`
    ExpiryYear      string `json:"expiryYear"`
    Ccv             string `json:"ccv"`
}

type PaymentCreditCardHolderInfo struct {
	Name                string `json:"name"`
    Email               string `json:"email"`
    CpfCnpj             string `json:"cpfCnpj"`
    PostalCode          string `json:"postalCode"`
    AddressNumber       string `json:"addressNumber"`
    AddressComplement   string `json:"addressComplement"`
    Phone               string `json:"phone"`
    MobilePhone         string `json:"mobilePhone"`
}

type PaymentBoleto struct {
	Customer          string  `json:"customer"`
	DueDate           string  `json:"dueDate"`
	Value             float32 `json:"value"`
	ExternalReference string  `json:"externalReference"`
	Description       string  `json:"description"`
}

type PaymentCard struct {
	Customer                string  `json:"customer"`
	DueDate                 string  `json:"dueDate"`
	Value                   float32 `json:"value"`
    InstallmentCount        float32 `json:"installmentCount"`
	InstallmentValue        float32 `json:"InstallmentValue"`
	ExternalReference       string  `json:"externalReference"`
	Description             string  `json:"description"`
    CreditCard              PaymentCreditCard `json:"creditCard"`
	CreditCardHolderInfo    PaymentCreditCardHolderInfo `json:"creditCardHolderInfo"`
}


type PaymentIdentificationField struct {
	IdentificationField     string `json:"identificationField"`
	NossoNumero             string `json:"nossoNumero"`
}

type PaymentDelete struct {
	Deleted     bool `json:"deleted"`
	ID          string `json:"id"`
}


func (asaas *AsaasClient) PaymentBoleto(mode string, req PaymentBoleto) (*Payment, *Error, error) {
	payment := Payment{
		Customer:          req.Customer,
		BillingType:       "BOLETO",
		DueDate:           req.DueDate,
		Value:             req.Value,
		Description:       req.Description,
		ExternalReference: req.ExternalReference,
		PostalService:     false,
	}
	data, _ := json.Marshal(payment)
	var response *Payment
	err, errAPI := asaas.Request(mode, "POST", fmt.Sprintf("payments"), data, &response)
	if err != nil {
		return nil, nil, err
	}
	if errAPI != nil {
		return nil, errAPI, nil
	}
	return response, nil, nil
}

func (asaas *AsaasClient) PaymentCard(mode string, req PaymentCard) (*Payment, *Error, error) {
	payment := Payment{
		Customer:          req.Customer,
		BillingType:       "CREDIT_CARD",
		DueDate:           req.DueDate,
		///Value:             req.Value,
		Description:       req.Description,
		ExternalReference: req.ExternalReference,
		PostalService:     false,
        CreditCard: PaymentCreditCard{
            HolderName:     req.CreditCard.HolderName,
            Number:         req.CreditCard.Number,
            ExpiryMonth:    req.CreditCard.ExpiryMonth,
            ExpiryYear:     req.CreditCard.ExpiryYear,
            Ccv:            req.CreditCard.Ccv,
        },
        CreditCardHolderInfo: PaymentCreditCardHolderInfo{
            Name:               req.CreditCardHolderInfo.Name,
            Email:              req.CreditCardHolderInfo.Email,
            CpfCnpj:            req.CreditCardHolderInfo.CpfCnpj,
            PostalCode:         req.CreditCardHolderInfo.PostalCode,
            AddressNumber:      req.CreditCardHolderInfo.AddressNumber,
            AddressComplement:  req.CreditCardHolderInfo.AddressComplement,
            Phone:              req.CreditCardHolderInfo.Phone,
            MobilePhone:        req.CreditCardHolderInfo.MobilePhone,
        },
	}
    
    if req.InstallmentCount > 1 {
        payment.InstallmentCount = req.InstallmentCount
        payment.InstallmentValue = req.InstallmentValue
    } else {
        payment.Value = req.Value
    }
    
	data, _ := json.Marshal(payment)
	var response *Payment
	err, errAPI := asaas.Request(mode, "POST", fmt.Sprintf("payments"), data, &response)
	if err != nil {
		return nil, nil, err
	}
	if errAPI != nil {
		return nil, errAPI, nil
	}
	return response, nil, nil
}

func (asaas *AsaasClient) GetPayment(mode, id string) (*Payment, *Error, error) {
	var response *Payment
	err, errAPI := asaas.Request(mode, "GET", fmt.Sprintf("payments/%s", id), nil, &response)
	if err != nil {
		return nil, nil, err
	}
	if errAPI != nil {
		return nil, errAPI, nil
	}
	return response, nil, nil
}

func (asaas *AsaasClient) GetIdentificationField(mode, id string) (*PaymentIdentificationField, *Error, error) {
	var response *PaymentIdentificationField
	err, errAPI := asaas.Request(mode, "GET", fmt.Sprintf("payments/%s/identificationField", id), nil, &response)
	if err != nil {
		return nil, nil, err
	}
	if errAPI != nil {
		return nil, errAPI, nil
	}
	return response, nil, nil
}

func (asaas *AsaasClient) DeletePayment(mode, id string) (*PaymentDelete, *Error, error) {
	var response *PaymentDelete
	err, errAPI := asaas.Request(mode, "DELETE", fmt.Sprintf("payments/%s", id), nil, &response)
	if err != nil {
		return nil, nil, err
	}
	if errAPI != nil {
		return nil, errAPI, nil
	}
	return response, nil, nil
}
