protoc -I . --go_out . --go_opt paths=source_relative api/*.proto
# --go-grpc_out gen --go-grpc_opt paths=source_relative \
