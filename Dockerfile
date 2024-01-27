FROM golang:1.20

LABEL project=assestment_krisna_satriadi

RUN mkdir /app

WORKDIR /app

ENV APP_NAME=test

# APP_HOST should use 0.0.0.0 due to docker known host
ENV APP_HOST=0.0.0.0 

ENV APP_PORT=4000

ENV APP_STAGE=dev

# DB_HOST should be db container name and whould ber register in same docker network
# to create docker network use command `docker network create [NETWORK NAME]`
# to connect running container to network use `docker network connect  [NETWORK NAME] [CONTAINER NAME]`
# to create container with network connection use flag --network [NETWORK NAME]
ENV DB_HOST=mysql5

ENV DB_PORT=3306

ENV DB_USER=root

ENV DB_PASS=password

COPY . .
COPY .env .

RUN GOPROXY="https://goproxy.io" go mod download

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/main ./cmd/api/main.go

EXPOSE 4000

CMD ["./bin/main"]
