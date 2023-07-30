#!/bin/bash
# -n 10000 -c 1000
#hey -n 10 -c 5 -m POST \
hey -n 300 -c 100 -m POST \
-d 'John Doe' \
"http://localhost:8080" #> go-extism-report.txt
