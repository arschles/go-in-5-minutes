#!/bin/bash

JWT_FILE="gifm-jwt.json"
touch $JWT_FILE
JWT=$(docker run --rm -e GCSUP_JWT=$GCSUP_JWT -v $PWD:/pwd -w /pwd alpine:3.3 echo "$GCSUP_JWT" | base64 -D)
echo "$JWT" > $JWT_FILE

# note the need for GOOGLE_APPLICATION_CREDENTIALS. it's needed by gcsup.
# see https://github.com/arschles/gcsup/issues/4
docker run --rm -v $PWD:/pwd -w /pwd \
  -e GCSUP_JWT_FILE_LOCATION=/pwd/$JWT_FILE \
  -e GCSUP_PROJECT_NAME=$GCSUP_PROJECT_NAME \
  -e GCSUP_BUCKET_NAME=$GCSUP_BUCKET_NAME \
  -e GCSUP_LOCAL_FOLDER=www/public \
  -e GOOGLE_APPLICATION_CREDENTIALS=/pwd/$JWT_FILE \
  quay.io/arschles/gcsup:latest gcsup
