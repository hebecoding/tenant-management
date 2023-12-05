FROM golang:1.21.5-alpine3.17

WORKDIR /app

COPY . .

RUN go build -o main cmd/tenant-management/

EXPOSE 8443

CMD ["/app/main"]