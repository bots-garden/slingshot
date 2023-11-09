#!/bin/bash

curl --verbose \
http://localhost:7070 \
-H 'content-type: text/plain; charset=utf-8' \
-d '😄 Bob Morane'
echo ""

