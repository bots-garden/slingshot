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

