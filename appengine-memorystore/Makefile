enable-vpc-api:
	gcloud services enable vpcaccess.googleapis.com

create-serverless-vpc:
	gcloud beta compute networks vpc-access connectors create ${CONNECTOR_NAME} \
    --network ${VPC_NETWORK} \
    --region ${REGION} \
    --range ${IP_RANGE}

deploy:
	gcloud app deploy --project $(PROJECT_ID) --quiet ./app/app.yaml


