compile:
	GOOS=js GOARCH=wasm go build -o  ./assets/json.wasm
	cp "$(shell go env GOROOT)/misc/wasm/wasm_exec.js" ./assets

serve:
	go run cmd/server.go
