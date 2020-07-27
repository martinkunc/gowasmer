
build-wasm-1-13:
	docker run --rm -v "${PWD}":/go/src/github.com/martinkunc/gowasmer -w /go/src/github.com/martinkunc/gowasmer golang:1.13 bash -c "GOFLAGS=-mod=vendor GOOS=js GOARCH=wasm go build -o wasmrun.wasm wasmrun/*.go"
	test $$? -eq 0 && cp wasmrun.wasm cmd/

build-wasm-1-14:
	docker run --rm -v "${PWD}":/go/src/github.com/martinkunc/gowasmer -w /go/src/github.com/martinkunc/gowasmer golang:1.14 bash -c "GOFLAGS=-mod=vendor GOOS=js GOARCH=wasm go build -o wasmrun.wasm wasmrun/*.go"
	test $$? -eq 0 && cp wasmrun.wasm cmd/


build:
	go build -o wasmcmd ./cmd/main.go

