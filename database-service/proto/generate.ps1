# Generate protobuf Go files
# Requires: protoc in PATH (download from https://github.com/protocolbuffers/protobuf/releases)
# Run from database-service folder: .\proto\generate.ps1

protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/database.proto
Write-Host "Proto generated: proto/*.pb.go" -ForegroundColor Green
