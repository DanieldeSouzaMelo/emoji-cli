package main

import (
	"fmt"
	"os"
	"github.com/loudercake/emoji-cli/utils"
)


func main() {
	query := os.Args[1]
	emoji_list := utils.GetEmojis()
	fmt.Println(utils.SearchEmojis(&emoji_list, query))
}
