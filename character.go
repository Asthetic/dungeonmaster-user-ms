package main

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

const (
	bard    Class = "Bard"
	cleric  Class = "Cleric"
	druid   Class = "Druid"
	paladin Class = "Paladin"
	ranger  Class = "Ranger"
	rogue   Class = "Rogue"
	warrior Class = "Warrior"
	wizard  Class = "Wizard"
)

// Character type defines an individual character
type Character struct {
	UID          int             `json:"uid,omitempty"`
	Name         string          `json:"string,omitempty"`
	Level        int             `json:"level,omitempty"`
	Strength     int             `json:"strength,omitempty"`
	Wisdom       int             `json:"wisdom,omitempty"`
	Charisma     int             `json:"charisma,omitempty"`
	Intelligence int             `json:"intelligence,omitempty"`
	Dexterity    int             `json:"dexterity,omitempty"`
	Class        string          `json:"class,omitempty"`
	Equipment    map[string]item `json:"equipment,omitempty"`
}

// Class is an enum that allows us to map to set classes
type Class string

// Profile returns a user's character profile
func Profile(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("method not yet implemented"))
}

// Create makes a new character for the user
func Create(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("method not yet implemented"))
}
