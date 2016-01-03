#!/bin/bash

JWT_FILE="gifm-jwt.json"
touch $JWT_FILE
JWT=$(echo "$GCSUP_JWT" | base64 -D)
echo "$JWT" > $JWT_FILE

# note the need for GOOGLE_APPLICATION_CREDENTIALS. it's needed by gcsup.
# see https://github.com/arschles/gcsup/issues/4
docker run --rm -v $PWD:/pwd -w /pwd \
  -e GCSUP_JWT_FILE_LOCATION=/pwd/$JWT_FILE \
  -e GCSUP_PROJECT_NAME=go-in-5-minutes \
  -e GCSUP_BUCKET_NAME=goin5minutes-site-test \
  -e GCSUP_LOCAL_FOLDER=www/public \
  -e GOOGLE_APPLICATION_CREDENTIALS=/pwd/$JWT_FILE \
  quay.io/arschles/gcsup:latest gcsup
