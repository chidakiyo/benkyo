run:
	go run app/main.go

config/build: config/clean
	$(shell docker run --rm -it -e SERVICE=default -v $(CURDIR):/tmp supinf/envsubst /tmp/app/app_template.yaml /tmp/app/env.txt >| app/app.yaml)

config/clean:
	rm -f app/app.yaml

deploy: config/build
	gcloud app deploy --project $(PROJECT_ID) --quiet app/app.yaml
	make config/clean

