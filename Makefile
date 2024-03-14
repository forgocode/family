.PHONY:
prepare:
	@echo "\n\n\n" >> /etc/hosts
	@echo "192.168.0.200 mysql.test.com" >> /etc/hosts
	@echo "192.168.0.200 redis.test.com" >> /etc/hosts
	@echo "192.168.0.200 mongo.test.com" >> /etc/hosts
	
	@echo "10.182.34.112 mysql.test.com" >> /etc/hosts
	@echo "10.182.34.112 redis.test.com" >> /etc/hosts
	@echo "10.182.34.112 mongo.test.com" >> /etc/hosts
	@echo "system prepare ready"
	@cp ./internal/conf/config.yaml /root/tmp/.config.yaml
	@echo "prepare db env"
	@bash ./scripts/prepare.sh

.PHONY:
messagegrpc:
	@protoc --go_out=./internal/grpcserver/proto/message/ ./internal/grpcserver/proto/message/message.proto
	@protoc --go-grpc_out=./internal/grpcserver/proto/message/ ./internal/grpcserver/proto/message/message.proto
	@echo "message grpc protobuf generate successfully"

.PHONY:
loggrpc:
	@protoc --go_out=./internal/grpcserver/proto/log/ ./internal/grpcserver/proto/log/log.proto
	@protoc --go-grpc_out=./internal/grpcserver/proto/log/ ./internal/grpcserver/proto/log/log.proto
	@echo "log grpc protobuf generate successfully"

.PHONY:
noticegrpc:
	@protoc --go_out=./internal/grpcserver/proto/station_notice/ ./internal/grpcserver/proto/station_notice/station_notice.proto
	@protoc --go-grpc_out=./internal/grpcserver/proto/station_notice/ ./internal/grpcserver/proto/station_notice/station_notice.proto
	@echo "station_notice grpc protobuf generate successfully"
	
.PHONY:
build:
	@rm -rf bin
	@mkdir bin
	@go build -o bin/family_webservice cmd/webservice/main.go
	@echo "build family_webservice successfully!"
	@./bin/family_webservice


