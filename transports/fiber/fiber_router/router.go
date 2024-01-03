package fiberrouter

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

type routerFiberGeneral interface {
	NotFound() fiber.Handler
}

type routerFiberCustomer interface {
	CreateCustomer() fiber.Handler
	GetListCustomer() fiber.Handler
	UpdateCustomerById() fiber.Handler
	DeleteCustomerById() fiber.Handler
}

type router struct {
	routerFiberGeneral
	routerFiberCustomer
}

func NewRouter(rfGeneral routerFiberGeneral, rfCustomer routerFiberCustomer) *router {
	return &router{
		routerFiberGeneral:  rfGeneral,
		routerFiberCustomer: rfCustomer,
	}
}

func (r *router) Route(app fiber.Router, logger func(*fiber.Ctx) error) {
	route := app.Group("api/v1")
	route.Post("/customer", logger, r.CreateCustomer())
	route.Get("/customer", logger, r.GetListCustomer())
	route.Put("/customer/:customerId", logger, r.UpdateCustomerById())
	route.Delete("/customer/:customerId", logger, r.DeleteCustomerById())
	app.Get("/monitor", monitor.New(monitor.Config{Title: "Invoice App Monitoring"}))
	app.Use(r.NotFound())
}
