#!/bin/bash
curl -X POST --verbose \
http://localhost:8080 \
-H 'content-type: text/plain; charset=utf-8' \
-d '😄 Bob Morane'
echo ""


#curl -X POST \
#http://extism-arm.local:8080 \
#-H 'content-type: text/plain; charset=utf-8' \
#-d '😄 Bob Morane'
#echo ""


#curl --verbose http://localhost:8080