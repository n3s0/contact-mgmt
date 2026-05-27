FROM golang:1.26.3-alpine3.13

WORKDIR /backend

COPY . .

RUN go get -d -v ./...

RUN go build ./cmd/api -o contact-mgmt-api .

EXPOSE 8000

CMD ["./contact-mgmt-api"]
