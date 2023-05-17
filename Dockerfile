FROM golang:alpine AS builder
RUN mkdir /go-api-server
COPY . /go-api-server
WORKDIR /go-api-server
RUN go build .

FROM alpine
WORKDIR /go-api-server
COPY --from=builder /go-api-server/ /go-api-server/

# Expose port 8080 to the outside world
EXPOSE 7127

# Command to run the executable
ENTRYPOINT ["./go-api-server"]
CMD ["server"]