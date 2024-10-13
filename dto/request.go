package dto

type Request struct {
	Command    string      `json:"command"`
	Key        string      `json:"key,omitempty"`
	Value      interface{} `json:"value,omitempty"`
	Expiration int64       `json:"expiration,omitempty"`
}

type Requests struct {
	Requests []Request `json:"requests"`
}
