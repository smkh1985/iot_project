FROM golang:1.24.3

WORKDIR /app
COPY . .

RUN go mod download
RUN go build -o consumer-influx .

CMD ["./consumer-influx"]