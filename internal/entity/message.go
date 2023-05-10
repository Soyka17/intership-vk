package entity

type Message struct {
	Date                  int    `json:"date"`
	FromID                int    `json:"from_id"`
	ID                    int    `json:"id"`
	Out                   int    `json:"out"`
	Attachments           []any  `json:"attachments"`
	ConversationMessageID int    `json:"conversation_message_id"`
	FwdMessages           []any  `json:"fwd_messages"`
	Important             bool   `json:"important"`
	IsHidden              bool   `json:"is_hidden"`
	Payload               string `json:"payload"`
	PeerID                int    `json:"peer_id"`
	RandomID              int    `json:"random_id"`
	Text                  string `json:"text"`
}

type MessageParentResponse struct {
	Message Message `json:"message"`
}
