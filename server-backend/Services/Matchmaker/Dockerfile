FROM golang:1.24-alpine

ENV GIN_MODE=release
ENV PORT=8080
ENV Environment=production
ENV GO11MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    BASE_PATH=/go/src/app

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o ./service ./cmd

EXPOSE 8080

ENTRYPOINT [ "./service" ]