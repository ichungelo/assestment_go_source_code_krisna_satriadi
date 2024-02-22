package di

import (
	"github.com/ichungelo/assestment_go_source_code_krisna_satriadi/config"
	"github.com/ichungelo/assestment_go_source_code_krisna_satriadi/core/services"
	gormadapter "github.com/ichungelo/assestment_go_source_code_krisna_satriadi/sources/gorm/gorm_adapter"
	fiberhandler "github.com/ichungelo/assestment_go_source_code_krisna_satriadi/transports/fiber/fiber_handler"
	fiberrouter "github.com/ichungelo/assestment_go_source_code_krisna_satriadi/transports/fiber/fiber_router"
	gormconnection "github.com/ichungelo/assestment_go_source_code_krisna_satriadi/sources/gorm/gorm_connection"
)

func Initializer(cfg *config.Config) *fiberrouter.Router {
	var (
		gormClient = gormconnection.GetInstanceDB(cfg)

		gormAdapter = gormadapter.NewGormAdapter(gormClient)

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

		sQuantity = services.NewServiceQuantity(gormAdapter)
		hQuantity = fiberhandler.NewQuantityHandler(sQuantity)
	)

	router := fiberrouter.NewRouter(hCustomer, hInvoice, hItemType, hItem, hMisc, hQuantity)
	return router

}
