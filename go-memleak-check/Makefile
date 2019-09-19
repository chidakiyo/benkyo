
run:
	go run main.go

start-database:
	@echo "FIRESTORE_EMULATOR_HOST 環境変数を設定してください。"
	gcloud beta emulators datastore start --no-store-on-disk

env-init:
	$(gcloud beta emulators datastore env-init)