FROM golang:1.20-alpine

WORKDIR /app

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh

COPY go.mod go.sum ./

RUN go mod download

COPY . ./

RUN go build -o /server

EXPOSE 8080

CMD [ "/server" ]