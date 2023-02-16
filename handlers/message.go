package handlers

import (
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/lemon113/cringe-bot/data"
)

type Message struct {
	BotId     string
	BotPrefix string
	l         *log.Logger
	db        *data.DB
}

func NewMessage(l *log.Logger, BotId, BotPrefix string, db *data.DB) *Message {
	return &Message{BotId, BotPrefix, l, db}
}

func (m *Message) MessageHandler(s *discordgo.Session, mc *discordgo.MessageCreate) {
	if mc.Author.ID == m.BotId {
		return
	}

	m.l.Printf("Message from %s %#v", mc.Author, mc.Content)
	if m.db.HasAnyTriggerWords(mc.Content) {
		response := m.db.GenerateRandomResponse()
		s.ChannelMessageSend(mc.ChannelID, response)
	}
}
