FROM golang:1.13

WORKDIR /basicauth
COPY . .
RUN go test ./...
