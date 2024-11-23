# codebase/golang/api_service/Dockerfile

# Use golang as the build image
FROM golang:1.23.3 AS build

WORKDIR /go/src

# Copy source code
COPY api_service/ ./api_service
# Also copy gRPC module
COPY gogrpc/ ./gogrpc

# Disable CGO for static pure Go binary
ENV CGO_ENABLED=0

WORKDIR /go/src/api_service

# Build the application binary
RUN go build -o /go/openapi ./cmd/main.go

# Start a new stage with minimal scratch image
FROM scratch AS runtime

# Set Gin framework to release mode
ENV GIN_MODE=release

# Copy binary from build stage
COPY --from=build /go/openapi /openapi

# Expose API port
EXPOSE 8080

# Set the entry point
ENTRYPOINT ["/openapi"]