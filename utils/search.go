package utils

import (
	"encoding/json"
	"os"
	"log"
	"strings"
)

const emoji_json_path = "./data-by-emoji.json"

type Emoji struct {
	Emoji string 
	Name string 
	Skin_tone_support bool 
}


func GetEmojis() []Emoji {
	file, err := os.ReadFile(emoji_json_path)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(string(file))
	var emojis []Emoji
	err = json.Unmarshal(file, &emojis)
	if err != nil {
		log.Fatal(err)
	}
	return emojis
}

func SearchEmojis(emojis *[]Emoji, query string) []Emoji {
	query = strings.ToUpper(query)
	var emoji_array []Emoji
	for _, emoji := range *emojis {
		emoji_name := strings.ToUpper(emoji.Name)
		if (strings.Contains(emoji_name, query)) {
			emoji_array = append(emoji_array, emoji)
		}
	}
	return emoji_array
}	

