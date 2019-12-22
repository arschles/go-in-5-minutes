#!/bin/bash

# This file is deprecated and unused. It was needed when the GIFM
# site was hosted on Google Cloud Storage. It was since moved to 
# Netlify (https://netlify.com) 

gsutil rsync -R public gs://www.goin5minutes.com
gsutil -m acl ch -R -u AllUsers:R gs://www.goin5minutes.com
gsutil web set -m index.html -e 404.html gs://www.goin5minutes.com
gsutil web get gs://www.goin5minutes.com
