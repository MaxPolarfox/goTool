package ihttp

type Options struct {
	URL        string             `json:"url"`
	Headers    *map[string]string `json:"headers,omitempty"`
	TimeoutMs  int                `json:"timeoutMs"`
}
