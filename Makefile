test:
	go mod tidy
	go test -v ./...

install:
	go install

lint:
	golangci-lint run ./...