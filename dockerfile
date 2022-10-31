FROM golang:latest

WORKDIR /var/api

COPY . ./

RUN go mod download
RUN go build ./cmd/api

CMD ["./api"]