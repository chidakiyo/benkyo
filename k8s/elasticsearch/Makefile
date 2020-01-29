# run at k8s docker for mac.

00-apply-crd:
	kubectl apply -f https://download.elastic.co/downloads/eck/1.0.0-beta1/all-in-one.yaml

01-apply-namespace:
	kubectl apply -f namespace.yaml

02-apply-elasticsearch:
	kubectl apply -f elasticsearch.yaml

03-password:
	PASSWORD=$(kubectl get secret elasticsearch-es-elastic-user -o=jsonpath='{.data.elastic}' -n elastic | base64 --decode)

04-curl_test:
	curl -u "elastic:$PASSWORD" http://localhost:9200

05-apply-kibana:
	kubectl apply -f kibana.yaml
