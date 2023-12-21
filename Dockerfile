FROM golang

COPY . /server
WORKDIR /server

RUN go build -v ./cmd/app

CMD [ "./app" ]