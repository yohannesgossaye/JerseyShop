package initiator

import (
	customerhandler "github.com/yohannesgossaye/internal/handlers/rest/http/customeraccount"
	customerservice "github.com/yohannesgossaye/internal/service/customeraccount"
	"github.com/yohannesgossaye/pkgs/logger"
)

func InitHandler(service *customerservice.CustomerService, log logger.Logger) *customerhandler.CustomerAccountHandler {
	return customerhandler.NewCustomerAccountHandler(*service, log)
}
