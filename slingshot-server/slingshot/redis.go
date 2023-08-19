package slingshot

type RedisClientRecord struct {
	Id  string `json:"id"`
	Uri string `json:"uri"`
}

type RedisClientArguments struct {
	Id    string `json:"id"`
	Key   string `json:"key"`
	Value string `json:"value"`
}

type RedisClientMessageArguments struct {
	Id      string `json:"id"`
	Channel string `json:"channel"`
	Payload string `json:"payload"`
}
