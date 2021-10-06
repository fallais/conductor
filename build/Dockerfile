# Build
FROM golang:latest as builder
WORKDIR /go/src/conductor
ADD . /go/src/conductor
RUN go get -d -v ./...
RUN go build -o /go/bin/conductor

# Final
FROM gcr.io/distroless/base-debian10
COPY --from=builder /go/bin/conductor /
ENTRYPOINT [ "/conductor" ]
CMD [ "--help" ]
