package asaas_test

import (
	"os"
	"testing"

	"github.com/rafaeltokyo/asaas-go-sdk"
)

func TestNewCustomer(t *testing.T) {
	client := asaas.NewAsaasClient(os.Getenv("ASAAS_TOKEN"))

	customerData := asaas.Customer{
		Name:                 "Marcelo Almeida",
		Email:                "marcelo.almeida@gmail.com",
		Phone:                "4738010919",
		MobilePhone:          "4799376637",
		CpfCnpj:              "24971563792",
		PostalCode:           "01310-000",
		Address:              "Av. Paulista",
		AddressNumber:        "150",
		Complement:           "Sala 201",
		Province:             "Centro",
		ExternalReference:    "12987382",
		NotificationDisabled: false,
		AdditionalEmails:     "marcelo.almeida2@gmail.com,marcelo.almeida3@gmail.com",
	}

	customerResponse, errAPI, err := client.NewCustomer(customerData)
	if err != nil {
		t.Errorf("err : %s", err)
		return
	}

	if errAPI != nil {
		t.Errorf("errAPI : %#v", errAPI)
		return
	}
	if customerResponse == nil {
		t.Error("customerResponse is null")
		return
	}

}
package asaas_test

import (
	"os"
	"testing"

	"github.com/rafaeltokyo/asaas-go-sdk"
)

func TestNewCustomer(t *testing.T) {
	client := asaas.NewAsaasClient(os.Getenv("ASAAS_TOKEN"))

	customerData := asaas.Customer{
		Name:                 "Marcelo Almeida",
		Email:                "marcelo.almeida@gmail.com",
		Phone:                "4738010919",
		MobilePhone:          "4799376637",
		CpfCnpj:              "24971563792",
		PostalCode:           "01310-000",
		Address:              "Av. Paulista",
		AddressNumber:        "150",
		Complement:           "Sala 201",
		Province:             "Centro",
		ExternalReference:    "12987382",
		NotificationDisabled: false,
		AdditionalEmails:     "marcelo.almeida2@gmail.com,marcelo.almeida3@gmail.com",
	}

	customerResponse, errAPI, err := client.NewCustomer(customerData)
	if err != nil {
		t.Errorf("err : %s", err)
		return
	}

	if errAPI != nil {
		t.Errorf("errAPI : %#v", errAPI)
		return
	}
	if customerResponse == nil {
		t.Error("customerResponse is null")
		return
	}

}
