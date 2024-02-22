package ports

import (
	"github.com/ichungelo/assestment_go_source_code_krisna_satriadi/core/model"
	utilerrors "github.com/ichungelo/assestment_go_source_code_krisna_satriadi/utils/util_errors"
)

type RepositoryCustomer interface {
	CreateCustomer(customer *model.Customer) error
	GetListCustomer() ([]model.ResponseGetListCustomer, error)
	UpdateCustomerById(customer *model.Customer) error
	DeleteCustomerById(customerId *int) error
}

type ServiceCustomer interface {
	CreateCustomer(req *model.RequestCreateCustomer) *utilerrors.HttpError
	GetListCustomer() ([]model.ResponseGetListCustomer, *utilerrors.HttpError)
	UpdateCustomerById(req *model.RequestUpdateCustomerById) *utilerrors.HttpError
	DeleteCustomerById(req *model.RequestDeleteCustomerById) *utilerrors.HttpError
}
