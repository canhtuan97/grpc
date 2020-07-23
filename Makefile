gen-cal:
	protoc calculatorpb/calculator.proto  --go_out=plugins=grpc:.
gen-app-server:
    protoc proto/app-server.proto --go_out=plugins=grpc:.
run-server:
	go run server/server.go