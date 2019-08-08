#!/bin/sh

set -eou pipefail

docker run -p 3000:3000 -e GO_ENV=development -e ATHENS_GO_GET_WORKERS=5 gomods/athens:v0.5.0
