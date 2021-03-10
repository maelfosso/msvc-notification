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

# Build a small image
# FROM busybox

# COPY --from=builder /dist/main /

# Command to run
# ENTRYPOINT ["/main"]

# FROM golang:1.15.6
# # github.com/maelfosso/msvc-notification
# WORKDIR /go/src/guitou.com/notification-msvc
# RUN go get -d -v golang.org/x/net/html  
# COPY main.go .
# RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# FROM alpine:latest  
# RUN apk --no-cache add ca-certificates
# WORKDIR /root/
# COPY --from=0 /go/src/guitou.com/notification-msvc .
# CMD ["./main"]  
