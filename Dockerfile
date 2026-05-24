FROM golang:1.16.3-alpine3.13

WORKDIR /backend

COPY . .

RUN go get -d -v ./...

RUN go build ./cmd/api -o subscriber-mgmt-backend .

EXPOSE 8000

CMD ["./subscriber-mgmt-backend"]
