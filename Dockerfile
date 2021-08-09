# Start from golang:1.15-alpine base image
FROM golang:1.15-alpine3.14
# Adding git, bash and openssh to the image
RUN apk update && apk upgrade && apk add --no-cache bash git openssh
# Set the current Working Directory inside the container
WORKDIR /server
# Download all dependencies
RUN go get -u github.com/gorilla/websocket
RUN go get -u github.com/pkg/errors
# Copy the source from the current directory to the Working Directory inside the container
COPY . .
# Build the Go app
RUN go build -o main main.go
# Expose port 8080 to the outside world
EXPOSE 8080
# Run the executable
CMD ["./main"]
