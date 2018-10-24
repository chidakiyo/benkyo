KEYCHAIN_NAME := benkyo-keychain
KEY_NAME := benkyo-key

# cloud build が cloud KMS を利用できるように IAM を追加する
apply-iam-policy/cloud-build-use-kms:
	gcloud --project ${PROJECT_ID} kms keys add-iam-policy-binding \
	$(KEY_NAME) --location=global --keyring=$(KEYCHAIN_NAME) \
	--member=serviceAccount:${SERVICE_ACCOUNT}@cloudbuild.gserviceaccount.com \
	--role=roles/cloudkms.cryptoKeyEncrypterDecrypter

# cat用のコンテナをgcrに登録しておく
docker-build/cat-container:
	cd cat && gcloud builds submit --project ${PROJECT_ID} --tag gcr.io/${PROJECT_ID}/cat

# Cloud Build 上で暗号化されているファイルを復号化する
submit/enc-text:
	gcloud --project ${PROJECT_ID} builds submit --config=cloudbuild_01.yaml ./

