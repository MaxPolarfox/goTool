package client

type Options struct {
	URL        string             `json:"url"`
	Headers    *map[string]string `json:"headers,omitempty"`
	RetryCount int                `json:"retryCount"`
	TimeoutMs  int                `json:"timeoutMs"`
}
