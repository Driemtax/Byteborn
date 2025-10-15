.DEFAULT_GOAL := start

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: vet
vet: fmt
	go vet ./...

.PHONY: test
test:
	go test ./...

.PHONY: build
build: vet
	go build -o byteborn-game ./cmd/byteborn/main.go

.PHONY: start
start: build
	./byteborn-game

commit: test
	git commit

.PHONY: clean
clean:
	rm -rf ./byteborn-game
