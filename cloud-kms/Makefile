KEYCHAIN_NAME := benkyo-keychain
KEY_NAME := benkyo-key

# keychainを作成する
create/keychain:
	gcloud --project ${PROJECT_ID} kms keyrings create $(KEYCHAIN_NAME) --location global

# keyを作成する
create/key:
	gcloud --project ${PROJECT_ID} kms keys create $(KEY_NAME) --location global --keyring $(KEYCHAIN_NAME) --purpose encryption

# 作成したkeyを確認する
check/keys:
	gcloud --project ${PROJECT_ID} kms keys list --location global --keyring $(KEYCHAIN_NAME)

# base.txt を 暗号化した base.text.enc に変換する
enc/text:
	gcloud --project ${PROJECT_ID} kms encrypt \
	--plaintext-file=base.txt \
	--ciphertext-file=base.txt.enc \
	--location=global \
	--keyring=$(KEYCHAIN_NAME) \
	--key=$(KEY_NAME)

# 暗号化された base.text.enc をもとに戻して base.text.dec として出力する
dec/text:
	gcloud --project ${PROJECT_ID} kms decrypt \
	--ciphertext-file=base.txt.enc \
	--plaintext-file=base.txt.dec \
	--location=global \
	--keyring=$(KEYCHAIN_NAME) \
	--key=$(KEY_NAME)

# 変数を直接暗号化して出力する
SAMPLE_TEXT := "以下、話は、東京中心であるから、そのつもりで、きいていただきたい。"
enc/inline:
	echo -n $(SAMPLE_TEXT) | \
	gcloud --project ${PROJECT_ID} kms encrypt \
	--plaintext-file=- \
	--ciphertext-file=- \
	--location=global \
	--keyring=$(KEYCHAIN_NAME) \
	--key=$(KEY_NAME) | base64

# encodeされている変数を直接復号化して出力する（うまくいかない）
SAMPLE_ENC_TEXT := "CiQANnq0qOdjCv1xSfIDJiEuR+mhMHDknH0XVc37iWQmgH1IlXsSjQEALx9DRYZExBoaFpX03Zpd9vSyUv+3AylvqONEfBiP9CWRegWQzob6IZx6zLhqM+Y6ooPWpZ6fXHehMbdWR6qoSZBpaCsjNNeMzQHXPi0SSimGtB184IQNdf0f1dRC9R62gess+BKaYPE3LcE8cvLeVhFVHHp/BXM0Au5rh34jv/okmd2MGhySGbcOqhM="
dec/inline:
	echo -n $(SAMPLE_ENC_TEXT) | base64 -D | \
	gcloud --project ${PROJECT_ID} kms decrypt \
	--ciphertext-file=- \
	--plaintext-file=- \
	--location=global \
	--keyring=$(KEYCHAIN_NAME) \
	--key=$(KEY_NAME)

