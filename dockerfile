FROM golang:1.20-alpine3.16
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go build -o main cmd/main.go
CMD ["/app/main"]



