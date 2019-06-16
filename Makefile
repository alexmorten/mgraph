gen:
	protoc --go_out=plugins=grpc:. proto/db.proto
	protoc --go_out=plugins=grpc:. proto/request.proto
