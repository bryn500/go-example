FROM golang:1.21.2

# Set the working directory in the container
WORKDIR /app

# Copy the Go application source code into the container
COPY . .

# Build the Go application
RUN go build -o api

EXPOSE 8080

CMD ["./api"]