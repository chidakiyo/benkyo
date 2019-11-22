
build-all: docker-build-v1 docker-build-v2
push-all: docker-push-v1 docker-push-v2
all: build-all push-all

docker-build-v1:
	docker build -t ${ECR_URI}:v1 -f ./container/v1.Dockerfile ./container

docker-build-v2:
	docker build -t ${ECR_URI}:v2 -f ./container/v2.Dockerfile ./container

docker-push-v1:
	docker push ${ECR_URI}:v1

docker-push-v2:
	docker push ${ECR_URI}:v2

ecr_login:
	aws ecr get-login --no-include-email --region us-east-2 > ./docker_login.sh
	chmod +x ./docker_login.sh

# --- debug ---

docker-run-v1:
	docker run --rm -itd -p 8080:8080 --name rolling_test ${ECR_URI}:v1

docker-run-v2:
	docker run --rm -itd -p 8080:8080 --name rolling_test ${ECR_URI}:v2

docker-stop:
	docker stop rolling_test

docker-clean:
	docker rm `docker ps -a -q`
