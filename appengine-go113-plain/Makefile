run:
	go run app/main.go

config/build/a: config/clean
	$(shell docker run --rm -it -e SERVICE=svc-a -v $(CURDIR):/tmp supinf/envsubst /tmp/app/app_template.yaml /tmp/app/env.txt >| app/app.yaml)

config/build/b: config/clean
	$(shell docker run --rm -it -e SERVICE=svc-b -v $(CURDIR):/tmp supinf/envsubst /tmp/app/app_template.yaml /tmp/app/env.txt >| app/app.yaml)

config/build/default: config/clean
	$(shell docker run --rm -it -e SERVICE=default -v $(CURDIR):/tmp supinf/envsubst /tmp/app/app_template.yaml /tmp/app/env.txt >| app/app.yaml)

config/clean:
	rm -f app/app.yaml

deploy/default: config/build/default
	gcloud app deploy --project $(PROJECT_ID) --quiet app/app.yaml
	make config/clean

deploy/a: config/build/a
	gcloud app deploy --project $(PROJECT_ID) --quiet app/app.yaml
	make config/clean

deploy/b: config/build/b
	gcloud app deploy --project $(PROJECT_ID) --quiet app/app.yaml
	make config/clean

deploy/all: deploy/a deploy/b
