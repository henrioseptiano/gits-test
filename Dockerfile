FROM golang:1.20-alpine AS builder

RUN apk update && apk add --no-cache git

WORKDIR /opt/gits/

COPY . .

ENV GOSUMDB=off
COPY go.mod .
COPY go.sum .
RUN go mod download
RUN go mod verify

RUN GOOS=linux GOARCH=amd64 go build -o /go/bin/gits

FROM alpine:3.15

RUN apk add --no-cache tzdata

# Add the .env file to the image
COPY .env /app/.env

# Source the .env file to set environment variables
RUN set -o allexport; source /app/.env; set +o allexport

COPY --from=builder /go/bin/gits /go/bin/gits

ENV FIBER_MODE=release
ENV TZ=Asia/Jakarta

ENTRYPOINT ["/go/bin/gits"]
