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

	if data.CallbackQuery.Data != "" {
		err = answerCallbackQuery(data.CallbackQuery.ID, data.CallbackQuery.Data)
		if err != nil {
			fmt.Println("73 → Error:", err)
			return
		}
	}

	if data.Message.Text != "" {
		err = sendMessage(data.Message.Chat.Id, data.Message.Text)
		if err != nil {
			fmt.Println("81 → Error:", err)
			return
		}
	}

	showJsonData(w, r)
}

func main() {
	fmt.Println(TOKEN)
	fmt.Println(URL)

	http.HandleFunc("/", showStart)
	http.HandleFunc("/"+TOKEN, showTg)
	fmt.Println("Server starting on :8080...")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}
