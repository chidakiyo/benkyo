# デプロイする
gae/deploy:
	GO111MODULE=on go mod vendor
	GO111MODULE=off gcloud app deploy --project ${PROJECT_ID} --no-promote --quiet ./app/app.yaml
	rm -rf ./vendor