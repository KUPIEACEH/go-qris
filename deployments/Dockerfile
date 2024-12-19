FROM scratch

# Copy compiled binary into the container
COPY go-qris /go-qris

# Expose port to the outside world
EXPOSE 1337

# Set default command to run application
ENTRYPOINT ["/go-qris","run"]