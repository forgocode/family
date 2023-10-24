.PHONY:
prepare:
	@echo "\n\n\n" >> /etc/hosts
	@echo "192.168.0.202 mysql.test.com" >> /etc/hosts
	@echo "192.168.0.202 redis.test.com" >> /etc/hosts
	@echo "192.168.0.202 mongo.test.com" >> /etc/hosts
	@echo "system prepare ready"

.PHONY:
loggrpc:
	@protoc --go_out=./internal/grpcserver/proto/log/ ./internal/grpcserver/proto/log/log.proto
	@protoc --go-grpc_out=./internal/grpcserver/proto/log/ ./internal/grpcserver/proto/log/log.proto
	@echo "log grpc protobuf generate successfully"
