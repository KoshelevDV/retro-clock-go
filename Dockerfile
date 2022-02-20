FROM golang:1.17.5-alpine3.15 as build
WORKDIR /app
COPY . ./
RUN go build -o target/go-clock ./cmd/clock/main.go

FROM alpine:latest
COPY --from=build /app/target/go-clock /usr/bin
RUN chmod +x /usr/bin/go-clock && \
    apk update && \
    apk add --no-cache tzdata
ENTRYPOINT [ "go-clock" ]
