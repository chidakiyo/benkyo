run:
	go run app/main.go

deploy/gae:
	gcloud app deploy --project $(PROJECT_ID) --quiet app/app.yaml

deploy/run:
	gcloud builds submit --project $(PROJECT_ID) --config cloudbuild.yaml .

