.PHONY: build

live:
	air

run:
	go run cmd/api/main.go

build:
	go build -o bin/assessment_krisna cmd/api/main.go