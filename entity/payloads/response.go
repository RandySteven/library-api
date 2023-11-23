package payloads

type Response struct {
	Message string   `json:"message,omitempty"`
	Errors  []string `json:"errors,omitempty"`
	Data    any      `json:"data,omitempty"`
}
