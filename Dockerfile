FROM golang:1.24 AS baseimage

WORKDIR /app

RUN apt-get update && apt-get install libssl-dev libmcrypt-dev -y

COPY server/go.mod server/go.sum ./server/

RUN go install github.com/air-verse/air@latest

WORKDIR /app/server
RUN go mod download && go mod verify


FROM baseimage AS development

ARG SERVER_PORT=2609
ENV SERVER_PORT=${SERVER_PORT}

WORKDIR /app

COPY .air.toml .
COPY ./server ./server

EXPOSE ${SERVER_PORT}

CMD ["air", "-c", ".air.toml"]