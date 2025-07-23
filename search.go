package main

import (
	"encoding/json"
	"os"
	"log"
	"fmt"
	"strings"
)

const emoji_json_path = "./data-by-emoji.json"

type Emoji struct {
	Emoji string 
	Name string 
	Skin_tone_support bool 
}


func parse_emoji_json() []Emoji {
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
	fmt.Println(emojis)
	return emojis
}

func search_emoji(emojis *[]Emoji, query string) {
	query = strings.ToUpper(query)
	for _, emoji := range *emojis {
		emoji_name := strings.ToUpper(emoji.Name)
		if (strings.Contains(emoji_name, query)) {
			fmt.Println(emoji.Name)
			fmt.Println(emoji.Emoji)
		}
	}
}	

func main() {
	emojis := parse_emoji_json()
	search_emoji(&emojis, "eggplant")
}
