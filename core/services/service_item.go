package services

import (
	"github.com/ichungelo/assestment_go_source_code_krisna_satriadi/core/model"
	"github.com/ichungelo/assestment_go_source_code_krisna_satriadi/core/ports"
	"github.com/ichungelo/assestment_go_source_code_krisna_satriadi/utils"
)

type serviceItem struct {
	RepositoryItem ports.RepositoryItem
}

func NewServiceItem(rItem ports.RepositoryItem) *serviceItem {
	return &serviceItem{
		RepositoryItem: rItem,
	}
}

func (s *serviceItem) CreateItem(req *model.RequestCreateItem) *utils.ErrorCode {
	var (
		item = model.Item{
			Name:       &req.Name,
			UnitPrice:  &req.UnitPrice,
			ItemTypeId: &req.ItemTypeId,
		}
	)

	err := s.RepositoryItem.CreateItem(&item)
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

func (s *serviceItem) GetListItem() ([]model.ResponseGetListItem, *utils.ErrorCode) {
	res, err := s.RepositoryItem.GetListItem()
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

func (s *serviceItem) UpdateItemById(req *model.RequestUpdateItemById) *utils.ErrorCode {
	var (
		item = model.Item{
			Id:         &req.ItemId,
			Name:       &req.Name,
			UnitPrice:  &req.UnitPrice,
			ItemTypeId: &req.ItemTypeId,
		}
	)

	err := s.RepositoryItem.UpdateItemById(&item)
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

func (s *serviceItem) DeleteItemById(req *model.RequestDeleteItemById) *utils.ErrorCode {
	err := s.RepositoryItem.DeleteItemById(&req.ItemId)
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
