package fiberhandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ichungelo/assestment_go_source_code_krisna_satriadi/core/model"
	"github.com/ichungelo/assestment_go_source_code_krisna_satriadi/core/ports"
	fiberpresenter "github.com/ichungelo/assestment_go_source_code_krisna_satriadi/transports/fiber/fiber_presenter"
	utilerrors "github.com/ichungelo/assestment_go_source_code_krisna_satriadi/utils/util_errors"
	utillogger "github.com/ichungelo/assestment_go_source_code_krisna_satriadi/utils/util_logger"
	utilvalidator "github.com/ichungelo/assestment_go_source_code_krisna_satriadi/utils/util_validator"
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
		if err != nil {
			utillogger.Error(err, nil)
			httpError := utilerrors.NewHttpError(utilerrors.ErrParseData, err)

			return fiberpresenter.Presenter(c, nil, nil, httpError)
		}

		itemId, err := c.ParamsInt("itemId", 0)
		if err != nil {
			utillogger.Error(err, nil)
			httpError := utilerrors.NewHttpError(utilerrors.ErrParseData, err)

			return fiberpresenter.Presenter(c, nil, nil, httpError)
		}

		req.InvoiceId = invoiceId
		req.ItemId = itemId

		err = utilvalidator.Validate(req)
		if err != nil {
			utillogger.Error(err, nil)
			httpError := utilerrors.NewHttpError(utilerrors.ErrValidate, err)

			return fiberpresenter.Presenter(c, nil, nil, httpError)
		}

		httpError := h.ServiceQuantity.DeleteQuantityById(&req)
		if httpError != nil {
			return fiberpresenter.Presenter(c, nil, nil, httpError)
		}

		return fiberpresenter.Presenter(c, nil, nil, nil)
	}
}
