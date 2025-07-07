# Stage 1
FROM golang:1.21.5-alpine3.19 AS builder

WORKDIR /app

RUN apk update && apk upgrade && apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o tastybites-api ./...

# Stage 2
FROM alpine:3.19

WORKDIR /app

COPY --from=builder /app/tastybites-api .

EXPOSE 8080

CMD ["./tastybites-api"]