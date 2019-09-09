# ローカル環境で動かす
run:
	dev_appserver.py --enable_host_checking=False ./app/app.yaml

deploy_client:
	gcloud app deploy --project $(PROJECT_ID) --quiet ./client_app/app.yaml

deploy_server:
	gcloud app deploy --project $(PROJECT_ID) --quiet ./server_app/app.yaml

deploy_all: deploy_client deploy_server