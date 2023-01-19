API_PROTO_FILES=$(shell find api -name *.proto)

.PHONY: proto
proto:
	protoc --proto_path=. \
           --proto_path=./third_party \
           --go_out=paths=source_relative:. \
           --go-grpc_out=paths=source_relative:. \
           --gin-http_out=paths=source_relative:. \
           --go-errors_out=paths=source_relative:. \
           --validate_out=paths=source_relative,lang=go:. \
           --openapiv2_out . \
           --openapiv2_opt logtostderr=true \
           $(API_PROTO_FILES)


.PHONY: build
build:
	docker build -t app .
	docker-compose up -d