# syntax=docker/dockerfile:1

##
## Build
##
FROM golang:1.17.7-alpine3.15 AS build

WORKDIR /build/

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o ./nba-api ./cmd/main.go

##
## Deploy
##
FROM alpine:3.15

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=build /build/configs ./configs
COPY --from=build /build/nba-api ./

EXPOSE 8080

RUN chmod +x ./nba-api

ENTRYPOINT ["./nba-api"]