# yaml から chi-server の go コードを生成する
gen/gocode:
	oapi-codegen -package api -old-config-style -generate "chi-server","types","spec" openapi.yaml > api/api.gen.go # default : types,client,server,spec

# サーバを起動する
run:
	go run main.go
# airで起動する
run/air:
	air

# oapi-codegen を install する
# 利用するには ${GOPATH}/bin にパスを通している必要があります
init/install/oapi:
	go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest

# air を install する
init/install/air:
	go install github.com/cosmtrek/air@latest

# air の初期設定（ファイル生成）する
init/gen/air:
	air init