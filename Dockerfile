FROM golang:1.20 as builder
WORKDIR /app
copy go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN CGO_ENABLED=0 go build -v -ldflags="-s -w" -o /app/ ./...

FROM scratch

WORKDIR /app
COPY --from=builder /app/simple-single-share /app

EXPOSE 8080

ENTRYPOINT ["/app/simple-single-share", "server"]
