#/bin/sh

mkdir -p protos

protoc --go_out=protos --go_opt=paths=source_relative \
    --go-grpc_out=protos --go-grpc_opt=paths=source_relative \
    echo.proto
