package slingshot

type HTTPResponse struct {
	JsonBody   map[string]interface{} `json:"jsonBody"`
	TextBody   string                 `json:"textBody"`
	Headers    map[string]string      `json:"headers"`
	StatusCode int                    `json:"statusCode"`
}

type HTTPRequest struct {
	Method string `json:"method"`
	BaseUrl string `json:"baseUrl"`
	Body   string `json:"body"`
	//JsonBody   map[string]interface{} `json:"jsonBody"`
	//TextBody   string                 `json:"textBody"`
	Headers    map[string]string      `json:"headers"`
}
