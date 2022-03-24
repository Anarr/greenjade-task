FROM golang:1.16

WORKDIR /usr/src/app

COPY . .
RUN go mod download
RUN go build -o app main.go

CMD ["./app"]