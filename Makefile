clean:
	rm internal/channels/grpc/pb/*
	rm internal/swagger/*
gen:
	protoc --proto_path=tools/proto tools/proto/*.proto  --go_out=. --go-grpc_out=. --grpc-gateway_out=. --openapiv2_out=:internal/swagger
run:
	go run cmd/main.go
run-client:
	go run client/cmd/main.go