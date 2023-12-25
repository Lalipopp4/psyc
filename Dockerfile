FROM golang:latest-alpine

COPY . /app
WORKDIR /app

RUN go build -v -o app ./cmd/app

CMD [ "./app" ]