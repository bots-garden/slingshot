#!/bin/bash

curl --verbose \
http://localhost:7071 \
-H 'content-type: text/plain; charset=utf-8' \
-d '😄 Bob Morane'
echo ""

curl --verbose \
http://localhost:7072 \
-H 'content-type: text/plain; charset=utf-8' \
-d '😄 Bob Morane'
echo ""


#curl http://198.19.249.146:7071 -H "content-type: text/plain; charset=utf-8" -d "😄 Bob Morane"
#curl http://198.19.249.146:7072 -H "content-type: text/plain; charset=utf-8" -d "😄 Bob Morane"

curl https://hello-wasm.localto.net -H "content-type: text/plain; charset=utf-8" -d "😄 Bob Morane"

curl https://hello-one.localto.net -H "content-type: text/plain; charset=utf-8" -d "😄 Bob Morane"
curl https://hello-two.localto.net -H "content-type: text/plain; charset=utf-8" -d "😄 Bob Morane"

