run:
	go run main.go

unit:
	go clean -testcache && go test -tags=unit -v ./...
