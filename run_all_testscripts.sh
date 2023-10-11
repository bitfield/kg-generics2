#!/bin/bash

set -e
for LISTING in solutions/*; do
  # Update deps:
  # cd $LISTING && go get -t -u && cd -
  ./run_testscript.sh $LISTING
done