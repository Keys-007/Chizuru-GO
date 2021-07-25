package main

import (
	"Chizuru-GO/chizuru/bot"
	"Chizuru-GO/chizuru/config"
	"Chizuru-GO/chizuru/database"
	"Chizuru-GO/chizuru/telegraph"
	"fmt"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/bigkevmcd/go-configparser"
	"log"
	"net/http"
)

func main() {
	p, err := configparser.NewConfigParserFromFile("chizuru.cfg")
	if err != nil {
		log.Fatal("an error in getting cfg file", err)
	}

	v, err := p.Get("bot", "token")
	if err != nil {
		log.Fatal("an error in getting token: ", err)
	}
	config.SetVars(p)
	// Create bot
	b, err := gotgbot.NewBot(v, &gotgbot.BotOpts{
		Client:      http.Client{},
		GetTimeout:  gotgbot.DefaultGetTimeout,
		PostTimeout: gotgbot.DefaultPostTimeout,
	})
	if err != nil {
		log.Fatal("failed to create new bot: ", err)
	}

	// Create updater and dispatcher.
	updater := ext.NewUpdater(nil)
	dispatcher := updater.Dispatcher
	bot.LoadHandlers(dispatcher)
	err = updater.StartPolling(b, &ext.PollingOpts{DropPendingUpdates: true})
	if err != nil {
		log.Fatal("failed to start polling: ", err.Error())
	}
	database.StartDB()
	database.EnsureBotInDb(b)
	tf, err := telegraph.Register(b.User.FirstName, b.User.Username, fmt.Sprintf("https://t.me/%s", b.User.Username))
	if err != nil {
		log.Fatalf("Error while registering on telegra.ph: %s", err.Error())
		return
	}
	config.TelegraphToken = tf.Result.AccessToken

	log.Printf("Logged in to telegra.ph as %s", tf.Result.AuthorName)
	log.Printf("%s has been started...\n", b.User.Username)

	// Idle, to keep updates coming in, and avoid bot stopping.
	updater.Idle()
}
