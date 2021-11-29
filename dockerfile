FROM golang:alpine AS builder

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Move to working directory /build
WORKDIR /build

# Copy and download dependency using go mod
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy the code into the container
COPY . .

# Build the application
RUN  go build -o main .

# Move to /dist directory as the place for resulting binary folder
WORKDIR /dist

# Copy binary from build to main folder
RUN cp /build/main .
RUN cp /build/config.json .

# Build a small image
FROM scratch

COPY --from=builder /dist/* /

ENV JAEGER_AGENT_HOST=jaeger
ENV JAEGER_AGENT_PORT=6831

# command to run with jaeger
ENTRYPOINT ["./main"]