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
SAMPLE_ENC_TEXT := "CiQANnq0qL3tXKRtAPbBpysbluIbCWUSG31sdDccDhbmzIMnc30SkAEALx9DRaq1sr1tGrmmpzslkXEK0zLL3KaOwJj1dv459Xey2pRcpNGn3ugE8wRGKXmN0kNADZEDnSg2BJqY9LmdG08dY/MCJj4AwSMkyhY24yPnJQ9NCiPWWWXKy+FLxDW3NUyV66BaaMuy5Az1CWEp8tfYYALtvAZwZ9+VuR6W8v9O8nxPTWV7OIe0q1BjBWE="
dec/inline:
	echo -n $(SAMPLE_ENC_TEXT) | base64 -D | \
	gcloud --project ${PROJECT_ID} kms decrypt \
	--ciphertext-file=- \
	--plaintext-file=- \
	--location=global \
	--keyring=$(KEYCHAIN_NAME) \
	--key=$(KEY_NAME)

