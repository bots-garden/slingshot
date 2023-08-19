package slingshot

type RedisConfig struct {
	Id  string `json:"id"`
	Uri string `json:"uri"`
}

type RedisRecord struct {
	Id    string `json:"id"`
	Key   string `json:"key"`
	Value string `json:"value"`
}

type RedisMessage struct {
	Id      string `json:"id"`
	Channel string `json:"channel"`
	Payload string `json:"payload"`
}
