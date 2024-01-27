.PHONY: build

live:
	air

run:
	go run cmd/api/main.go

build:
	go build -o bin/assessment_krisna cmd/api/main.go

container:
	docker build -t assestment-krisna-satriadi .
	docker run -d -p 4000:4000 --network test-network --name assestment-api assestment-krisna-satriadi:latest