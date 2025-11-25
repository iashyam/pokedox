package main

import (
	"bufio"
	"fmt"
	"os"
	"pokedox/internal/pokecache"
	"strings"
)

var conf = NewConfig()
var cache = pokecache.NewCache()

func cleanInput(text string) []string {
	if text == "" {
		return []string{}
	}

	text_smallcase := strings.ToLower(text)
	removed_spaces := strings.Trim(text_smallcase, " ")

	splitted := strings.Split(removed_spaces, " ")

	return splitted
}

func main() {

	commands := getCommands(conf, cache)

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedox> ")
		scanner.Scan()
		text := scanner.Text()
		cleaned_text := cleanInput(text)
		if len(cleaned_text) == 0 {
			continue
		}
		command := cleaned_text[0]
		args := cleaned_text[1:]
		cli_command, exists := commands[command]
		if exists {
			conf.History = append(conf.History, strings.Join(cleaned_text, " "))
			err := cli_command.execute(args...)
			if err != nil {
				fmt.Println("Error executing command:", err)
			}
			continue
		} else {
			fmt.Printf("Unknown command: %s\n", command)
		}
	}
}
