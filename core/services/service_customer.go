package services

import (
	"github.com/ichungelo/assestment_go_source_code_krisna_satriadi/core/model"
	"github.com/ichungelo/assestment_go_source_code_krisna_satriadi/core/ports"
	utilerrors "github.com/ichungelo/assestment_go_source_code_krisna_satriadi/utils/util_errors"
	utillogger "github.com/ichungelo/assestment_go_source_code_krisna_satriadi/utils/util_logger"
)

type serviceCustomer struct {
	RepositoryCustomer ports.RepositoryCustomer
}

func NewServiceCustomer(rCustomer ports.RepositoryCustomer) *serviceCustomer {
	return &serviceCustomer{
		RepositoryCustomer: rCustomer,
	}
}

func (s *serviceCustomer) CreateCustomer(req *model.RequestCreateCustomer) *utilerrors.HttpError {
	var (
		customer = model.Customer{
			Name:    &req.Name,
			Address: &req.Address,
		}
	)

	err := s.RepositoryCustomer.CreateCustomer(&customer)
	if err != nil {
		utillogger.Error(err, nil)
		httpError := utilerrors.NewHttpError(utilerrors.ErrFailedCreateCustomer, err)

		return httpError
	}

	return nil
}

func (s *serviceCustomer) GetListCustomer() ([]model.ResponseGetListCustomer, *utilerrors.HttpError) {
	res, err := s.RepositoryCustomer.GetListCustomer()
	if err != nil {
		utillogger.Error(err, nil)
		httpError := utilerrors.NewHttpError(utilerrors.ErrFailedGetCustomer, err)

		return nil, httpError
	}

	return res, nil
}

func (s *serviceCustomer) UpdateCustomerById(req *model.RequestUpdateCustomerById) *utilerrors.HttpError {
	var (
		customer = model.Customer{
			Id:      &req.CustomerId,
			Name:    &req.Name,
			Address: &req.Address,
		}
	)

	err := s.RepositoryCustomer.UpdateCustomerById(&customer)
	if err != nil {
		utillogger.Error(err, nil)
		httpError := utilerrors.NewHttpError(utilerrors.ErrFailedUpdateCustomer, err)

		return httpError
	}

	return nil
}

func (s *serviceCustomer) DeleteCustomerById(req *model.RequestDeleteCustomerById) *utilerrors.HttpError {
	err := s.RepositoryCustomer.DeleteCustomerById(&req.CustomerId)
	if err != nil {
		utillogger.Error(err, nil)
		httpError := utilerrors.NewHttpError(utilerrors.ErrFailedDeleteCustomer, err)

		return httpError
	}

	return nil
}
