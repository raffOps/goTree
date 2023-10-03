install:
	go install honnef.co/go/tools/cmd/staticcheck@latest
	go install golang.org/x/lint/golint@latest
	go install github.com/kisielk/errcheck@latest

pre-commit:
	pre-commit run --all-files

lint:
	go vet ./...
	golint ./...
	staticcheck ./...
	errcheck ./...

test:
	go test ./... -v

coverage:
	go test ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out
