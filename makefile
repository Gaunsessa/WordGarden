build: assets/index.wasm

assets/index.wasm: $(wildcard cmd/client/*.go)
	cd cmd/client; GOOS=js GOARCH=wasm go build -o ../../assets/index.wasm

run: build
	cd cmd/server; go run main.go