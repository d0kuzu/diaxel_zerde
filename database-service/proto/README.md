# Protobuf generation

1) Install protoc for Windows
   - Download from https://github.com/protocolbuffers/protobuf/releases (win64 zip)
   - Extract and add `bin` to your PATH

2) Install Go plugins (already done):
   ```sh
   go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
   go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
   ```

3) Generate Go files
   ```sh
   protoc --go_out=. --go-grpc_out=. proto/database.proto
   ```
   Or run from database-service folder:
   ```sh
   .\proto\generate.ps1
   ```

After generation you should see:
- `proto/database.pb.go`
- `proto/database_grpc.pb.go`

Then run `go mod tidy` to resolve imports.
