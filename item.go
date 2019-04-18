package main

type item struct {
	LevelRequired int `json:"levelRequired,omitempty"`
	Strength      int `json:"strength,omitempty"`
	Wisdom        int `json:"wisdom,omitempty"`
	Charisma      int `json:"charisma,omitempty"`
	Intelligence  int `json:"intelligence,omitempty"`
	Dexterity     int `json:"dexterity,omitempty"`
}
