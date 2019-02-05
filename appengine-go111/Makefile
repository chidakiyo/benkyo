PROJ=chida-test-012

# ローカル環境で動かす
run:
	go run ./app/main.go # TODO

# GAEにデプロイする
deploy:
	gcloud app deploy --project $(PROJ) --version $(shell env TZ=Asia/Tokyo date "+%Y%m%d-%H%M%S")-$(shell git show -s --format=%H) --quiet ./app/app.yaml

# パフォーマンス用にデプロイする
deploy-perf:
	gcloud app deploy --project ${GAE_PROJECT} --version $(shell env TZ=Asia/Tokyo date "+%Y%m%d-%H%M%S")-$(shell git show -s --format=%H) --quiet ./app/app.yaml
