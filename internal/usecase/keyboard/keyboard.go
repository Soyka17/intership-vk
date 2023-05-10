package keyboard

import (
	"VK_BOT_RAW/internal/entity"
	"encoding/json"
)

func NewButton(tp string, pl string, lb string, clr string) *entity.Button {
	setting := entity.ButtonSettings{Type: tp, Payload: pl, Label: lb}
	return &entity.Button{Action: setting, Color: clr}
}

func GetKeyboardMap() map[string]interface{} {
	kbMap := make(map[string]interface{})
	kbMap["BASE"] = getBaseKeyboard
	kbMap["ACT1"] = getACT1Keyboard
	kbMap["ACT2"] = getACT2Keyboard
	kbMap["ACT3"] = getACT3Keyboard
	kbMap["ACT4"] = getACT4Keyboard
	kbMap["OK"] = getBaseKeyboard
	kbMap["NO"] = getBaseKeyboard

	return kbMap
}

func getBaseKeyboard() string {
	var keyboard entity.Keyboard

	keyboard.OneTime = false
	var buttons [][]entity.Button
	var buttRow []entity.Button
	act1B := NewButton("text", "{\"button\": \"ACT1\"}", "Действие 1", "primary")
	act2B := NewButton("text", "{\"button\": \"ACT2\"}", "Действие 2", "primary")
	act3B := NewButton("text", "{\"button\": \"ACT3\"}", "Действие 3", "primary")
	act4B := NewButton("text", "{\"button\": \"ACT4\"}", "Действие 4", "primary")
	buttRow = append(buttRow, *act1B)
	buttRow = append(buttRow, *act2B)
	buttons = append(buttons, buttRow)
	buttRow = nil
	buttRow = append(buttRow, *act3B)
	buttRow = append(buttRow, *act4B)
	buttons = append(buttons, buttRow)
	keyboard.Buttons = buttons

	kbJson, _ := json.Marshal(keyboard)
	return string(kbJson)
}

func getACT1Keyboard() string {
	var keyboard entity.Keyboard

	keyboard.OneTime = false

	var buttons [][]entity.Button
	var buttRow []entity.Button

	b1 := NewButton("text", "{\"button\": \"OK\"}", "Подтвердить", "positive")
	b2 := NewButton("text", "{\"button\": \"NO\"}", "Отклонить", "negative")
	b3 := NewButton("text", "{\"button\": \"BASE\"}", "На главную", "secondary")

	buttRow = append(buttRow, *b1)
	buttRow = append(buttRow, *b2)
	buttRow = append(buttRow, *b3)

	buttons = append(buttons, buttRow)
	keyboard.Buttons = buttons

	kbJson, _ := json.Marshal(keyboard)
	return string(kbJson)
}

func getACT2Keyboard() string {
	var keyboard entity.Keyboard

	keyboard.OneTime = false

	var buttons [][]entity.Button
	var buttRow []entity.Button

	b1 := NewButton("text", "{\"button\": \"OK\"}", "Подтвердить", "positive")
	b2 := NewButton("text", "{\"button\": \"NO\"}", "Отклонить", "negative")
	b3 := NewButton("text", "{\"button\": \"BASE\"}", "На главную", "secondary")

	buttRow = append(buttRow, *b1)
	buttRow = append(buttRow, *b2)
	buttRow = append(buttRow, *b3)

	buttons = append(buttons, buttRow)
	keyboard.Buttons = buttons

	kbJson, _ := json.Marshal(keyboard)
	return string(kbJson)
}

func getACT3Keyboard() string {
	var keyboard entity.Keyboard

	keyboard.OneTime = false

	var buttons [][]entity.Button
	var buttRow []entity.Button

	b1 := NewButton("text", "{\"button\": \"OK\"}", "Подтвердить", "positive")
	b2 := NewButton("text", "{\"button\": \"NO\"}", "Отклонить", "negative")
	b3 := NewButton("text", "{\"button\": \"BASE\"}", "На главную", "secondary")

	buttRow = append(buttRow, *b1)
	buttRow = append(buttRow, *b2)
	buttRow = append(buttRow, *b3)

	buttons = append(buttons, buttRow)
	keyboard.Buttons = buttons

	kbJson, _ := json.Marshal(keyboard)
	return string(kbJson)
}

func getACT4Keyboard() string {
	var keyboard entity.Keyboard

	keyboard.OneTime = false

	var buttons [][]entity.Button
	var buttRow []entity.Button

	b1 := NewButton("text", "{\"button\": \"OK\"}", "Подтвердить", "positive")
	b2 := NewButton("text", "{\"button\": \"NO\"}", "Отклонить", "negative")
	b3 := NewButton("text", "{\"button\": \"BASE\"}", "На главную", "secondary")

	buttRow = append(buttRow, *b1)
	buttRow = append(buttRow, *b2)
	buttRow = append(buttRow, *b3)

	buttons = append(buttons, buttRow)
	keyboard.Buttons = buttons

	kbJson, _ := json.Marshal(keyboard)
	return string(kbJson)
}
