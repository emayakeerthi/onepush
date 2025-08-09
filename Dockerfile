FROM golang:1.24 AS baseimage

WORKDIR /app

RUN apt-get update && apt-get install libssl-dev libmcrypt-dev git -y

COPY server/go.mod server/go.sum ./server/

RUN go install github.com/air-verse/air@latest

WORKDIR /app/server
RUN go mod download && go mod verify


FROM baseimage AS development

ARG SERVER_PORT=2609
ENV SERVER_PORT=${SERVER_PORT}
ENV CGO_ENABLED=0
ENV GOOS=linux

WORKDIR /app

COPY .air.toml .
COPY ./server ./server

EXPOSE ${SERVER_PORT}

CMD ["air", "-c", ".air.toml"]

FROM baseimage AS production

ARG SERVER_PORT=2609
ENV SERVER_PORT=${SERVER_PORT}
ENV CGO_ENABLED=0
ENV GOOS=linux

WORKDIR /app

COPY ./server ./server

WORKDIR /app/server
RUN go build -buildvcs=false -a -ldflags '-extldflags "-static"' -o /app/server-binary ./cmd/server

WORKDIR /app

EXPOSE ${SERVER_PORT}

CMD ["/app/server-binary"]