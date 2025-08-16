package main

import (
	"log"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	token := "8262839503:AAGeXC5t_TwuvJH5A0ZZuR6hoHzKuV_5CPg"

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatal(err)
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)

	waiting := map[int64]bool{}
	menu := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Yordam"),
		),
	)

	for update := range updates {
		if update.Message == nil {
			continue
		}
		chatID := update.Message.Chat.ID
		text := strings.TrimSpace(update.Message.Text)

		if text == "/start" {
			waiting[chatID] = false
			msg := tgbotapi.NewMessage(chatID, "Assalomu alaykum!\nSiz bilan Ijara.uz mas’ul xizmati bog‘lanmoqda.\n\nXizmatdan foydalanish, hamkorlik yoki boshqa masalalar bo‘yicha savolingiz bo‘lsa, menyudan tanlang yoki yozib qoldiring.")
			msg.ReplyMarkup = menu
			bot.Send(msg)
			continue
		}

		if strings.EqualFold(text, "Yordam") || text == "/help" {
			waiting[chatID] = true
			msg := tgbotapi.NewMessage(chatID, "Iltimos, savolingizni batafsil yozib qoldiring. Mas’ul guruh 24 soat ichida siz bilan bog‘lanadi.")
			msg.ReplyMarkup = menu
			bot.Send(msg)
			continue
		}

		if waiting[chatID] {
			waiting[chatID] = false
			msg := tgbotapi.NewMessage(chatID, "Savolingiz qabul qilindi. Mas’ul guruh 24 soat ichida siz bilan bog‘lanadi. Rahmat!")
			msg.ReplyMarkup = menu
			bot.Send(msg)
			continue
		}
	}
}
