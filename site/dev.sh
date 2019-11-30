#!/bin/bash

# This script runs commands in parallel, and stops them all when 
# a kill signal (ctrl+c) is pressed
#
# I took it from
# https://stackoverflow.com/a/10909842

# for cmd in "$@"; do {
#   echo "Process \"$cmd\" started";
#   $cmd & pid=$!
#   PID_LIST+=" $pid";
# } done
cd api && buffalo dev &
PID_LIST+=$!
cd static && netlify dev &
PID_LIST+=$!
cd static && npm watch &
PID_LIST=$!

trap "kill $PID_LIST" SIGINT

echo "Parallel processes have started";

wait $PID_LIST

echo
echo "All processes have completed";
