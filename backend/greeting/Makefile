postgres:
	podman run --name greet -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:15-alpine
createdb:
	 podman exec -it greet createdb --username=root --owner=root greet
dropdb:
	podman exec -it greet dropdb kiteapi

.Phony:
	postgres createdb dropdb