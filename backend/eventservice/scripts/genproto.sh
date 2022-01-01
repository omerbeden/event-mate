
protodir=../../../proto/grpc/event/v1

protoc --go_out=plugins=grpc:genproto -I $protodir $protodir/event-service.proto
protoc --go_out=plugins=grpc:genproto -I proto/grpc/event/v1/ proto/grpc/event/v1/event-service.proto