run:
	go run main.go

unit:
	go clean -testcache && go test -tags=unit -v ./...

unit-race:
	go test -race -tags=unit -v ./...

unit-coverage:
	go clean -testcache && go test -tags=unit -coverprofile=coverage.out ./... && go tool cover -html=coverage.out -o coverage.html && open coverage.html
