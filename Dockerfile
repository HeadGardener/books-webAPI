FROM golang:1.18.3

RUN go version
ENV GOPATH=/

COPY ./ ./


# build go app
RUN go mod download
RUN go build -o books-webAPI ./cmd/main.go

CMD ["./books-webAPI"]