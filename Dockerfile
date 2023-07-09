FROM golang:latest AS builder
COPY ./app /app
WORKDIR /app
RUN go mod download

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -a -o /app .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /app ./
RUN chmod +x ./app
ENTRYPOINT ["./app"]
EXPOSE 80