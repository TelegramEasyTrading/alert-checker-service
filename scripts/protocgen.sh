#!bin/sh

# Generate the protobuf files
protoc --proto_path=proto --go_out=internal/model --go_opt=paths=source_relative alert.proto