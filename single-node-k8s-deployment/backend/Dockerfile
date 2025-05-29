FROM golang:1.24.2
WORKDIR /app
COPY main.go .
RUN go build -o server main.go
CMD ["./server"]
