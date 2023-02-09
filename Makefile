
# Parameters
REGION=ap-southeast-2
PROFILE=coredev-dev
AWS=aws --profile $(PROFILE) --region $(REGION)
BUCKET_NAME=weather-event-sub
LOCATIONS=locations.json
CLIENTS=clients.json
SUBS=subscriptions.json

copy-locations:
	${AWS} s3 cp ./res/${LOCATIONS} s3://${BUCKET_NAME}/${LOCATIONS}

copy-clients:
	${AWS} s3 cp ./res/${CLIENTS} s3://${BUCKET_NAME}/${CLIENTS}

copy-subs:
	${AWS} s3 cp ./res/${SUBS} s3://${BUCKET_NAME}/${SUBS}

copy-all: copy-locations copy-clients copy-subs


