GOOS=js GOARCH=wasm go build -o main.wasm -ldflags "-X main.CURRENT_USER=b"
