package slingshot

import "github.com/valyala/fastjson"

var parser = fastjson.Parser{}

// GetJsonFromBytes
/*
	Convert a buffer (`[]byte`) into a JSON value
*/
func GetJsonFromBytes(buffer []byte)  (*fastjson.Value, error) {
	return parser.ParseBytes(buffer)
}
