package ports

import (
	"time"

	"github.com/ichungelo/assestment_go_source_code_krisna_satriadi/core/model"
	utilerrors "github.com/ichungelo/assestment_go_source_code_krisna_satriadi/utils/util_errors"
)

type RepositoryInvoice interface {
	CreateInvoice(req *model.RequestCreateInvoice) error
	GetListInvoice(isDelete bool, limit int, offset int, issueDate *time.Time, subject *string, totalItems *int, customer *string, dueDate *time.Time, InvoiceId *int) (*model.ResponseGetListInvoice, error)
	GetInvoiceById(invoiceId *int) (*model.ResponseInvoiceById, error)
	UpdateInvoiceById(req *model.RequestUpdateInvoiceById) error
	DeleteInvoiceById(invoiceId *int) error
}

type ServiceInvoice interface {
	CreateInvoice(req *model.RequestCreateInvoice) *utilerrors.ErrorCode
	GetListInvoice(req *model.RequestGetListInvoice) (*model.ResponseGetListInvoice, *utilerrors.ErrorCode)
	GetInvoiceById(req *model.RequestGetInvoiceById) (*model.ResponseInvoiceById, *utilerrors.ErrorCode)
	UpdateInvoiceById(req *model.RequestUpdateInvoiceById) *utilerrors.ErrorCode
	DeleteInvoiceById(req *model.RequestDeleteInvoiceById) *utilerrors.ErrorCode
}
