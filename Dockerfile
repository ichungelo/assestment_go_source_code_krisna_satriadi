# ---------------------- make images from repo --------------------------

FROM golang:1.20

LABEL project=assestment_krisna_satriadi

RUN mkdir /app

WORKDIR /app

COPY . .

ENV APP_NAME=${APP_NAME}

# APP_HOST should use 0.0.0.0 due to docker known host
ENV APP_HOST=${APP_HOST} 

ENV APP_PORT=${APP_PORT}

ENV APP_STAGE=${APP_STAGE}

# DB_HOST should be db container name and whould ber register in same docker network
# to create docker network use command `docker network create [NETWORK NAME]`
# to connect running container to network use `docker network connect  [NETWORK NAME] [CONTAINER NAME]`
# to create container with network connection use flag --network [NETWORK NAME]
ENV DB_HOST=${DB_HOST}

ENV DB_PORT=${DB_PORT}

ENV DB_USER=${DB_USER}

ENV DB_PASS=${DB_PASS}

RUN GOPROXY="https://goproxy.io" go mod download

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/main ./cmd/api/main.go

EXPOSE ${APP_PORT}

CMD ["./bin/main"]

# ---------------------- make images from compiled bin file --------------------------

# FROM ubuntu:latest

# LABEL project=assessment_krisna_satriadi

# RUN mkdir /app

# WORKDIR /app

# COPY ./bin/assessment_krisna .

# ENV APP_NAME=${APP_NAME}

# # APP_HOST should use 0.0.0.0 due to docker known host
# ENV APP_HOST=${APP_HOST} 

# ENV APP_PORT=${APP_PORT}

# ENV APP_STAGE=${APP_STAGE}

# # DB_HOST should be db container name and whould ber register in same docker network
# # to create docker network use command `docker network create [NETWORK NAME]`
# # to connect running container to network use `docker network connect  [NETWORK NAME] [CONTAINER NAME]`
# # to create container with network connection use flag --network [NETWORK NAME]
# ENV DB_HOST=${DB_HOST}

# ENV DB_PORT=${DB_PORT}

# ENV DB_USER=${DB_USER}

# ENV DB_PASS=${DB_PASS}

# EXPOSE ${APP_PORT}

# CMD ["./assessment_krisna"]
