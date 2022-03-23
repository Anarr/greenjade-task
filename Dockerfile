FROM golang:1.16

WORKDIR /usr/src/app

COPY . .
RUN go get ./.. && go mod tidy

CMD ["go run main.go"]