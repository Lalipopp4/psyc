FROM golang:alpine

COPY . /app
WORKDIR /app

RUN go build -v -o app ./cmd/app

EXPOSE 8000

CMD [ "./app" ]