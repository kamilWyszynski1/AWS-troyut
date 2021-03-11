build:
	mkdir -p /bin
	go build -o /bin/main cmd/main.go

clean:
	rm -rf /bin

test:
	go test -v ./...

run:
	make build
	go run ../bin/main