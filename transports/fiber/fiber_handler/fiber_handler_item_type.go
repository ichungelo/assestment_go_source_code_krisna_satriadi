package fiberhandler

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/ichungelo/assestment_go_source_code_krisna_satriadi/core/model"
	"github.com/ichungelo/assestment_go_source_code_krisna_satriadi/core/ports"
	fiberpresenter "github.com/ichungelo/assestment_go_source_code_krisna_satriadi/transports/fiber/fiber_presenter"
	"github.com/ichungelo/assestment_go_source_code_krisna_satriadi/utils"
)

type RouterFiberItemType interface {
	CreateItemType() fiber.Handler
	GetListItemType() fiber.Handler
	UpdateItemTypeById() fiber.Handler
	DeleteItemTypeById() fiber.Handler
}

type handlerItemType struct {
	ports.ServiceItemType
}

func NewItemTypeHandler(sItemType ports.ServiceItemType) *handlerItemType {
	return &handlerItemType{
		ServiceItemType: sItemType,
	}
}

func (h *handlerItemType) CreateItemType() fiber.Handler {
	return func(c *fiber.Ctx) error {
		req := model.RequestCreateItemType{}

		err := json.Unmarshal(c.Body(), &req)
		if err != nil {
			utils.Error(err, nil)
			errCode := utils.ErrorCode{
				Code: utils.ERR_FAILED_UNMARSHAL_JSON,
				Err:  err,
			}

			return fiberpresenter.Presenter(c, nil, nil, &errCode)
		}

		err = utils.Validate(req)
		if err != nil {
			utils.Error(err, nil)
			errCode := utils.ErrorCode{
				Code: utils.ERR_VALIDATE_STRUCT,
				Err:  err,
			}

			return fiberpresenter.Presenter(c, nil, nil, &errCode)
		}

		errCode := h.ServiceItemType.CreateItemType(&req)
		if errCode != nil {
			return fiberpresenter.Presenter(c, nil, nil, errCode)
		}

		return fiberpresenter.Presenter(c, nil, nil, nil)
	}
}

func (h *handlerItemType) GetListItemType() fiber.Handler {
	return func(c *fiber.Ctx) error {

		data, errCode := h.ServiceItemType.GetListItemType()
		if errCode != nil {
			return fiberpresenter.Presenter(c, nil, nil, errCode)
		}

		return fiberpresenter.Presenter(c, data, nil, nil)
	}
}

func (h *handlerItemType) UpdateItemTypeById() fiber.Handler {
	return func(c *fiber.Ctx) error {
		req := model.RequestUpdateItemTypeById{}

		err := json.Unmarshal(c.Body(), &req)
		if err != nil {
			utils.Error(err, nil)
			errCode := utils.ErrorCode{
				Code: utils.ERR_FAILED_UNMARSHAL_JSON,
				Err:  err,
			}

			return fiberpresenter.Presenter(c, nil, nil, &errCode)
		}

		itemTypeId, err := c.ParamsInt("itemTypeId", 0)
		if err != nil {
			utils.Error(err, nil)
			errCode := utils.ErrorCode{
				Code: utils.ERR_PARSE_DATA,
				Err:  err,
			}

			return fiberpresenter.Presenter(c, nil, nil, &errCode)
		}

		req.ItemTypeId = itemTypeId

		err = utils.Validate(req)
		if err != nil {
			utils.Error(err, nil)
			errCode := utils.ErrorCode{
				Code: utils.ERR_VALIDATE_STRUCT,
				Err:  err,
			}

			return fiberpresenter.Presenter(c, nil, nil, &errCode)
		}

		errCode := h.ServiceItemType.UpdateItemTypeById(&req)
		if errCode != nil {
			return fiberpresenter.Presenter(c, nil, nil, errCode)
		}

		return fiberpresenter.Presenter(c, nil, nil, nil)
	}
}

func (h *handlerItemType) DeleteItemTypeById() fiber.Handler {
	return func(c *fiber.Ctx) error {
		req := model.RequestDeleteItemTypeById{}

		itemTypeId, err := c.ParamsInt("itemTypeId", 0)
		if err != nil {
			utils.Error(err, nil)
			errCode := utils.ErrorCode{
				Code: utils.ERR_PARSE_DATA,
				Err:  err,
			}

			return fiberpresenter.Presenter(c, nil, nil, &errCode)
		}

		req.ItemTypeId = itemTypeId

		err = utils.Validate(req)
		if err != nil {
			utils.Error(err, nil)
			errCode := utils.ErrorCode{
				Code: utils.ERR_VALIDATE_STRUCT,
				Err:  err,
			}

			return fiberpresenter.Presenter(c, nil, nil, &errCode)
		}

		errCode := h.ServiceItemType.DeleteItemTypeById(&req)
		if errCode != nil {
			return fiberpresenter.Presenter(c, nil, nil, errCode)
		}

		return fiberpresenter.Presenter(c, nil, nil, nil)
	}
}
