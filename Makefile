test:
	go test -cover -v ./...

cover:
	go test -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out
