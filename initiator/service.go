package initiator

import (
	"github.com/yohannesgossaye/internal/persistence/postgres/customeraccount"
	customerservice "github.com/yohannesgossaye/internal/service/customeraccount"
	"github.com/yohannesgossaye/pkgs/utils/email"
)

func InitService(repo *postgres.CustomerRepositary, sender email.Sender) *customerservice.CustomerService {
	return customerservice.NewCustomerService(repo, sender)
}
