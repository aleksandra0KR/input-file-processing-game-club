FROM golang:1.22 AS builder

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . .

ENV GOOS linux

RUN go build -o main ./cmd/main.go

FROM alpine AS runner

WORKDIR /root/

COPY --from=builder /app/main .

ENTRYPOINT ["./main"]