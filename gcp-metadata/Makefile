# ローカル環境で動かす
run:
	go run ./app/main.go # TODO

# GAEにデプロイする
deploy:
	gcloud app deploy --project $(PROJECT_ID) --quiet ./app/app.yaml

# Runにデプロイする
run/build:
	GO111MODULE=on gcloud builds submit --project ${PROJECT_ID} --tag gcr.io/${PROJECT_ID}/meta

run/deploy:
	GO111MODULE=on gcloud beta run deploy --project ${PROJECT_ID} --image gcr.io/${PROJECT_ID}/meta --region us-central1 meta --quiet --allow-unauthenticated

run/build-and-deploy: run/build run/deploy