sqlc:
	docker run --rm -v "D:\_PROJECTS\_GIT\GoMockProject\Source\go-mock-project:/src" -w /src kjconroy/sqlc generate

proto:
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
    --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
    proto/*.proto

flight-grpc:
	go run grpc/flight-grpc/main.go
booking-grpc:
	go run grpc/booking-grpc/main.go
user-grpc:
	go run grpc/user-grpc/main.go

flight-api:
	go run api/flight-api/main.go
booking-api:
	go run api/booking-api/main.go
customer-api:
	go run api/customer-api/main.go

.PHONY: sqlc proto flight-grpc booking-grpc user-grpc flight-api booking-api customer-api