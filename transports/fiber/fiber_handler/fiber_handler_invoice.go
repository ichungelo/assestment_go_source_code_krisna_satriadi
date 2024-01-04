package fiberhandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ichungelo/assestment_go_source_code_krisna_satriadi/core/ports"
)

type RouterFiberInvoice interface {
	CreateInvoice() fiber.Handler
	GetListInvoice() fiber.Handler
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
		return c.Status(fiber.StatusOK).JSON([]string{"a", "b"})
	}
}

func (h *handlerInvoice)GetListInvoice() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON([]string{"a", "b"})
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
