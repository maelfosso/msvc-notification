FROM golang:1.15.6 AS builder
# golang:1.15.6-alpine3.13 AS builder 
# golang:1.15.6 AS builder

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
  CGO_ENABLED=0 \
  GOOS=linux \
  GOARCH=amd64

WORKDIR /
ADD wait-for-it.sh wait-for-it.sh
RUN chmod +x /wait-for-it.sh

# Move to working directory /build
WORKDIR /src

# Copy and download dependency using go mod
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy the code into the container
COPY . .

# Build the application
RUN go build -o main .

CMD ["/src/main"]
