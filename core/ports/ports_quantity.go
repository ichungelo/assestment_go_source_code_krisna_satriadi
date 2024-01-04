package ports

import (
	"github.com/ichungelo/assestment_go_source_code_krisna_satriadi/core/model"
	"github.com/ichungelo/assestment_go_source_code_krisna_satriadi/utils"
)

type RepositoryQuantity interface {
	DeleteQuantityById(itemId *int, invoiceId *int) error
}

type ServiceQuantity interface {
	DeleteQuantityById(req *model.RequestDeleteQuantityById) *utils.ErrorCode
}