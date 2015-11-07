#!/bin/bash

gsutil rsync -R public gs://www.goin5minutes.com
gsutil -m acl ch -R -u AllUsers:R gs://www.goin5minutes.com
gsutil web set -m index.html -e 404.html gs://www.goin5minutes.com
gsutil web get gs://www.goin5minutes.com
