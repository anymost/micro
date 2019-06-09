build:
	protoc -I . ./proto/consignment/consignment.proto --go_out=plugins=grpc:./
	GOOS=linux GOARCH=amd64 go build
	docker build -t creatuture/consignment-service:0.1 .
run:
	docker run -d -p 9090:9090 creatuture/consignment-service:0.1