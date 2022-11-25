PORT := 8000
IP := localhost

build: assets/index.wasm

assets/index.wasm: $(wildcard cmd/client/*.go)
	cd cmd/client; GOOS=js GOARCH=wasm go build -o ../../assets/index.wasm -ldflags "-X main.PORT=$(PORT) -X main.IP=$(IP)"

run: build
	cd cmd/server; PORT=8000 go run main.go