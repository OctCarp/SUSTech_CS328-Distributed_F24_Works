# codebase/golang/db_service/Dockerfile

# Similar to api_service Dockerfile

FROM golang:1.23.3 AS build

WORKDIR /go/src

COPY db_service/ ./db_service
COPY gogrpc/ ./gogrpc

ENV CGO_ENABLED=0

WORKDIR /go/src/db_service

RUN go build -o /go/db_service ./cmd/main.go


FROM scratch AS runtime

COPY --from=build /go/db_service /db_service

EXPOSE 50051

CMD ["/db_service"]
