#!/bin/bash

JWT_FILE="gifm-jwt.json"
touch $JWT_FILE
JWT=$(echo "$GCSUP_JWT" | base64 -D)
echo "$JWT" > $JWT_FILE

# build in hugo
docker run --rm -v "$PWD/www":/www -w /www quay.io/arschles/hugo:latest hugo -v

docker run --rm -v $PWD:/pwd -w /pwd \
  -e GCSUP_JWT_FILE_LOCATION=$JWT_FILE \
  -e GCSUP_PROJECT_NAME=$GCS_PROJECT_NAME \
  -e GCSUP_BUCKET_NAME=$GCSUP_BUCKET_NAME \
  -e GCSUP_LOCAL_FOLDER=www/public quay.io/arschles/gcsup:0.0.1 gcsup
