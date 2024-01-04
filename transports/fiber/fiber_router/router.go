package fiberrouter

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	fiberhandler "github.com/ichungelo/assestment_go_source_code_krisna_satriadi/transports/fiber/fiber_handler"
)

type Router struct {
	fiberhandler.RouterFiberCustomer
	fiberhandler.RouterFiberInvoice
	fiberhandler.RouterFiberItemType
	fiberhandler.RouterFiberItem
	fiberhandler.RouterFiberMisc
}

func NewRouter(rfCustomer fiberhandler.RouterFiberCustomer, rfInvoice fiberhandler.RouterFiberInvoice, rfItemType fiberhandler.RouterFiberItemType, rfItem fiberhandler.RouterFiberItem, rfMisc fiberhandler.RouterFiberMisc) *Router {
	return &Router{
		RouterFiberCustomer: rfCustomer,
		RouterFiberInvoice:  rfInvoice,
		RouterFiberItemType: rfItemType,
		RouterFiberItem:     rfItem,
		RouterFiberMisc:     rfMisc,
	}
}

func (r *Router) Route(app fiber.Router, logger func(*fiber.Ctx) error) {
	route := app.Group("api/v1")

	//! customer
	route.Post("/customers", logger, r.CreateCustomer())
	route.Get("/customers", logger, r.GetListCustomer())
	route.Put("/customers/:customerId", logger, r.UpdateCustomerById())
	route.Delete("/customers/:customerId", logger, r.DeleteCustomerById())

	//! invoice
	route.Post("/invoices", logger, r.CreateInvoice())
	route.Get("/invoices", logger, r.GetListInvoice())
	route.Put("/invoices/:invoicesId", logger, r.UpdateInvoiceById())
	route.Delete("/invoices/:invoicesId", logger, r.DeleteInvoiceById())

	//! misc
	app.Get("/monitor", monitor.New(monitor.Config{Title: "Invoice App Monitoring"}))
	app.Use(r.NotFound())
}
