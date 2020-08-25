APP=http200

setup:
	go get -u github.com/securego/gosec/cmd/gosec
	go get github.com/golangci/golangci-lint/cmd/golangci-lint

build:
	go build -o bin/${APP} cmd/http200/main.go
	@echo "build complete"

clean:
	rm -rf .bin/${APP}
	@echo "cleaned"

fmt:
	goimports -w .

lint:
		golangci-lint run --enable-all --disable lll --disable misspell --disable gochecknoglobals --disable goconst

test:
	go test ./...

gosec:
	gosec -tests ./...

find-updates:
	go list -u -m -json all | go-mod-outdated -update -direct

.DEFAULT_GOAL := build
