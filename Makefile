.PHONY: proto
proto:
   protoc --proto_path=proto --go_out=internal/model --go_opt=paths=source_relative alert.proto \
run: 
   go run main.go