package ports

import (
	"github.com/ichungelo/assestment_go_source_code_krisna_satriadi/core/model"
	"github.com/ichungelo/assestment_go_source_code_krisna_satriadi/utils"
)

type RepositoryCustomer interface {
	CreateCustomer(customer *model.Customer) error
	GetListCustomer() ([]model.ResponseGetListCustomer, error)
	UpdateCustomerById(customer *model.Customer) error
	DeleteCustomerById(customerId *int) error
}

type ServiceCustomer interface {
	CreateCustomer(req *model.RequestCreateCustomer) *utils.ErrorCode
	GetListCustomer() ([]model.ResponseGetListCustomer, *utils.ErrorCode)
	UpdateCustomerById(req *model.RequestUpdateCustomerById) *utils.ErrorCode
	DeleteCustomerById(req *model.RequestDeleteCustomerById) *utils.ErrorCode
}
