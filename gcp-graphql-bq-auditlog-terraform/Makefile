

tools:
	@go build -mod=mod -o ./_bin/gqlgen		github.com/99designs/gqlgen

gqlgen-init:
	./_bin/gqlgen init

gqlgen-basic:
	./_bin/gqlgen generate

# add resolver.go -> //go:generate go run github.com/99designs/gqlgen
gqlgen:
	go generate ./...

.PHONY: build-cb
build-cb:
	gcloud --project ${PROJECT_ID} alpha builds submit \
	--pack=builder=gcr.io/buildpacks/builder,env=GOOGLE_BUILDABLE=./app,image=gcr.io/${PROJECT_ID}/app

# 一撃デプロイコマンド
deploy:
	gcloud --project ${PROJECT_ID} beta run deploy \
	app \
	--region asia-northeast1 \
	--platform managed \
	--allow-unauthenticated \
	--set-env-vars GCP_PROJECT=${PROJECT_ID} \
	--max-instances 1 \
	--source .