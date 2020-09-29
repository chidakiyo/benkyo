

tools:
	@go build -mod=mod -o ./_bin/gqlgen		github.com/99designs/gqlgen

gqlgen-init:
	./_bin/gqlgen init

gqlgen-basic:
	./_bin/gqlgen generate

# add resolver.go -> //go:generate go run github.com/99designs/gqlgen
gqlgen:
	go generate ./...
