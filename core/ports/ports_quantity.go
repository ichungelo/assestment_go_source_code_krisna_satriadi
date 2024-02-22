package ports

import (
	"github.com/ichungelo/assestment_go_source_code_krisna_satriadi/core/model"
	utilerrors "github.com/ichungelo/assestment_go_source_code_krisna_satriadi/utils/util_errors"
)

type RepositoryQuantity interface {
	DeleteQuantityById(itemId *int, invoiceId *int) error
}

type ServiceQuantity interface {
	DeleteQuantityById(req *model.RequestDeleteQuantityById) *utilerrors.HttpError
}
