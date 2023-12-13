# Use an official Golang runtime as a parent image
FROM golang:1.21

# Set the working directory to /app
WORKDIR /app

# Copy the current directory contents into the container at /app
ADD . /app

# Build the Golang app
RUN go build -o main .

# Expose port 8090
EXPOSE 8090

# Run main when the container launches
CMD ["./main"]
