#!/bin/bash
# -n 10000 -c 1000
#hey -n 10 -c 5 -m POST \
#hey -n 6000 -c 3000 -m POST \
#hey -n 5 -c 1 -m POST \
hey -n 3000 -c 1000 -m POST \
-d 'John Doe' \
"http://localhost:8080" #> go-extism-report.txt
