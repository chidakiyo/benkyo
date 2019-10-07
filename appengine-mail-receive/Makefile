# ローカル環境で動かす
run:
	dev_appserver.py --enable_host_checking=False ./app/app.yaml

deploy:
	gcloud app deploy --project $(PROJECT_ID) --quiet ./app/app.yaml

