package fiberhandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ichungelo/assestment_go_source_code_krisna_satriadi/core/ports"
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
		return c.Status(fiber.StatusOK).JSON([]string{"a", "b"})
	}
}

func (h *handlerItem) GetListItem() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON([]string{"a", "b"})
	}
}

func (h *handlerItem) UpdateItemById() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON([]string{"a", "b"})
	}
}

func (h *handlerItem) DeleteItemById() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON([]string{"a", "b"})
	}
}