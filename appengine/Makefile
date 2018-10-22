# ローカル環境で動かす
run:
	goapp serve ./app

# GAEにデプロイする
deploy:
	gcloud app deploy --project ${SANDBOX_PROJECT} --version 0 --quiet ./app/app.yaml