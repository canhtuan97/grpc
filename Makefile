gen-cal:
    protoc proto/*.proto --go_out=paths=source_relative,plugins=grpc:. --grpc-gateway_out=logtostderr=true,paths=source_relative:.
run-server:
	go run server/server.go