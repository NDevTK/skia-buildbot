#!/bin/bash

# List all secrets available.

REL=$(dirname "$0")
source ${REL}/config.sh

berglas list ${BUCKET_ID}
