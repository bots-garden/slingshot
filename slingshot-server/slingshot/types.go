package slingshot

// Used by Host Function Callbacks
type Result struct {
	Success []byte `json:"success"`
	Failure []byte `json:"failure"`
}

// Used by Host Function Callbacks
type StringResult struct {
	Success string `json:"success"`
	Failure string `json:"failure"`
}

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