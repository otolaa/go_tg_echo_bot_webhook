package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type DataStart struct {
	Ok      bool   `json:"ok"`
	Result  bool   `json:"result"`
	Version string `json:"version"`
	Error   bool   `json:"error"`
}

type From struct {
	Id           int64  `json:"id"`
	IsBot        bool   `json:"is_bot"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Username     string `json:"username"`
	LanguageCode string `json:"language_code"`
	IsPremium    bool   `json:"is_premium"`
}

type Chat struct {
	Id        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
	Type      string `json:"type"`
}

type Entities struct {
	Offset uint16 `json:"offset"`
	Length uint16 `json:"length"`
	Type   string `json:"bot_command"`
}

type Message struct {
	MessageID int64      `json:"message_id"`
	From      From       `json:"from"`
	Chat      Chat       `json:"chat"`
	Date      int        `json:"date"`
	Text      string     `json:"text"`
	Entities  []Entities `json:"entities"`
}

type CallbackQuery struct {
	ID   string `json:"id"`
	Data string `json:"data"`
}

type BotMessage struct {
	UpdateID      int64         `json:"update_id"`
	Message       Message       `json:"message"`
	CallbackQuery CallbackQuery `json:"callback_query"`
}

type InlineKeyboard [][]struct {
	Text         string `json:"text"`
	CallbackData string `json:"callback_data"`
}

type ReplyMarkup struct {
	InlineKeyboard InlineKeyboard `json:"inline_keyboard"`
}

type SendMessageBody struct {
	ChatID      int64       `json:"chat_id"`
	Text        string      `json:"text"`
	ParseMode   string      `json:"parse_mode"`
	ReplyMarkup ReplyMarkup `json:"reply_markup"`
}

type SetWebhook struct {
	Url            string `json:"url"`
	AllowedUpdates string `json:"allowed_updates"`
}

type AnswerCallbackQuery struct {
	CallbackQueryId string `json:"callback_query_id"`
	Text            string `json:"text"`
}

func answerCallbackQuery(CallbackQueryId string, Text string) error {
	res := &AnswerCallbackQuery{
		CallbackQueryId: CallbackQueryId,
		Text:            Text,
	}

	reqBytes, err := json.Marshal(res)
	if err != nil {
		return err
	}

	resp, err := http.Post(URL_API+TOKEN+"/answercallbackquery", "application/json", bytes.NewBuffer(reqBytes))
	if err != nil {
		fmt.Println(61)
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New("unexpected status" + resp.Status)
	}

	return err
}

func sendMessage(ChatID int64, Text string) error {
	reqBody := &SendMessageBody{
		ChatID:    ChatID,
		Text:      Text,
		ParseMode: "html",
		ReplyMarkup: ReplyMarkup{
			InlineKeyboard: InlineKeyboard{
				{{"понятно", "clear"}, {"неясно", "unclear"}},
			},
		},
	}

	reqBytes, err := json.Marshal(reqBody)
	if err != nil {
		return err
	}

	resp, err := http.Post(URL_API+TOKEN+"/sendMessage", "application/json", bytes.NewBuffer(reqBytes))
	if err != nil {
		fmt.Println(61)
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New("unexpected status" + resp.Status)
	}

	return err
}

func setWebhook() error {
	allowed_updates := []string{"message", "edited_channel_post", "callback_query"}
	jsonData, err := json.Marshal(allowed_updates)
	if err != nil {
		return err
	}

	setWebhook := &SetWebhook{
		Url:            URL + "/" + TOKEN,
		AllowedUpdates: string(jsonData),
	}

	buf, err := json.Marshal(setWebhook)
	if err != nil {
		return err
	}

	resp, err := http.Post(URL_API+TOKEN+"/setWebhook", "application/json", bytes.NewBuffer(buf))
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New("unexpected status" + resp.Status)
	}

	req := map[string]interface{}{}
	err = json.NewDecoder(resp.Body).Decode(&req)
	if err != nil {
		return err
	}

	fmt.Println(req)

	return err
}
