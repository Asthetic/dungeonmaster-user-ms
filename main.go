package main

import (
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/caarlos0/env"
	"gopkg.in/yaml.v2"
)

// config is the configuration of the bot
type config struct {
	Token string `env:"DISCORDTOKEN" yaml:"token"`
}

func getConfig(cfg *config) *config {
	if err := env.Parse(cfg); err != nil {
		log.Printf("unable to parse enviornment variables: %v\n", err)
	}

	yamlCfg, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Printf("unable to read config.yaml: %v", err)
		return nil
	}

	err = yaml.Unmarshal(yamlCfg, cfg)
	if err != nil {
		log.Printf("unable to parse config.yaml: %v\n", err)
	}

	return cfg
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	// If the message is from our bot return, don't process anything
	if s.State.Ready.User.ID == m.Author.ID {
		return
	}

	log.Printf("channelID: %s, authorName: %s, authorID: %s, content: %s\n", m.ChannelID, m.Author.Username, m.Author.ID, m.Content)

	if m.Content[:1] == "!" {
		command := strings.SplitN(strings.ToLower(m.Content[1:]), " ", 2)
		method := ""
		if len(command) > 0 {
			method = command[0]
		}

		switch method {
		case "profile":
			go Profile(s, m)
		case "create":
			go Create(s, m)
		}
	}
}

func main() {
	cfg := getConfig(&config{})

	var err error

	discord, err := discordgo.New("Bot " + cfg.Token)
	if err != nil {
		log.Fatalf("unable to create discordgo instance: %v\n", err)
	}

	discord.AddHandler(messageHandler)

	err = discord.Open()
	defer discord.Close()
	if err != nil {
		log.Fatalf("unable to connect to discord: %v\n", err)
	}

	log.Printf("Listening for discord messages\n")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
}
