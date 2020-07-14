.PHONY: proto
proto:
	protoc -I . --go_out=. --micro_out=. proto/vessel/vessel.proto

.PHONY: vendor
vendor:
	go mod vendor

.PHONY: runv
runv: vendor
	go run main.go

.PHONY: run
run:
	go run main.go

.PHONY: build
build:
	go build -o vessel *.go

.PHONY: clean
clean:
	rm -f vessel

.PHONY: test
test:
	go test -v ./... -cover
