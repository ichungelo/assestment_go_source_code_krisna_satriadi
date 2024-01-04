package services

import (
	"strconv"
	"time"

	"github.com/ichungelo/assestment_go_source_code_krisna_satriadi/core/model"
	"github.com/ichungelo/assestment_go_source_code_krisna_satriadi/core/ports"
	"github.com/ichungelo/assestment_go_source_code_krisna_satriadi/utils"
)

type serviceInvoice struct {
	RepositoryInvoice ports.RepositoryInvoice
}

func NewServiceInvoice(rInvoice ports.RepositoryInvoice) *serviceInvoice {
	return &serviceInvoice{
		RepositoryInvoice: rInvoice,
	}
}

func (s *serviceInvoice) CreateInvoice(req *model.RequestCreateInvoice) *utils.ErrorCode {
	err := s.RepositoryInvoice.CreateInvoice(req)
	if err != nil {
		utils.Error(err, nil)
		errData := utils.ErrorCode{
			Code: "999",
			Err:  err,
		}
		return &errData

	}

	return nil
}

func (s *serviceInvoice) GetListInvoice(req *model.RequestGetListInvoice) (listInvoice *model.ResponseGetListInvoice, apiErr *utils.ErrorCode) {

	var (
		isDelete   bool
		limit      int = -1
		offset     int = -1
		totalItems *int
		invoiceId  *int
		issueDate  *time.Time
		dueDate    *time.Time
	)

	if req.IsDelete != nil {
		v, err := strconv.ParseBool(*req.IsDelete)
		if err == nil {
			isDelete = v
		}
	}

	if req.Limit != nil {
		v, err := strconv.Atoi(*req.Limit)

		if err == nil {
			limit = v
		}
	}

	if req.Offset != nil {
		v, err := strconv.Atoi(*req.Offset)
		if err == nil {
			offset = v
		}
	}

	if req.TotalItems != nil {
		v, err := strconv.Atoi(*req.TotalItems)
		if err == nil {
			totalItems = &v
		}
	}

	if req.InvoiceId != nil {
		v, err := strconv.Atoi(*req.InvoiceId)
		if err == nil {
			invoiceId = &v
		}
	}

	if req.IssueDate != nil {
		v, err := time.Parse(time.RFC3339, *req.IssueDate)
		if err == nil {
			issueDate = &v
		}
	}

	if req.DueDate != nil {
		v, err := time.Parse(time.RFC3339, *req.DueDate)
		if err == nil {
			dueDate = &v
		}
	}

	res, err := s.RepositoryInvoice.GetListInvoice(isDelete, limit, offset, issueDate, req.Subject, totalItems, req.Customer, dueDate, invoiceId)
	if err != nil {
		utils.Error(err, nil)
		errData := utils.ErrorCode{
			Code: "999",
			Err:  err,
		}
		return nil, &errData
	}

	return res, nil
}

func (s *serviceInvoice) GetInvoiceById(invoiceId *int) (*model.Invoice, *utils.ErrorCode) {
	//!TODO Add Service
	return nil, nil
}

func (s *serviceInvoice) UpdateInvoiceById(req *model.RequestUpdateInvoiceById) *utils.ErrorCode {
	var (
		Invoice = model.Invoice{
			Id:         &req.InvoiceId,
			Subject:    &req.Subject,
			CustomerId: &req.CustomerId,
		}
	)

	err := s.RepositoryInvoice.UpdateInvoiceById(&Invoice)
	if err != nil {
		utils.Error(err, nil)
		errData := utils.ErrorCode{
			Code: "999",
			Err:  err,
		}
		return &errData

	}

	return nil
}

func (s *serviceInvoice) DeleteInvoiceById(req *model.RequestDeleteInvoiceById) *utils.ErrorCode {
	err := s.RepositoryInvoice.DeleteInvoiceById(&req.InvoiceId)
	if err != nil {
		utils.Error(err, nil)
		errData := utils.ErrorCode{
			Code: "999",
			Err:  err,
		}
		return &errData

	}

	return nil
}
