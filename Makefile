lint: load-lint run-lint
test: rundb unit integration
#before pr launched
pr: format tidy lint rundb test

run-lint:
	.bin/golangci-lint run --config golangci.yml

load-lint:
	mkdir -p .bin
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b .bin $(golangci_lint_version)

format:
	goimports -w .

tidy:
	go mod tidy

rundb:
	docker-compose up database redis -d

run:
	docker-compose up -d

down:
	docker-compose down

unit:
	gotest -v -tags unit ./...

integration:
	gotest -v -tags integration ./...

cover:
	go test -tags unit,integration -covermode=atomic -coverpkg=./... -coverprofile coverage.out ./... && \
	go tool cover -func=coverage.out && \
	go tool cover -html=coverage.out