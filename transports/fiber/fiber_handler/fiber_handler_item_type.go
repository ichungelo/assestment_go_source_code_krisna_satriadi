package fiberhandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ichungelo/assestment_go_source_code_krisna_satriadi/core/ports"
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
		return c.Status(fiber.StatusOK).JSON([]string{"a", "b"})
	}
}

func (h *handlerItemType) GetListItemType() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON([]string{"a", "b"})
	}
}

func (h *handlerItemType) UpdateItemTypeById() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON([]string{"a", "b"})
	}
}

func (h *handlerItemType) DeleteItemTypeById() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON([]string{"a", "b"})
	}
}