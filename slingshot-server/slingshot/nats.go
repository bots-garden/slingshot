package slingshot

type NatsConfig struct {
	Url string `json:"url"`
	Id  string `json:"id"`
}

type NatsSubscribeMessage struct {
	Subject string `json:"subject"`
	Data    string `json:"data"`
}

type NatsPublishMessage struct {
	Id      string `json:"id"`
	Url     string `json:"url"`
	Subject string `json:"subject"`
	Data    string `json:"data"`
}
