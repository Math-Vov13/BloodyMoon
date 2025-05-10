FROM golang:alpine

ENV GIN_MODE=release
ENV PORT=8080

WORKDIR /app

COPY . /app/

# Run the two commands below to install git and dependencies for the project. 
RUN apk update && apk add --no-cache git
RUN go get .

RUN go build .

EXPOSE 8080

ENTRYPOINT ["./BloodyMoon"]
