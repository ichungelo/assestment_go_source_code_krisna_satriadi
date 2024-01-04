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

func (r *Router) Route(app fiber.Router) {
	route := app.Group("/api/v1")

	//! customer
	customer := route.Group("/customers")
	customer.Post("/", r.CreateCustomer())
	customer.Get("/", r.GetListCustomer())
	customer.Put("/:customerId", r.UpdateCustomerById())
	customer.Delete("/:customerId", r.DeleteCustomerById())

	//! invoice
	invoice := route.Group("/invoices")
	invoice.Post("/", r.CreateInvoice())
	invoice.Get("/", r.GetListInvoice())
	invoice.Get("/:invoiceId", r.GetInvoiceById())
	invoice.Put("/:invoiceId", r.UpdateInvoiceById())
	invoice.Delete("/:invoiceId", r.DeleteInvoiceById())

	//! ItemType
	itemType := route.Group("/types")
	itemType.Post("/", r.CreateItemType())
	itemType.Get("/", r.GetListItemType())
	itemType.Put("/:itemTypeId", r.UpdateItemTypeById())
	itemType.Delete("/:itemTypeId", r.DeleteItemTypeById())

	//! ItemType
	item := route.Group("/items")
	item.Post("/", r.CreateItem())
	item.Get("/", r.GetListItem())
	item.Put("/:itemId", r.UpdateItemById())
	item.Delete("/:itemId", r.DeleteItemById())

	//! misc
	app.Get("/monitor", monitor.New(monitor.Config{Title: "Invoice App Monitoring"}))
	app.Use(r.NotFound())
}
