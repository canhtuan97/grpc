gen-cal:
	protoc calculatorpb/calculator.proto  --go_out=plugins=grpc:.
gen-app-server:
    protoc proto/app-server.proto --go_out=plugins=grpc:.
gen-gateway:
    protoc proto/DemoGateway.proto --go_out=plugins=grpc:. --grpc-gateway_out=:pb --swagger_out=:swagger
run-server:
	go run server/server.go