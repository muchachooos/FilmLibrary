FROM golang:1.21
WORKDIR /app
COPY server .
RUN go mod vendor
RUN go build cmd/main.go
ENTRYPOINT ["./main"]