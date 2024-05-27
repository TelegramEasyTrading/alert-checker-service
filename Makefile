.PHONY: proto
proto:
   protoc --proto_path=proto --go_out=out --go_opt=paths=source_relative alert.proto \
run: 
   go run main.go