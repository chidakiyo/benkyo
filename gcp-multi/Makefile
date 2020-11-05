
.PHONY: deploy
deploy:
	gcloud --project ${PROJECT_ID} app deploy --quiet app_a.yaml

.PHONY: build-cb
build-cb:
	gcloud --project ${PROJECT_ID} alpha builds submit --pack=builder=gcr.io/buildpacks/builder,env=GOOGLE_BUILDABLE=./app,image=gcr.io/${PROJECT_ID}/env-container