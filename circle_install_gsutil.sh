curl https://storage.googleapis.com/pub/gsutil.tar.gz > gsutil.tar.gz
mkdir -p $CIRCLE_BUILD_DIR/bin/gsutil
tar xfz gsutil.tar.gz -C $CIRCLE_BUILD_DIR/bin/gsutil

CFG=$'[Credentials]
gs_access_key_id = $GCS_ACCESS_KEY_ID
gs_secret_access_key = $GCS_SECRET_ACCESS_KEY
[Boto]
https_validate_certificates = True
[GSUtil]
content_language = en
default_api_version = 2
default_project_id = go-in-5-minutes'

echo "$CFG" > $BOTO_CONFIG
