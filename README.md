# Assestment Go Source Code PT Essensi Solusi Buana

## Prerequisites

- [Go Programming Language](https://go.dev/dl/) v1.20.x
- [MySQL](https://www.mysql.com/downloads/) v5.7.x
- [AIR (Go Live Reloader)](https://github.com/cosmtrek/air)
- Your favorite code editor

## Installation & Run

Before running API server, you should set your env. Your DB will automatically generated and the name will be combination of your app name and your stage. (ex: myapp_dev)

```bash
# Run
cd assestment_go_source_code_krisna_satriadi
make run
```

```bash
# Run (with live reload)
cd assestment_go_source_code_krisna_satriadi
make live
```

```bash
# Build
cd assestment_go_source_code_krisna_satriadi
make build
```

## Structure

```
.
├── cmd
│   └── api
│       └── main.go								// main package
├── config										// config control
│   └── config.go
├── core
│   ├── model									// data model
│   │   ├── model_customer.go
│   │   ├── model_invoice.go
│   │   ├── model_item.go
│   │   ├── model_item_type.go
│   │   ├── model_misc.go
│   │   └── model_quantity.go
│   ├── ports									// ports/interface
│   │   ├── ports_customer.go
│   │   ├── ports_invoice.go
│   │   ├── ports_item.go
│   │   ├── ports_item_type.go
│   │   ├── ports_misc.go
│   │   └── ports_quantity.go
│   └── services								// service layer
│       ├── service_customer.go
│       ├── service_invoice.go
│       ├── service_item.go
│       ├── service_item_type.go
│       ├── service_misc.go
│       └── service_quantity.go
├── di											// dependency injector
│   └── di.go
├── documentation								//	postman documentation
│   ├── assestment.postman_collection.json
│   └── assestment.postman_environment.json
├── go.mod
├── go.sum
├── Makefile
├── README.md									// YOU ARE HERE
├── sources										// data sources directory
│   └── gorm
│       ├── gorm_adapter						// gorm adapter/repository
│       │   ├── gorm_adapter_base.go
│       │   ├── gorm_adapter_customer.go
│       │   ├── gorm_adapter_invoice.go
│       │   ├── gorm_adapter_item.go
│       │   ├── gorm_adapter_item_type.go
│       │   └── gorm_adapter_quantity.go
│       └── gorm_connection						// gorm connector starter
│           └── gorm_connection.go
├── transports									// data presenter directory
│   └── fiber
│       ├── fiber_handler						// restAPI handler
│       │   ├── fiber_handler_customer.go
│       │   ├── fiber_handler_invoice.go
│       │   ├── fiber_handler_item.go
│       │   ├── fiber_handler_item_type.go
│       │   ├── fiber_handler_misc.go
│       │   └── fiber_handler_quantity.go
│       ├── fiber_middleware					// API middleware
│       │   └── middleware.go
│       ├── fiber_presenter						// json presenter
│       │   └── fiber_presenter.go
│       └── fiber_router						// http router
│           └── router.go
└── utils										// utils storage
    ├── checker.go
    ├── errors.go
    ├── logger.go
    ├── ulid.go
    └── validator.go
```

## API

#### /projects

- `GET` : Get all projects
- `POST` : Create a new project

#### /projects/:title

- `GET` : Get a project
- `PUT` : Update a project
- `DELETE` : Delete a project

#### /projects/:title/archive

- `PUT` : Archive a project
- `DELETE` : Restore a project

#### /projects/:title/tasks

- `GET` : Get all tasks of a project
- `POST` : Create a new task in a project

#### /projects/:title/tasks/:id

- `GET` : Get a task of a project
- `PUT` : Update a task of a project
- `DELETE` : Delete a task of a project

#### /projects/:title/tasks/:id/complete

- `PUT` : Complete a task of a project
- `DELETE` : Undo a task of a project

