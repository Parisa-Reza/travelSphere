FROM golang:1.26-alpine AS builder
WORKDIR /app
RUN apk add --no-cache git
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main .

FROM alpine:latest
WORKDIR /app
RUN apk --no-cache add ca-certificates
COPY --from=builder /app/main .
COPY --from=builder /app/views ./views
COPY --from=builder /app/static ./static 
COPY entrypoint.sh .
RUN chmod +x entrypoint.sh && mkdir -p conf
EXPOSE 8080
ENTRYPOINT ["./entrypoint.sh"]