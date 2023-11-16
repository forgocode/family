.PHONY:
prepare:
	@echo "\n\n\n" >> /etc/hosts
	@echo "192.168.0.202 mysql.test.com" >> /etc/hosts
	@echo "192.168.0.202 redis.test.com" >> /etc/hosts
	@echo "192.168.0.202 mongo.test.com" >> /etc/hosts
	
	@echo "10.182.34.112 mysql.test.com" >> /etc/hosts
	@echo "10.182.34.112 redis.test.com" >> /etc/hosts
	@echo "10.182.34.112 mongo.test.com" >> /etc/hosts

	@echo "system prepare ready"

	@cp ./internal/conf/config.yaml /root/tmp/.config.yaml

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