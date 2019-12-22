#!/bin/sh

set -e

# This script will build all of the elm files
# needed to power the static site.
#
# Taken from
# https://guide.elm-lang.org/optimization/asset_size.html#scripts

# js="elm.js"
# min="elm.min.js"

# elm make --optimize --output=$js $@
elm make elm/Screencasts.elm --output=static/js/screencasts.elm.js
# elm make elm/Screencast.Elm --output = static/js/screencast.elm.js

# uglifyjs $js --compress 'pure_funcs="F2,F3,F4,F5,F6,F7,F8,F9,A2,A3,A4,A5,A6,A7,A8,A9",pure_getters,keep_fargs=false,unsafe_comps,unsafe' | uglifyjs --mangle --output=$min

# echo "Compiled size:$(cat $js | wc -c) bytes  ($js)"
# echo "Minified size:$(cat $min | wc -c) bytes  ($min)"
# echo "Gzipped size: $(cat $min | gzip -c | wc -c) bytes"
