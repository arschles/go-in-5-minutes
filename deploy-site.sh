#!/bin/bash

JWT_FILE="gifm-jwt.json"
touch $JWT_FILE
echo $GCSUP_JWT_FILE > $JWT_FILE

docker run --rm -v $PWD:/pwd -w /pwd -e GCSUP_JWT_FILE_LOCATION=$JWT_FILE -e GCSUP_PROJECT_NAME=$GCS_PROJECT_NAME -e quay.io/arschles/gcsup:0.0.1 gcsup
