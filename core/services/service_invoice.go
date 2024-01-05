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
	status, err := utils.InvoiceStatusValidator(req.Status)
	if err != nil {
		utils.Error(err, nil)
		errData := utils.ErrorCode{
			Code: utils.ERR_VALIDATE,
			Err:  err,
		}
		return &errData
	}

	req.Status = *status

	err = s.RepositoryInvoice.CreateInvoice(req)
	if err != nil {
		utils.Error(err, nil)
		errData := utils.ErrorCode{
			Code: utils.ERR_FAILED_CREATE_INVOICE,
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
		v, err := time.Parse("2006-01-02", *req.IssueDate)

		if err != nil {
			utils.Error(err, nil)
			errData := utils.ErrorCode{
				Code: utils.ERR_PARSE_DATE,
				Err:  err,
			}
			return nil, &errData
		}
		issueDate = &v
	}

	if req.DueDate != nil {
		v, err := time.Parse("2006-01-02", *req.DueDate)

		if err != nil {
			utils.Error(err, nil)
			errData := utils.ErrorCode{
				Code: utils.ERR_PARSE_DATE,
				Err:  err,
			}
			return nil, &errData

		}
		dueDate = &v
	}

	res, err := s.RepositoryInvoice.GetListInvoice(isDelete, limit, offset, issueDate, req.Subject, totalItems, req.Customer, dueDate, invoiceId)
	if err != nil {
		utils.Error(err, nil)
		errData := utils.ErrorCode{
			Code: utils.ERR_FAILED_GET_INVOICE,
			Err:  err,
		}
		return nil, &errData
	}

	return res, nil
}

func (s *serviceInvoice) GetInvoiceById(req *model.RequestGetInvoiceById) (*model.ResponseInvoiceById, *utils.ErrorCode) {
	var (
		subTotal   float64
		tax        float64
		grandTotal float64
	)
	res, err := s.RepositoryInvoice.GetInvoiceById(&req.InvoiceId)
	if err != nil {
		utils.Error(err, nil)
		errData := utils.ErrorCode{
			Code: utils.ERR_FAILED_GET_INVOICE,
			Err:  err,
		}
		return nil, &errData

	}

	for _, v := range res.Items {
		var (
			quantity  = *v.Quantity
			unitPrice = *v.UnitPrice
		)
		subTotal += float64(quantity * unitPrice)
	}

	tax = float64(subTotal) * 0.1
	grandTotal = subTotal + tax

	res.SubTotal = &subTotal
	res.Tax = &tax
	res.GrandTotal = &grandTotal

	return res, nil
}

func (s *serviceInvoice) UpdateInvoiceById(req *model.RequestUpdateInvoiceById) *utils.ErrorCode {
	status, err := utils.InvoiceStatusValidator(req.Status)
	if err != nil {
		utils.Error(err, nil)
		errData := utils.ErrorCode{
			Code: utils.ERR_VALIDATE,
			Err:  err,
		}
		return &errData
	}

	req.Status = *status

	err = s.RepositoryInvoice.UpdateInvoiceById(req)
	if err != nil {
		utils.Error(err, nil)
		errData := utils.ErrorCode{
			Code: utils.ERR_FAILED_UPDATE_INVOICE,
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
			Code: utils.ERR_FAILED_DELETE_INVOICE,
			Err:  err,
		}
		return &errData

	}

	return nil
}
