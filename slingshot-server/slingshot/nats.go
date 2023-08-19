package slingshot

type NatsConfig struct {
	Id  string `json:"id"`
	Url string `json:"url"`
}

type NatsMessage struct {
	Id      string `json:"id"`
	Subject string `json:"subject"`
	Data    string `json:"data"`
}
