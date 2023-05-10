package messages

import (
	"VK_BOT_RAW/internal/entity"
	"VK_BOT_RAW/internal/usecase/keyboard"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
)

type MessageHandler struct {
	token     string
	v         string
	keyboards map[string]interface{}
}

func NewMessageHandler(t string, v string) *MessageHandler {
	return &MessageHandler{token: t, v: v, keyboards: keyboard.GetKeyboardMap()}
}

func (h *MessageHandler) getMessageStruct(data interface{}) entity.Message {
	var newMessage entity.MessageParentResponse

	jsonString, err := json.Marshal(data)

	if err != nil {
		log.Print("ERR:", err)
	}

	json.Unmarshal(jsonString, &newMessage)

	return newMessage.Message
}

func (h *MessageHandler) NewMessage(resp interface{}) {
	msg := h.getMessageStruct(resp)

	var payload entity.Payload
	json.Unmarshal([]byte(msg.Payload), &payload)

	if payload.Command == "start" {
		kb := h.keyboards["BASE"].(func() string)()
		h.SendMessage("Привет!\nТыкай на кнопки!", msg.FromID, kb)
	}

	if payload.Button != "" {
		var kb, newMsg string

		if v, ok := h.keyboards[payload.Button]; ok {
			kb = v.(func() string)()
			newMsg = fmt.Sprintf("Вы нажали кнопку (%s)", msg.Text)
		} else {
			kb = h.keyboards["BASE"].(func() string)()
			newMsg = fmt.Sprintf("Кнопка (%s) на стадии разработки(", msg.Text)
			log.Printf("BUTTON %s ERROR", msg.Text)
		}
		h.SendMessage(newMsg, msg.FromID, kb)
	}
}

func (h *MessageHandler) SendMessage(message string, toId int, kb string) error {
	reqUrl := "https://api.vk.com/method/messages.send"
	data := url.Values{
		"access_token": {h.token},
		"v":            {h.v},
		"user_id":      {strconv.Itoa(toId)},
		"random_id":    {strconv.Itoa(rand.Intn(int(^uint(0) >> 1)))},
		"message":      {message},
	}

	if kb != "" {
		data.Add("keyboard", kb)
	}

	_, err := http.PostForm(reqUrl, data)
	if err != nil {
		return err
	}
	return nil
}
