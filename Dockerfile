FROM golang:1.20.3-alpine3.17

WORKDIR /app

COPY . .

RUN go build -o main cmd/tenant-management/

EXPOSE 8443

CMD ["/app/main"]