FROM golang:1.23-alpine AS builder

WORKDIR /app

# copy and download all necessary modules
COPY go.mod .
COPY go.sum .

RUN go mod download

# copy all local files to Image
COPY . .

# build Docker Image with CGO Enabled
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o go-qris ./cmd

FROM alpine:latest

WORKDIR /app

# get Executable Binary file to new Image
COPY --from=builder /app/go-qris ./

# run the server executable
CMD ["./go-qris"]
