all:
	go build -o start.exe ./cmd/app
	./start

run:
	./start