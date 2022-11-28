WSIP := ws://localhost:8000/ws

build: assets/index.wasm

assets/index.wasm: $(wildcard cmd/client/*.go)
	cd cmd/client; GOOS=js GOARCH=wasm go build -o ../../assets/index.wasm -ldflags "-X main.WSIP=$(WSIP)"

run: build
	cd cmd/server; PORT=8000 go run main.go