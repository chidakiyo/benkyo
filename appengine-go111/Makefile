PROJ=chida-test-012

# ローカル環境で動かす
run:
	dev_appserver.py --enable_host_checking=False ./app/app.yaml

# GAEにデプロイする
deploy:
	gcloud app deploy --project $(PROJ) --version $(shell env TZ=Asia/Tokyo date "+%Y%m%d-%H%M%S")-$(shell git show -s --format=%H) --quiet ./app/app.yaml

# パフォーマンス用にデプロイする
deploy-perf:
	gcloud app deploy --project ${GAE_PROJECT} --version $(shell env TZ=Asia/Tokyo date "+%Y%m%d-%H%M%S")-$(shell git show -s --format=%H) --quiet ./app/app.yaml
