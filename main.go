package main

import (
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/lemon113/cringe-bot/data"
	"github.com/lemon113/cringe-bot/handlers"
	"github.com/lemon113/cringe-bot/memes"
)

var (
	Token     string
	BotPrefix string
	logger    *log.Logger
	m         *handlers.Message
	db        *data.DB
	MemeGen   *memes.Generator
)

var BotId string
var goBot *discordgo.Session

func Start() {
	logger.Println("Creating discordgo session")
	goBot, err := discordgo.New("Bot " + Token)

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

	m = handlers.NewMessage(logger, BotId, BotPrefix, db, MemeGen)
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
	logger = log.New(os.Stdout, "cringe-bot ", log.LstdFlags)
	Token = os.Getenv("CRINGE_BOT_KEY")
	MemeGen = memes.NewGenerator(os.Getenv("IMGFLIP_USERNAME"), os.Getenv("IMGFLIP_PWD"), logger)
	db = data.NewDB(logger)
}

func main() {
	rand.Seed(time.Now().Unix())
	Start()
	<-make(chan struct{})
	return
}
