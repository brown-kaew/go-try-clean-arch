# Dockerfile
FROM golang:1.21.4-alpine AS buildStage

WORKDIR /app

COPY go.mod .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go test -v --tags=unit ./...

RUN go build -o ./out/server .

# ===================

FROM alpine:3.16.2

COPY --from=buildStage /app/out/server /app/server

EXPOSE 2565

CMD ["/app/server"]