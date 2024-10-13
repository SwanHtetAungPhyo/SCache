package dto

type Request struct {
	Command string      `json:"command"`
	Key     string      `json:"key,omitempty"`
	Value   interface{} `json:"value,omitempty"`
}