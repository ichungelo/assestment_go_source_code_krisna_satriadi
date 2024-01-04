package di

import (
	"github.com/ichungelo/assestment_go_source_code_krisna_satriadi/core/services"
	gormadapter "github.com/ichungelo/assestment_go_source_code_krisna_satriadi/sources/gorm/gorm_adapter"
	fiberhandler "github.com/ichungelo/assestment_go_source_code_krisna_satriadi/transports/fiber/fiber_handler"
	fiberrouter "github.com/ichungelo/assestment_go_source_code_krisna_satriadi/transports/fiber/fiber_router"
	"gorm.io/gorm"
)

func Initializer(gormDB *gorm.DB) *fiberrouter.Router {
	var (
		gormAdapter = gormadapter.NewGormAdapter(gormDB)

		sCustomer = services.NewServiceCustomer(gormAdapter)
		hCustomer = fiberhandler.NewCustomerHandler(sCustomer)

		sInvoice = services.NewServiceInvoice(gormAdapter)
		hInvoice = fiberhandler.NewInvoiceHandler(sInvoice)

		sItemType = services.NewServiceItemType(gormAdapter)
		hItemType = fiberhandler.NewItemTypeHandler(sItemType)

		sItem = services.NewServiceItem(gormAdapter)
		hItem = fiberhandler.NewItemHandler(sItem)

		sMisc = services.NewServiceMisc(gormAdapter)
		hMisc = fiberhandler.NewMiscHandler(sMisc)
	)

	router := fiberrouter.NewRouter(hCustomer, hInvoice, hItemType, hItem, hMisc)
	return router

}
