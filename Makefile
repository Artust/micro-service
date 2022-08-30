.PHONY: generate_proto_gateway
generate_proto_gateway:
	protoc --go_out=./services/gateway/protos/gateway --go_opt=paths=source_relative \
	--go-grpc_out=./services/gateway/protos/gateway --go-grpc_opt=paths=source_relative \
	--proto_path=./api/grpc/ avatar.proto

	protoc --go_out=./services/gateway/protos/center --go_opt=paths=source_relative \
	--go-grpc_out=./services/gateway/protos/center --go-grpc_opt=paths=source_relative \
	--proto_path=./internal/api/grpc/ center.proto

	protoc --go_out=./services/gateway/protos/corporation --go_opt=paths=source_relative \
	--go-grpc_out=./services/gateway/protos/corporation --go-grpc_opt=paths=source_relative \
	--proto_path=./internal/api/grpc/ corporation.proto

	protoc --go_out=./services/gateway/protos/pos --go_opt=paths=source_relative \
	--go-grpc_out=./services/gateway/protos/pos --go-grpc_opt=paths=source_relative \
	--proto_path=./internal/api/grpc/ pos.proto

	protoc --go_out=./services/gateway/protos/streaming --go_opt=paths=source_relative \
	--go-grpc_out=./services/gateway/protos/streaming --go-grpc_opt=paths=source_relative \
	--proto_path=./internal/api/grpc/ streaming.proto

	protoc --go_out=./services/gateway/protos/talk_session --go_opt=paths=source_relative \
	--go-grpc_out=./services/gateway/protos/talk_session --go-grpc_opt=paths=source_relative \
	--proto_path=./internal/api/grpc/ talk_session.proto

	protoc --go_out=./services/gateway/protos/account_management --go_opt=paths=source_relative \
	--go-grpc_out=./services/gateway/protos/account_management --go-grpc_opt=paths=source_relative \
	--proto_path=./internal/api/grpc/ account_management.proto
	
.PHONY: generate_proto_account account_management
generate_proto_account:
	protoc --go_out=./services/account_management/ \
    --go-grpc_out=./services/account_management/protos \
	--go-grpc_opt=paths=source_relative \
    --proto_path=./internal/api/grpc/ account_management.proto

.PHONY: generate_proto_pos
generate_proto_pos:
	protoc --go_out=./services/pos/ \
    --go-grpc_out=./services/pos/protos \
	--go-grpc_opt=paths=source_relative \
    --proto_path=./internal/api/grpc/ pos.proto

.PHONY: generate_proto_corporation
generate_proto_corporation:
	protoc --go_out=./services/corporation/ \
    --go-grpc_out=./services/corporation/protos \
	--go-grpc_opt=paths=source_relative \
    --proto_path=./internal/api/grpc/ corporation.proto

.PHONY: generate_proto_streaming
generate_proto_streaming:
	protoc --go_out=./services/streaming/protos --go_opt=paths=source_relative \
	--go-grpc_out=./services/streaming/protos --go-grpc_opt=paths=source_relative \
	--proto_path=./internal/api/grpc/ streaming.proto

.PHONY: generate_proto_center
generate_proto_center:
	protoc --go_out=./services/center/protos --go_opt=paths=source_relative \
	--go-grpc_out=./services/center/protos --go-grpc_opt=paths=source_relative \
	--proto_path=./internal/api/grpc/ center.proto

.PHONY: generate_proto_talk_session
generate_proto_talk_session:
	protoc --go_out=./services/talk_session/protos --go_opt=paths=source_relative \
	--go-grpc_out=./services/talk_session/protos --go-grpc_opt=paths=source_relative \
	--proto_path=./internal/api/grpc/ talk_session.proto

.PHONY: generate_proto_all
generate_proto_all: generate_proto_streaming \
										generate_proto_gateway \
										generate_proto_pos \
										generate_proto_corporation \
										generate_proto_center \
										generate_proto_talk_session \
										generate_proto_account \

run_account_management:
	go run services/account_management/cmd/main.go

run_pos:
	go run services/pos/cmd/main.go
	
run_center:
	go run services/center/cmd/main.go

run_talk_session:
	go run services/talk_session/cmd/main.go

run_gateway: 
	go run services/gateway/cmd/main.go

run_streaming:
	go run services/streaming/cmd/main.go

run_corporation:
	go run services/corporation/cmd/main.go

run_upload:
	go run services/upload/cmd/main.go
	
# make run_all -j8 // Run all service in one command
run_all:run_pos \
				run_center \
				run_talk_session \
				run_streaming \
				run_corporation \
				run_upload \
				run_account_management \
				run_gateway \
				 
seed:
	docker exec -d gateway ./seed
