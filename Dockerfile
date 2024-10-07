FROM golang:1.23-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags '-s -w' -tags timetzdata -o cacher ./cmd/main.go

FROM scratch
COPY --from=builder /app/cacher /app/cacher
WORKDIR /app
ENTRYPOINT ["/app/cacher"]
