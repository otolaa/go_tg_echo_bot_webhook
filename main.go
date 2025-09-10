package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func showJsonData(w http.ResponseWriter, r *http.Request) {
	dataStart := &DataStart{
		Ok:      true,
		Result:  true,
		Version: VERSION,
		Error:   false,
		Method:  r.Method,
	}

	jsonData, _ := json.Marshal(dataStart)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(jsonData)
}

func show404(w http.ResponseWriter, r *http.Request) {
	data404 := &DataStart{
		Ok:      false,
		Result:  false,
		Version: VERSION,
		Error:   true,
		Method:  r.Method,
	}

	json404, _ := json.Marshal(data404)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(404)
	w.Write(json404)
}

func showStart(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		show404(w, r)
		return
	}

	showJsonData(w, r)
}

func showTg(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		show404(w, r)
		return
	}

	data := BotMessage{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		fmt.Println("66 → Error:", err)
		return
	}

	handleCallbackQuery(&data)
	handleMessage(&data)

	showJsonData(w, r)
}

func handleCallbackQuery(data *BotMessage) {
	if data.CallbackQuery.Data == "" {
		return
	}

	err := answerCallbackQuery(data.CallbackQuery.ID, data.CallbackQuery.Data)
	if err != nil {
		fmt.Println("73 → Error:", err)
		return
	}
}

func handleMessage(data *BotMessage) {
	if data.Message.Text == "" {
		return
	}

	p(2, data.Message.Chat.Id, " ~ ", data.Message.From.Username, " ~ ", data.Message.Text)

	LineKeyboard := InlineKeyboard{}
	MessageText := data.Message.Text

	if data.Message.Text == "/start" {
		LineKeyboard = InlineKeyboard{
			{{"понятно", "clear"}, {"неясно", "unclear"}},
		}

		MessageText = "Это тестовый бот для моделей ~ tg.go"
	}

	if data.Message.Text == "/version" {
		MessageText = "версия бота: " + VERSION
	}

	Markup := ReplyMarkup{
		InlineKeyboard: LineKeyboard,
	}

	err := sendMessage(data.Message.Chat.Id, MessageText, Markup)
	if err != nil {
		fmt.Println("81 → Error:", err)
		return
	}
}

func main() {
	p(4, TOKEN)
	p(4, URL)
	p(4, SuffixLine)

	http.HandleFunc("/", showStart)
	http.HandleFunc("/"+TOKEN, showTg)
	p(3, "Server starting on :8080...")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}
