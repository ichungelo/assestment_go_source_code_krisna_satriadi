.PHONY: build

live:
	air

run:
	go run cmd/api/main.go

build:
	go build -o bin/assessment_krisna cmd/api/main.go

container:
	docker build -t assessment-krisna-satriadi .
	docker run -d -p 4000:4000 --network test-network --name assessment-api --env-file .env assessment-krisna-satriadi:latest

clean:
	-docker rm -f assessment-api
	-docker rmi assessment-krisna-satriadi:latest
	-rm -rf ./bin
	-rm -rf tmp