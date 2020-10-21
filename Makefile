lint:
	go fmt ./...
	go vet ./...
	go get golang.org/x/lint/golint	
	golint -set_exit_status $(go list ./...)
	go get github.com/golangci/golangci-lint/cmd/golangci-lint
	golangci-lint run -E gofmt -E golint -E vet 
