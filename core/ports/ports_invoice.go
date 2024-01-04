package ports

import (
	"time"

	"github.com/ichungelo/assestment_go_source_code_krisna_satriadi/core/model"
	"github.com/ichungelo/assestment_go_source_code_krisna_satriadi/utils"
)

type RepositoryInvoice interface {
	CreateInvoice(req *model.RequestCreateInvoice) error
	GetListInvoice(isDelete bool, limit int, offset int, issueDate *time.Time, subject *string, totalItems *int, customer *string, dueDate *time.Time, InvoiceId *int) (*model.ResponseGetListInvoice, error)
	GetInvoiceById(invoiceId *int) (*model.ResponseInvoiceById, error)
	UpdateInvoiceById(invoice *model.Invoice) error
	DeleteInvoiceById(invoiceId *int) error
}

type ServiceInvoice interface {
	CreateInvoice(req *model.RequestCreateInvoice) *utils.ErrorCode
	GetListInvoice(req *model.RequestGetListInvoice) (*model.ResponseGetListInvoice, *utils.ErrorCode)
	GetInvoiceById(req *model.RequestGetInvoiceById) (*model.ResponseInvoiceById, *utils.ErrorCode)
	UpdateInvoiceById(req *model.RequestUpdateInvoiceById) *utils.ErrorCode
	DeleteInvoiceById(req *model.RequestDeleteInvoiceById) *utils.ErrorCode
}
