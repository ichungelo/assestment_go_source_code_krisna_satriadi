package services

import (
	"github.com/ichungelo/assestment_go_source_code_krisna_satriadi/core/model"
	"github.com/ichungelo/assestment_go_source_code_krisna_satriadi/core/ports"
	"github.com/ichungelo/assestment_go_source_code_krisna_satriadi/utils"
)

type serviceCustomer struct {
	RepositoryCustomer ports.RepositoryCustomer
}

func NewServiceCustomer(rCustomer ports.RepositoryCustomer) *serviceCustomer {
	return &serviceCustomer{
		RepositoryCustomer: rCustomer,
	}
}

func (s *serviceCustomer) CreateCustomer(req *model.RequestCreateCustomer) *utils.ErrorCode {
	var (
		customer = model.Customer{
			Name:    &req.Name,
			Address: &req.Address,
		}
	)

	err := s.RepositoryCustomer.CreateCustomer(&customer)
	if err != nil {
		utils.Error(err, nil)
		errData := utils.ErrorCode{
			Code: utils.ERR_FAILED_CREATE_CUSTOMER,
			Err:  err,
		}
		return &errData

	}

	return nil
}

func (s *serviceCustomer) GetListCustomer() ([]model.ResponseGetListCustomer, *utils.ErrorCode) {
	res, err := s.RepositoryCustomer.GetListCustomer()
	if err != nil {
		utils.Error(err, nil)
		errData := utils.ErrorCode{
			Code: utils.ERR_FAILED_GET_CUSTOMER,
			Err:  err,
		}
		return nil, &errData
	}

	return res, nil
}

func (s *serviceCustomer) UpdateCustomerById(req *model.RequestUpdateCustomerById) *utils.ErrorCode {
	var (
		customer = model.Customer{
			Id:      &req.CustomerId,
			Name:    &req.Name,
			Address: &req.Address,
		}
	)

	err := s.RepositoryCustomer.UpdateCustomerById(&customer)
	if err != nil {
		utils.Error(err, nil)
		errData := utils.ErrorCode{
			Code: utils.ERR_FAILED_UPDATE_CUSTOMER,
			Err:  err,
		}
		return &errData

	}

	return nil
}

func (s *serviceCustomer) DeleteCustomerById(req *model.RequestDeleteCustomerById) *utils.ErrorCode {
	err := s.RepositoryCustomer.DeleteCustomerById(&req.CustomerId)
	if err != nil {
		utils.Error(err, nil)
		errData := utils.ErrorCode{
			Code: utils.ERR_FAILED_DELETE_CUSTOMER,
			Err:  err,
		}
		return &errData

	}

	return nil
}
