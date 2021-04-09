# build stage image
FROM golang:1.16 AS build-stage
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64
WORKDIR /build
COPY . .
RUN go mod download
RUN go build -o main .

# production stage image
FROM alpine:latest AS production-stage
RUN apk --no-cache add ca-certificates
COPY --from=build-stage /build /app
EXPOSE 3000
# Command to run when starting the container
CMD ["/app/main"]