package ports

import (
	"time"

	"github.com/ichungelo/assestment_go_source_code_krisna_satriadi/core/model"
	"github.com/ichungelo/assestment_go_source_code_krisna_satriadi/utils"
)

type RepositoryInvoice interface {
	CreateInvoice(invoice *model.Invoice) error
	GetListInvoice(isDelete bool, limit int, offset int, issueDate *time.Time, subject *string, totalItems *int, customer *string,  dueDate *time.Time, InvoiceId *int) ([]model.Invoice, error)
	UpdateInvoiceById(invoice *model.Invoice) error
	DeleteInvoiceById(invoiceId *int) error
}

type ServiceInvoice interface {
	CreateInvoice(req *model.RequestCreateInvoice) *utils.ErrorCode
	GetListInvoice(req *model.RequestGetListInvoice) ([]model.Invoice, *utils.ErrorCode)
	UpdateInvoiceById(req *model.RequestUpdateInvoiceById) *utils.ErrorCode
	DeleteInvoiceById(req *model.RequestDeleteInvoiceById) *utils.ErrorCode
}