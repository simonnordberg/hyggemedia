package file

type Change struct {
	Error  error  `json:"error,omitempty"`
	Source string `json:"source"`
	Target string `json:"target"`
}

type Changes []*Change
