FROM scratch

# Copy compiled binary into the container
COPY go-qris /go-qris

# Set default command to run application
ENTRYPOINT ["/go-qris","run"]