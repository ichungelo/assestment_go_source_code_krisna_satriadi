package fiberhandler

import "github.com/ichungelo/assestment_go_source_code_krisna_satriadi/core/ports"

type handlerCustomer struct {
	ports.ServiceCustomer
}

func NewCustomerHandler(sCustomer ports.ServiceCustomer) *handlerCustomer {
	return &handlerCustomer{
		ServiceCustomer: sCustomer,
	}
}