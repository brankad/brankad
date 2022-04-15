# Start from the latest golang base image
FROM golang:1.17.9-alpine AS build_base
# Add Maintainer Info

RUN apk add --no-cache git

# Set the Current Working Directory inside the container
WORKDIR /app
COPY . .

# Download all dependencies
RUN go mod download
RUN go mod tidy

# Unit tests
RUN CGO_ENABLED=0 go test -v ./api
# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

CMD ["./main"]


