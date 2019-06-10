build:
	GOOS=linux GOARCH=amd64 go build
	docker build -t creatuture/consignment-cli:0.1 .
run:
	docker run -e MICRO_REGISTRY=mdns creatuture/consignment-cli:0.1