FROM golang:1.14

# Set the Current Working Directory inside the container
WORKDIR app

# Copy everything from the current directory to the PWD (Present Working Directory) inside the container
COPY . .

# Install the package
RUN go build cmd/main.go

# This container exposes port 8080 to the outside world
EXPOSE 8080

# Run the executable
CMD ["main"]