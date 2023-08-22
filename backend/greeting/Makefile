proto:
	protoc -I grpc/proto/ --go_opt=module=greeting --go_out=. --go-grpc_opt=module=greeting --go-grpc_out=. grpc/proto/*.proto
gqlgen:
	go mod tidy
	go get github.com/99designs/gqlgen
	go run github.com/99designs/gqlgen generate --config gqlgen.yml
postgres:
	podman run --name workshop -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:15-alpine
createdb:
	 podman exec -it workshop createdb --username=root --owner=root greet
dropdb:
	podman exec -it greet dropdb kiteapi
.PHONEY:
	proto createdb dropdb gqlgen