module redis-pub-pdk

go 1.20

require (
	github.com/bots-garden/slingshot/go-pdk v0.0.0-00010101000000-000000000000
	github.com/extism/go-pdk v0.0.0-20230816024928-ee09fee7466e // indirect
)

require github.com/valyala/fastjson v1.6.3 // indirect

replace github.com/bots-garden/slingshot/go-pdk => ../../../go-pdk
