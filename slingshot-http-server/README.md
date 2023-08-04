# First HTTP Server

```bash
curl -v -X POST \
http://localhost:8080 \
-H 'content-type: text/plain; charset=utf-8' \
-d '😄 Bob Morane'
```

## Load testing

```bash
hey -n 300 -c 100 -m POST \
-d 'John Doe' \
"http://localhost:8080" 
```

```bash
Summary:
  Total:	0.0669 secs
  Slowest:	0.0467 secs
  Fastest:	0.0005 secs
  Average:	0.0171 secs
  Requests/sec:	4484.9541
  
  Total data:	9900 bytes
  Size/request:	33 bytes

Response time histogram:
  0.001 [1]	|
  0.005 [4]	|■
  0.010 [54]	|■■■■■■■■■■■■■■■■■■■
  0.014 [111]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.019 [39]	|■■■■■■■■■■■■■■
  0.024 [10]	|■■■■
  0.028 [16]	|■■■■■■
  0.033 [35]	|■■■■■■■■■■■■■
  0.037 [21]	|■■■■■■■■
  0.042 [6]	|■■
  0.047 [3]	|■


Latency distribution:
  10% in 0.0081 secs
  25% in 0.0101 secs
  50% in 0.0123 secs
  75% in 0.0256 secs
  90% in 0.0330 secs
  95% in 0.0348 secs
  99% in 0.0458 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0045 secs, 0.0005 secs, 0.0467 secs
  DNS-lookup:	0.0006 secs, 0.0000 secs, 0.0077 secs
  req write:	0.0014 secs, 0.0001 secs, 0.0271 secs
  resp wait:	0.0087 secs, 0.0002 secs, 0.0149 secs
  resp read:	0.0025 secs, 0.0000 secs, 0.0128 secs

Status code distribution:
  [200]	300 responses
```