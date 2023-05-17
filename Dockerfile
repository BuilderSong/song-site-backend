# Use the official Golang image as the parent image
FROM golang:1.20

# Set the working directory to /app
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY . /app

# Install any needed dependencies
RUN go mod download

# Build the Go app
RUN go build -o main .

# Expose port 8080 for the Gin server to listen on
EXPOSE 8080

# Define the command to run the executable
CMD ["./main"]
