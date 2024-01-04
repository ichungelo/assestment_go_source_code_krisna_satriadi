package fiberhandler

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/ichungelo/assestment_go_source_code_krisna_satriadi/core/model"
	"github.com/ichungelo/assestment_go_source_code_krisna_satriadi/core/ports"
	"github.com/ichungelo/assestment_go_source_code_krisna_satriadi/utils"
)

type RouterFiberInvoice interface {
	CreateInvoice() fiber.Handler
	GetListInvoice() fiber.Handler
	GetInvoiceById() fiber.Handler
	UpdateInvoiceById() fiber.Handler
	DeleteInvoiceById() fiber.Handler
}

type handlerInvoice struct {
	ports.ServiceInvoice
}

func NewInvoiceHandler(sInvoice ports.ServiceInvoice) *handlerInvoice {
	return &handlerInvoice{
		ServiceInvoice: sInvoice,
	}
}

func (h *handlerInvoice)CreateInvoice() fiber.Handler {
	return func(c *fiber.Ctx) error {
		req := model.RequestCreateInvoice{}

		err := json.Unmarshal(c.Body(), &req)
		if err != nil {
			errCode := utils.ErrorCode{
				Code: utils.ERR_FAILED_UNMARSHAL_JSON,
				Err:  err,
			}

			return model.Presenter(c, nil, nil, &errCode)
		}

		err = utils.Validate(req)
		if err != nil {
			errCode := utils.ErrorCode{
				Code: utils.ERR_VALIDATE_STRUCT,
				Err:  err,
			}

			return model.Presenter(c, nil, nil, &errCode)
		}

		errCode := h.ServiceInvoice.CreateInvoice(&req)
		if errCode != nil {
			return model.Presenter(c, nil, nil, errCode)
		}

		return model.Presenter(c, nil, nil, nil)
	}
}

func (h *handlerInvoice)GetListInvoice() fiber.Handler {
	return func(c *fiber.Ctx) error {
		req := model.RequestGetListInvoice{}

		err := utils.Validate(req)
		if err != nil {
			errCode := utils.ErrorCode{
				Code: utils.ERR_VALIDATE_STRUCT,
				Err:  err,
			}

			return model.Presenter(c, nil, nil, &errCode)
		}

		data, errCode := h.ServiceInvoice.GetListInvoice(&req)
		if errCode != nil {
			return model.Presenter(c, nil, nil, errCode)
		}

		return model.Presenter(c, data, nil, nil)
	}
}

func (h *handlerInvoice)GetInvoiceById() fiber.Handler {
	return func(c *fiber.Ctx) error {
		req := model.RequestGetListInvoice{}

		err := utils.Validate(req)
		if err != nil {
			errCode := utils.ErrorCode{
				Code: utils.ERR_VALIDATE_STRUCT,
				Err:  err,
			}

			return model.Presenter(c, nil, nil, &errCode)
		}

		data, errCode := h.ServiceInvoice.GetListInvoice(&req)
		if errCode != nil {
			return model.Presenter(c, nil, nil, errCode)
		}

		return model.Presenter(c, data, nil, nil)
	}
}

func (h *handlerInvoice)UpdateInvoiceById() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON([]string{"a", "b"})
	}	
}

func (h *handlerInvoice)DeleteInvoiceById() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON([]string{"a", "b"})
	}
}
