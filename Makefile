API_PROTO_FILES=$(shell find api -name *.proto)

.PHONY: proto
proto:
	protoc --proto_path=. \
           --proto_path=./pkg \
           --go_out=. --go_opt=paths=source_relative \
           --go-grpc_out=. --go-grpc_opt=paths=source_relative \
           --go-errors_out=paths=source_relative:. \
           $(API_PROTO_FILES)

.PHONY: grpc
# generate grpc code
grpc:
	protoc --proto_path=. \
           --proto_path=./third_party \
           --go_out=paths=source_relative:. \
           --go-grpc_out=paths=source_relative:. \
           $(API_PROTO_FILES)

.PHONY: http
# generate http code
http:
	protoc --proto_path=. \
           --proto_path=./third_party \
           --go_out=paths=source_relative:. \
           --go-http_out=paths=source_relative:. \
           $(API_PROTO_FILES)

.PHONY: swagger
# generate swagger file
swagger:
	protoc --proto_path=. \
		--proto_path=./third_party \
		--openapiv2_out . \
		--openapiv2_opt logtostderr=true \
		$(API_PROTO_FILES)


.PHONY: build
build:
	docker build -t app .
	docker-compose up -d