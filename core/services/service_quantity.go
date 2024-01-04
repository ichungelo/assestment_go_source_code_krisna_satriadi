package services

import (
	"github.com/ichungelo/assestment_go_source_code_krisna_satriadi/core/model"
	"github.com/ichungelo/assestment_go_source_code_krisna_satriadi/core/ports"
	"github.com/ichungelo/assestment_go_source_code_krisna_satriadi/utils"
)

type serviceQuantity struct {
	RepositoryQuantity ports.RepositoryQuantity
}

func NewServiceQuantity(rQuantity ports.RepositoryQuantity) *serviceQuantity {
	return &serviceQuantity{
		RepositoryQuantity: rQuantity,
	}
}

func (s *serviceQuantity) DeleteQuantityById(req *model.RequestDeleteQuantityById) *utils.ErrorCode {
	err := s.RepositoryQuantity.DeleteQuantityById(&req.ItemId, &req.InvoiceId)
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
