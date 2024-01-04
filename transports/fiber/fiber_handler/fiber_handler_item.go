package fiberhandler

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/ichungelo/assestment_go_source_code_krisna_satriadi/core/model"
	"github.com/ichungelo/assestment_go_source_code_krisna_satriadi/core/ports"
	fiberpresenter "github.com/ichungelo/assestment_go_source_code_krisna_satriadi/transports/fiber/fiber_presenter"
	"github.com/ichungelo/assestment_go_source_code_krisna_satriadi/utils"
)

type RouterFiberItem interface {
	CreateItem() fiber.Handler
	GetListItem() fiber.Handler
	UpdateItemById() fiber.Handler
	DeleteItemById() fiber.Handler
}

type handlerItem struct {
	ports.ServiceItem
}

func NewItemHandler(sItem ports.ServiceItem) *handlerItem {
	return &handlerItem{
		ServiceItem: sItem,
	}
}

func (h *handlerItem) CreateItem() fiber.Handler {
	return func(c *fiber.Ctx) error {
		req := model.RequestCreateItem{}

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

		errCode := h.ServiceItem.CreateItem(&req)
		if errCode != nil {
			return fiberpresenter.Presenter(c, nil, nil, errCode)
		}

		return fiberpresenter.Presenter(c, nil, nil, nil)
	}
}

func (h *handlerItem) GetListItem() fiber.Handler {
	return func(c *fiber.Ctx) error {

		data, errCode := h.ServiceItem.GetListItem()
		if errCode != nil {
			return fiberpresenter.Presenter(c, nil, nil, errCode)
		}

		return fiberpresenter.Presenter(c, data, nil, nil)
	}
}

func (h *handlerItem) UpdateItemById() fiber.Handler {
	return func(c *fiber.Ctx) error {
		req := model.RequestUpdateItemById{}

		err := json.Unmarshal(c.Body(), &req)
		if err != nil {
			utils.Error(err, nil)
			errCode := utils.ErrorCode{
				Code: utils.ERR_FAILED_UNMARSHAL_JSON,
				Err:  err,
			}

			return fiberpresenter.Presenter(c, nil, nil, &errCode)
		}

		itemId, err := c.ParamsInt("itemId", 0)
		if err != nil {
			utils.Error(err, nil)
			errCode := utils.ErrorCode{
				Code: utils.ERR_PARSE_DATA,
				Err:  err,
			}

			return fiberpresenter.Presenter(c, nil, nil, &errCode)
		}

		req.ItemId = itemId

		err = utils.Validate(req)
		if err != nil {
			utils.Error(err, nil)
			errCode := utils.ErrorCode{
				Code: utils.ERR_VALIDATE_STRUCT,
				Err:  err,
			}

			return fiberpresenter.Presenter(c, nil, nil, &errCode)
		}

		errCode := h.ServiceItem.UpdateItemById(&req)
		if errCode != nil {
			return fiberpresenter.Presenter(c, nil, nil, errCode)
		}

		return fiberpresenter.Presenter(c, nil, nil, nil)
	}
}

func (h *handlerItem) DeleteItemById() fiber.Handler {
	return func(c *fiber.Ctx) error {
		req := model.RequestDeleteItemById{}

		itemId, err := c.ParamsInt("itemId", 0)
		if err != nil {
			utils.Error(err, nil)
			errCode := utils.ErrorCode{
				Code: utils.ERR_PARSE_DATA,
				Err:  err,
			}

			return fiberpresenter.Presenter(c, nil, nil, &errCode)
		}

		req.ItemId = itemId

		err = utils.Validate(req)
		if err != nil {
			utils.Error(err, nil)
			errCode := utils.ErrorCode{
				Code: utils.ERR_VALIDATE_STRUCT,
				Err:  err,
			}

			return fiberpresenter.Presenter(c, nil, nil, &errCode)
		}

		errCode := h.ServiceItem.DeleteItemById(&req)
		if errCode != nil {
			return fiberpresenter.Presenter(c, nil, nil, errCode)
		}

		return fiberpresenter.Presenter(c, nil, nil, nil)
	}
}
