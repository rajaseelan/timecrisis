FROM golang:1.14.4-alpine3.12 AS builder

RUN mkdir /app
ADD . /app
WORKDIR /app
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o timecrisis ./cmd/web
CMD ["/app/timecrisis"]

FROM alpine:3.12 as production
COPY --from=builder /app/timecrisis .
ENV TIMECRISIS_PORT
EXPOSE 4000
CMD ["/timecrisis"]