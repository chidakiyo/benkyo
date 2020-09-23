.PHONY: tools
tools:
	@go build -mod=mod -o ./_bin/wire github.com/google/wire/cmd/wire

.PHONY: wire-gen
wire-gen:
	./_bin/wire gen