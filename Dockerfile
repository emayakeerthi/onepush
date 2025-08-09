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

FROM golang:1.24 AS production

ARG SERVER_PORT=2609
ENV SERVER_PORT=${SERVER_PORT}

WORKDIR /app

# Install system dependencies
RUN apt-get update && apt-get install -y libssl-dev libmcrypt-dev git

# Copy the entire server directory
COPY ./server ./

# Download dependencies
RUN go mod download && go mod verify

# Debug: Show Go environment
RUN go env

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -buildvcs=false -v -a -ldflags "-s -w" -o onepushserver ./cmd/server

EXPOSE ${SERVER_PORT}

# Make the binary executable
RUN chmod +x onepushserver

CMD ["./onepushserver"]