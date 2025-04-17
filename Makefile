build:
	tinygo build -target=wasm-unknown -o spacetime.wasm ./bindings.go
test:
	tinygo test ./...
