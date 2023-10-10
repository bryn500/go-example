# Use a minimal Go image as the base
FROM golang:1.21.2 AS build

# Set the working directory in the container
WORKDIR /app

# Copy the Go application source code into the container
COPY . .

# Build the Go application
RUN go build -o api

# Create a lightweight final image
FROM alpine:latest

# Set the working directory in the final image
WORKDIR /app

# Copy the binary from the build image to the final image
COPY --from=build /app/api ./api

# Expose the port your application will run on
EXPOSE 8080

# Define the command to run your application
CMD ["/app/api"]
