#!/bin/bash
set -o allexport; source .env; set +o allexport
# you need these environment variables
# - REDIS_URI
echo "REDIS_URI: ${REDIS_URI}"
go test 
