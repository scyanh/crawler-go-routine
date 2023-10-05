test:
	go test ./...

test-race:
	go test -race ./...

coverage:
	go test -cover ./...

coverage-func:
	go test -coverprofile=coverage.out ./... ;    go tool cover -func=coverage.out

coverage-html:
	go test -coverprofile=coverage.out ./... ;    go tool cover -html=coverage.out

mockgen:
	mockgen -source=pkg/domain/interfaces/http_client.go -destination=pkg/domain/interfaces/http_client_mock.go -package=interfaces
	mockgen -source=pkg/domain/interfaces/repository.go -destination=pkg/domain/interfaces/repository_mock.go -package=interfaces