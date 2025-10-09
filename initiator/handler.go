package initiator

import (
	customerhandler "github.com/yohannesgossaye/internal/handlers/rest/http/customeraccount"
	customerservice "github.com/yohannesgossaye/internal/service/customeraccount"
)

func InitHandler(service *customerservice.CustomerService) *customerhandler.CustomerAccountHandler {
	return customerhandler.NewCustomerAccountHandler(*service)
}
