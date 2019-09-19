
run:
	go run ./...

start-database:
	@echo "FIRESTORE_EMULATOR_HOST 環境変数を設定してください。"
	gcloud beta emulators datastore start --no-store-on-disk

env-init:
	$(gcloud beta emulators datastore env-init)

run-gae:
	dev_appserver.py --enable_host_checking=False ./app.yaml