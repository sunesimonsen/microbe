FROM golang:1.26-alpine AS build

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN go build -o microbe .

FROM alpine:3.20

RUN apk add --no-cache ca-certificates

WORKDIR /app
COPY --from=build /app/assets /app/assets
COPY --from=build /app/microbe /app/microbe

EXPOSE 8080
ENTRYPOINT ["/app/microbe"]
