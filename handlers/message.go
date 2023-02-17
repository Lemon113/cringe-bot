package handlers

import (
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/lemon113/cringe-bot/data"
	"github.com/lemon113/cringe-bot/memes"
)

type Message struct {
	BotId     string
	BotPrefix string
	l         *log.Logger
	db        *data.DB
	memGen    *memes.Generator
}

func NewMessage(l *log.Logger, BotId, BotPrefix string, db *data.DB, memGen *memes.Generator) *Message {
	return &Message{BotId, BotPrefix, l, db, memGen}
}

func (m *Message) MessageHandler(s *discordgo.Session, mc *discordgo.MessageCreate) {
	if mc.Author.ID == m.BotId {
		return
	}

	m.l.Printf("Message from %s %#v", mc.Author, mc.Content)
	rs, hasBotRequest := hasBotRequest(mc.Content)
	if hasBotRequest {
		memS, hasMemRequest := hasMemeRequest(rs)
		if hasMemRequest {
			imgUrl, err := m.memGen.Generate(strings.Split(memS, ";"))
			if err == nil {
				s.ChannelMessageSend(mc.ChannelID, imgUrl.Data.URL)
			}
		}
	} else if m.db.HasAnyTriggerWords(mc.Content) {
		response := m.db.GenerateRandomResponse()
		s.ChannelMessageSend(mc.ChannelID, response)
	}
}

var botRequestTriggers []string = []string{
	"дед",
	"дедушка",
	"глад",
	"дедуля",
	"богдан",
}

func hasBotRequest(s string) (string, bool) {
	ls := strings.ToLower(s)
	for _, trigger := range botRequestTriggers {
		prefix := trigger + ","
		if strings.HasPrefix(ls, prefix) {
			return strings.TrimPrefix(ls, prefix), true
		}
	}
	return "", false
}

func hasMemeRequest(s string) (string, bool) {
	if strings.Contains(s, "мем:") {
		arr := strings.Split(s, "мем:")
		return arr[1], true
	}
	return "", false
}
