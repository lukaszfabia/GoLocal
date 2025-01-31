FROM golang:1.23-bookworm AS builder

WORKDIR /build
COPY go.mod go.sum golocal-firebase.json ./
RUN go mod download

COPY . .

RUN go build -o backend ./cmd/api

FROM alpine:latest AS run

RUN apk add --no-cache libc6-compat

WORKDIR /app
COPY --from=builder /build/golocal-firebase.json /app/
COPY --from=builder /build/backend .
COPY --from=builder /build/media /app/media 

RUN chmod +x ./backend

CMD ["./backend"]
