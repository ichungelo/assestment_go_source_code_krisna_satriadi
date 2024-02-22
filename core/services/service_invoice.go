package services

import (
	"strconv"
	"time"

	"github.com/ichungelo/assestment_go_source_code_krisna_satriadi/core/model"
	"github.com/ichungelo/assestment_go_source_code_krisna_satriadi/core/ports"
	utilenum "github.com/ichungelo/assestment_go_source_code_krisna_satriadi/utils/util_enum"
	utilerrors "github.com/ichungelo/assestment_go_source_code_krisna_satriadi/utils/util_errors"
	utillogger "github.com/ichungelo/assestment_go_source_code_krisna_satriadi/utils/util_logger"
)

type serviceInvoice struct {
	RepositoryInvoice ports.RepositoryInvoice
}

func NewServiceInvoice(rInvoice ports.RepositoryInvoice) *serviceInvoice {
	return &serviceInvoice{
		RepositoryInvoice: rInvoice,
	}
}

func (s *serviceInvoice) CreateInvoice(req *model.RequestCreateInvoice) *utilerrors.HttpError {
	status, err := utilenum.EnumCheckInvoiceStatus(req.Status)
	if err != nil {
		utillogger.Error(err, nil)
		httpError := utilerrors.NewHttpError(utilerrors.ErrValidate, err)

		return httpError
	}

	req.Status = *status

	err = s.RepositoryInvoice.CreateInvoice(req)
	if err != nil {
		utillogger.Error(err, nil)
		httpError := utilerrors.NewHttpError(utilerrors.ErrFailedCreateInvoice, err)

		return httpError
	}

	return nil
}

func (s *serviceInvoice) GetListInvoice(req *model.RequestGetListInvoice) (listInvoice *model.ResponseGetListInvoice, apiErr *utilerrors.HttpError) {

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
			utillogger.Error(err, nil)
			httpError := utilerrors.NewHttpError(utilerrors.ErrParseDate, err)

			return nil, httpError
		}
		issueDate = &v
	}

	if req.DueDate != nil {
		v, err := time.Parse("2006-01-02", *req.DueDate)

		if err != nil {
			utillogger.Error(err, nil)
			httpError := utilerrors.NewHttpError(utilerrors.ErrParseDate, err)

			return nil, httpError
		}
		dueDate = &v
	}

	res, err := s.RepositoryInvoice.GetListInvoice(isDelete, limit, offset, issueDate, req.Subject, totalItems, req.Customer, dueDate, invoiceId)
	if err != nil {
		utillogger.Error(err, nil)
		httpError := utilerrors.NewHttpError(utilerrors.ErrFailedGetInvoice, err)

		return nil, httpError
	}

	return res, nil
}

func (s *serviceInvoice) GetInvoiceById(req *model.RequestGetInvoiceById) (*model.ResponseInvoiceById, *utilerrors.HttpError) {
	var (
		subTotal   float64
		tax        float64
		grandTotal float64
	)
	res, err := s.RepositoryInvoice.GetInvoiceById(&req.InvoiceId)
	if err != nil {
		utillogger.Error(err, nil)
		httpError := utilerrors.NewHttpError(utilerrors.ErrFailedGetInvoice, err)

		return nil, httpError

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

func (s *serviceInvoice) UpdateInvoiceById(req *model.RequestUpdateInvoiceById) *utilerrors.HttpError {
	status, err := utilenum.EnumCheckInvoiceStatus(req.Status)
	if err != nil {
		utillogger.Error(err, nil)
		httpError := utilerrors.NewHttpError(utilerrors.ErrValidate, err)

		return httpError
	}

	req.Status = *status

	for i, v := range req.Items {
		action, err := utilenum.EnumCheckInvoiceItemAction(v.Action)
		if err != nil {
			utillogger.Error(err, nil)
			httpError := utilerrors.NewHttpError(utilerrors.ErrValidate, err)

			return httpError
		}

		req.Items[i].Action = *action
	}
	err = s.RepositoryInvoice.UpdateInvoiceById(req)
	if err != nil {
		utillogger.Error(err, nil)
		httpError := utilerrors.NewHttpError(utilerrors.ErrFailedUpdateInvoice, err)

		return httpError

	}

	return nil
}

func (s *serviceInvoice) DeleteInvoiceById(req *model.RequestDeleteInvoiceById) *utilerrors.HttpError {
	err := s.RepositoryInvoice.DeleteInvoiceById(&req.InvoiceId)
	if err != nil {
		utillogger.Error(err, nil)
		httpError := utilerrors.NewHttpError(utilerrors.ErrFailedDeleteInvoice, err)

		return httpError

	}

	return nil
}
