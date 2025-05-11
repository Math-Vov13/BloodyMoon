FROM golang:tip-alpine3.21

ENV GIN_MODE=release
ENV PORT=8080
ENV Environment=production

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o ./bloodymoon-app ./cmd

EXPOSE 8080

ENTRYPOINT [ "./bloodymoon-app" ]