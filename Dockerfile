FROM golang:latest

LABEL maintainer="Jefferson <jeffersonlacerda@live.com>"

WORKDIR /app

COPY go.mod . 

COPY go.sum .

RUN go mod download

COPY . .

ENV PORT=8000

RUN go build ./cmd/main.go

CMD [ "./main" ]