package slingshot

type Result struct {
	Success []byte `json:"success"`
	Failure []byte `json:"failure"`
}

type StringResult struct {
	Success string `json:"success"`
	Failure string `json:"failure"`
}