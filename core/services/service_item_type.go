package services

import (
	"github.com/ichungelo/assestment_go_source_code_krisna_satriadi/core/model"
	"github.com/ichungelo/assestment_go_source_code_krisna_satriadi/core/ports"
	"github.com/ichungelo/assestment_go_source_code_krisna_satriadi/utils"
)

type serviceItemType struct {
	RepositoryItemType ports.RepositoryItemType
}

func NewServiceItemType(rItemType ports.RepositoryItemType) *serviceItemType {
	return &serviceItemType{
		RepositoryItemType: rItemType,
	}
}

func (s *serviceItemType) CreateItemType(req *model.RequestCreateItemType) *utils.ErrorCode {
	var (
		ItemType = model.ItemType{
			Name: &req.Name,
		}
	)

	err := s.RepositoryItemType.CreateItemType(&ItemType)
	if err != nil {
		utils.Error(err, nil)
		errData := utils.ErrorCode{
			Code: utils.ERR_FAILED_CREATE_ITEM_TYPE,
			Err:  err,
		}
		return &errData

	}

	return nil
}

func (s *serviceItemType) GetListItemType() ([]model.ResponseGetListItemType, *utils.ErrorCode) {
	res, err := s.RepositoryItemType.GetListItemType()
	if err != nil {
		utils.Error(err, nil)
		errData := utils.ErrorCode{
			Code: utils.ERR_FAILED_GET_ITEM_TYPE,
			Err:  err,
		}
		return nil, &errData
	}

	return res, nil
}

func (s *serviceItemType) UpdateItemTypeById(req *model.RequestUpdateItemTypeById) *utils.ErrorCode {
	var (
		itemType = model.ItemType{
			Id:   &req.ItemTypeId,
			Name: &req.Name,
		}
	)

	err := s.RepositoryItemType.UpdateItemTypeById(&itemType)
	if err != nil {
		utils.Error(err, nil)
		errData := utils.ErrorCode{
			Code: utils.ERR_FAILED_UPDATE_ITEM_TYPE,
			Err:  err,
		}
		return &errData

	}

	return nil
}

func (s *serviceItemType) DeleteItemTypeById(req *model.RequestDeleteItemTypeById) *utils.ErrorCode {
	err := s.RepositoryItemType.DeleteItemTypeById(&req.ItemTypeId)
	if err != nil {
		utils.Error(err, nil)
		errData := utils.ErrorCode{
			Code: utils.ERR_FAILED_DELETE_ITEM_TYPE,
			Err:  err,
		}
		return &errData

	}

	return nil
}
