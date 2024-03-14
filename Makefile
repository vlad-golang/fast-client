mocks:
	mockery \
		--all --case=underscore --recursive \
		--with-expecter \
		--dir=. --output=mocks --keeptree

install:
	go install github.com/vektra/mockery/v2@v2.42.0
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.56.2
	go install github.com/onsi/ginkgo/v2/ginkgo

lint:
	golangci-lint run

tests:
	ginkgo ./...
