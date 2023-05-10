package entity

type ButtonSettings struct {
	Type    string `json:"type"`
	Payload string `json:"payload"`
	Label   string `json:"label"`
}

type Button struct {
	Action ButtonSettings `json:"action"`
	Color  string         `json:"color"`
}
