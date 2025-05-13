# Build Stage
FROM golang:1.24.3 AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# ✅ STATIC BUILD — avoids GLIBC issues
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main .

# Minimal final image
FROM scratch
WORKDIR /root/
COPY --from=builder /app/main /main
COPY --from=builder /app/config.yaml ./config.yaml

CMD ["/main"]
