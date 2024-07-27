# Build
FROM golang:1.22 AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

ENV CGO_ENABLED 0
ENV GOOS linux
RUN go build -o pokecalc cmd/server/main.go

# Deploy
# hadolint ignore=DL3007
FROM gcr.io/distroless/base-debian11:latest

WORKDIR /

COPY --from=builder /app/pokecalc .

EXPOSE 8080

# アプリケーションを実行
ENTRYPOINT ["./pokecalc"]
