package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/atotto/clipboard"
	"github.com/kenshaw/emoji"
	"github.com/ktr0731/go-fuzzyfinder"
)

var gemoji = emoji.Gemoji()

func main() {
	var nFlag = flag.Bool("n", false, "no newline (for piping into other commands)")
	flag.Parse()

	idx, err := fuzzyfinder.Find(
		gemoji,
		func(i int) string {
			return gemoji[i].Description
		},
		fuzzyfinder.WithPromptString("Emoji: "),
		fuzzyfinder.WithPreviewWindow(func(i, w, h int) string {
			if i == -1 {
				return ""
			}
			return fmt.Sprintf("Emoji: %s\nAlias: %s\nDescription: %s",
				gemoji[i].Emoji,
				gemoji[i].Aliases,
				gemoji[i].Description)
		}))
	if err != nil {
		log.Fatal(err)
	}

	emoji := gemoji[idx].Emoji
	if *nFlag {
		fmt.Printf("%v", emoji)
	} else {
		fmt.Printf("%v\n", emoji)
	}
	clipboard.WriteAll(emoji)
}
