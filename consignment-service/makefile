build:
	protoc -I . ./proto/consignment/consignment.proto --go_out=plugins=micro:./
	GOOS=linux GOARCH=amd64 go build
	docker build -t creatuture/consignment-service:0.1 .
run:
	docker run -p 9090:9090 -e MICRO_SERVER_ADDRESS=:9090 -e MICRO_REGISTRY=mdns creatuture/consignment-service:0.1