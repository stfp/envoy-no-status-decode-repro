FROM golang:latest AS builder

WORKDIR /build
COPY ./ .
RUN go get google.golang.org/grpc/internal/transport@v1.35.0
RUN go build -o echoserver

FROM debian:testing-slim
COPY --from=builder /build/echoserver /bin
CMD ["/bin/echoserver"]
