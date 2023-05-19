FROM golang:1.20 as builder
WORKDIR /app
copy go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /app/ ./...

FROM golang:1.20

WORKDIR /app
COPY --from=builder /app/simple-single-share /app/

ENTRYPOINT ["/app/simple-single-share", "server"]
