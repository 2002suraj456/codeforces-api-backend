FROM golang:1.24

WORKDIR /app

# Set the entrypoint to run main.go
CMD ["go", "run", "main.go"]