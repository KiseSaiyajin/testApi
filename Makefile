build:
	go build -o bin/testApi cmd/main.go
test:
	go test -v ./...
run:build
	./bin/testApi