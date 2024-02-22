package services

import (
	"github.com/ichungelo/assestment_go_source_code_krisna_satriadi/core/model"
	"github.com/ichungelo/assestment_go_source_code_krisna_satriadi/core/ports"
	utilerrors "github.com/ichungelo/assestment_go_source_code_krisna_satriadi/utils/util_errors"
	utillogger "github.com/ichungelo/assestment_go_source_code_krisna_satriadi/utils/util_logger"
)

type serviceQuantity struct {
	RepositoryQuantity ports.RepositoryQuantity
}

func NewServiceQuantity(rQuantity ports.RepositoryQuantity) *serviceQuantity {
	return &serviceQuantity{
		RepositoryQuantity: rQuantity,
	}
}

func (s *serviceQuantity) DeleteQuantityById(req *model.RequestDeleteQuantityById) *utilerrors.ErrorCode {
	err := s.RepositoryQuantity.DeleteQuantityById(&req.ItemId, &req.InvoiceId)
	if err != nil {
		utillogger.Error(err, nil)
		errData := utilerrors.ErrorCode{
			Code: utilerrors.ErrFailedDeleteQuantity,
			Err:  err,
		}
		return &errData
	}

	return nil
}
