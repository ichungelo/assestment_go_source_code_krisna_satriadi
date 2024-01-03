package fiberrouter

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

type routerFiberMisc interface {
	NotFound() fiber.Handler
}

type routerFiberCustomer interface {
	CreateCustomer() fiber.Handler
	GetListCustomer() fiber.Handler
	UpdateCustomerById() fiber.Handler
	DeleteCustomerById() fiber.Handler
}

type router struct {
	routerFiberMisc
	routerFiberCustomer
}

func NewRouter(rfMisc routerFiberMisc, rfCustomer routerFiberCustomer) *router {
	return &router{
		routerFiberMisc:  rfMisc,
		routerFiberCustomer: rfCustomer,
	}
}

func (r *router) Route(app fiber.Router, logger func(*fiber.Ctx) error) {
	route := app.Group("api/v1")
	route.Post("/customers", logger, r.CreateCustomer())
	route.Get("/customers", logger, r.GetListCustomer())
	route.Put("/customers/:customerId", logger, r.UpdateCustomerById())
	route.Delete("/customers/:customerId", logger, r.DeleteCustomerById())
	app.Get("/monitor", monitor.New(monitor.Config{Title: "Invoice App Monitoring"}))
	app.Use(r.NotFound())
}
