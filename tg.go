package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

type DataStart struct {
	Ok      bool   `json:"ok"`
	Result  bool   `json:"result"`
	Version string `json:"version"`
	Error   bool   `json:"error"`
	Method  string `json:"metod"`
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

type BotCommand struct {
	Command     string `json:"command"`
	Description string `json:"description"`
}

func answerCallbackQuery(CallbackQueryId string, Text string) error {
	res := &AnswerCallbackQuery{
		CallbackQueryId: CallbackQueryId,
		Text:            Text,
	}

	err := setPost("answercallbackquery", res)
	if err != nil {
		return err
	}

	return err
}

func sendMessage(ChatID int64, Text string, Markup ReplyMarkup) error {
	res := &SendMessageBody{
		ChatID:      ChatID,
		Text:        Text,
		ParseMode:   "html",
		ReplyMarkup: Markup,
	}

	err := setPost("sendMessage", res)
	if err != nil {
		return err
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

	err = setPost("setWebhook", setWebhook)
	if err != nil {
		return err
	}

	return err
}

func setMyCommands() error {

	commands := []BotCommand{}
	commanStart := BotCommand{
		Command:     "start",
		Description: "🚀 The start bot",
	}

	commanVersion := BotCommand{
		Command:     "version",
		Description: "👾 version bot",
	}

	commands = append(commands, commanStart, commanVersion)

	jsonData, err := json.Marshal(commands)
	if err != nil {
		return err
	}

	BotCommand := &map[string]string{
		"commands": string(jsonData),
	}

	err = setPost("setMyCommands", BotCommand)
	if err != nil {
		return err
	}

	return err
}

// return post from URL method
func setPost(method string, data any) error {
	buf, err := json.Marshal(data)
	if err != nil {
		return err
	}

	urlPost := URL_API + TOKEN
	urlPost += "/" + method

	resp, err := http.Post(urlPost, "application/json", bytes.NewBuffer(buf))
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New("unexpected status " + resp.Status)
	}

	req := map[string]interface{}{}
	err = json.NewDecoder(resp.Body).Decode(&req)
	if err != nil {
		return err
	}

	// p(2, req)

	// for k, v := range req {
	// 	p(2, k, " ~ ", v)
	// }

	return err
}
