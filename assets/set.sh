cd ../src

export GOOS=js
export GOARCH=wasm
go build -o ../assets/main.wasm