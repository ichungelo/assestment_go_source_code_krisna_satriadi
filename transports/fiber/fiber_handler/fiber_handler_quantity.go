package fiberhandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ichungelo/assestment_go_source_code_krisna_satriadi/core/model"
	"github.com/ichungelo/assestment_go_source_code_krisna_satriadi/core/ports"
	fiberpresenter "github.com/ichungelo/assestment_go_source_code_krisna_satriadi/transports/fiber/fiber_presenter"
	"github.com/ichungelo/assestment_go_source_code_krisna_satriadi/utils"
)

type RouterFiberQuantity interface {
	DeleteQuantityById() fiber.Handler
}

type handlerQuantity struct {
	ports.ServiceQuantity
}

func NewQuantityHandler(sQuantity ports.ServiceQuantity) *handlerQuantity {
	return &handlerQuantity{
		ServiceQuantity: sQuantity,
	}
}

func (h *handlerQuantity) DeleteQuantityById() fiber.Handler {
	return func(c *fiber.Ctx) error {
		req := model.RequestDeleteQuantityById{}

		invoiceId, err := c.ParamsInt("invoiceId", 0)
		itemId, err := c.ParamsInt("itemId", 0)
		if err != nil {
			utils.Error(err, nil)
			errCode := utils.ErrorCode{
				Code: utils.ERR_PARSE_DATA,
				Err:  err,
			}

			return fiberpresenter.Presenter(c, nil, nil, &errCode)
		}

		req.InvoiceId = invoiceId
		req.ItemId = itemId

		err = utils.Validate(req)
		if err != nil {
			utils.Error(err, nil)
			errCode := utils.ErrorCode{
				Code: utils.ERR_VALIDATE,
				Err:  err,
			}

			return fiberpresenter.Presenter(c, nil, nil, &errCode)
		}

		errCode := h.ServiceQuantity.DeleteQuantityById(&req)
		if errCode != nil {
			return fiberpresenter.Presenter(c, nil, nil, errCode)
		}

		return fiberpresenter.Presenter(c, nil, nil, nil)
	}
}
