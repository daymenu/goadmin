run:
	go run ./cmd/admin
build:
	go build ./cmd/admin
protoc:
	build/package/protoc.sh
ent:
	go run ./internal/ent/gen