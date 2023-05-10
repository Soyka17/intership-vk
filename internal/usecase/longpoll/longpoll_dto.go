package longpoll

type GetLongPollServerResponse struct {
	Response struct {
		Key    string `json:"key"`
		Server string `json:"server"`
		Ts     string `json:"ts"`
	} `json:"response"`
}

type LongPollResponse struct {
	Ts      string          `json:"ts"`
	Updates []LongPollEvent `json:"updates"`
	Failed  int             `json:"failed"`
}

type LongPollEvent struct {
	Type    string      `json:"type"`
	EventID string      `json:"event_id"`
	V       string      `json:"v"`
	Object  interface{} `json:"object"`
	GroupID int         `json:"group_id"`
}
