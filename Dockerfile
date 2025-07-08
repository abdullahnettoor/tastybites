# Stage 1
FROM golang:1.23-alpine AS builder

WORKDIR /app

RUN apk update && apk upgrade && apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o tastybites-api cmd/tastybites/main.go

# Stage 2
FROM alpine:3.19

WORKDIR /app

# Install curl for healthcheck
RUN apk --no-cache add curl

COPY --from=builder /app/tastybites-api .
COPY --from=builder /app/.env .

EXPOSE 8080

CMD ["./tastybites-api"]