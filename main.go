package main

import (
	"flag"
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/lemon113/cringe-bot/data"
	"github.com/lemon113/cringe-bot/handlers"
)

var (
	Token     *string
	BotPrefix string
	logger    *log.Logger
	m         *handlers.Message
	db        *data.DB
)

var BotId string
var goBot *discordgo.Session

func Start() {
	logger.Println("Creating goBot ", *Token, BotPrefix)
	goBot, err := discordgo.New("Bot " + *Token)

	if err != nil {
		logger.Println(err.Error())
		return
	}

	logger.Println("Creating user")
	u, err := goBot.User("@me")

	if err != nil {
		logger.Println(err.Error())
		return
	}

	BotId = u.ID

	m = handlers.NewMessage(logger, BotId, BotPrefix, db)
	goBot.AddHandler(messageHandler)
	log.Print("Opening session")
	goBot.Identify.Intents |= discordgo.IntentMessageContent
	err = goBot.Open()

	if err != nil {
		logger.Println(err.Error())
		return
	}

	logger.Println("Bot is running!")
}

func messageHandler(s *discordgo.Session, mc *discordgo.MessageCreate) {
	m.MessageHandler(s, mc)
}

func init() {
	Token = flag.String("t", "", "discord bot token")
}

func main() {
	flag.Parse()
	logger = log.New(os.Stdout, "cringe-bot ", log.LstdFlags)
	db = data.NewDB(logger)
	logger.Printf("%#v\n", *db)

	Start()
	<-make(chan struct{})
	return
}
