// mplayer project mplayer.go
package main

import (
	"bufio"
	"fmt"
	"mplayer/mlib"
	"mplayer/mp"
	"os"
	"strconv"
	"strings"
)

func handlePlayCommand(lib *mlib.MusicManager, tokens []string) {
	if len(tokens) != 2 {
		fmt.Println("USAGE: play <name>")

	}
	e := lib.Find(tokens[1])
	if e == nil {
		fmt.Println("The music", tokens[1], "does not exist.")
		return
	}
	mp.Play(e.Source, e.Type)
}

func handleLibCommands(lib *mlib.MusicManager, tokens []string) {
	switch tokens[1] {
	case "list":
		for i := 0; i < lib.Len(); i++ {
			e, _ := lib.Get(i)
			fmt.Println(i+1, ":", e.Name, e.Artist, e.Source, e.Type)
		}
	case "add":
		if len(tokens) == 6 {
			l := lib.Len()
			id := l
			if l > 0 {
				m, _ := lib.Get(l - 1)
				v, _ := strconv.Atoi(m.Id)
				id = v + 1
			}
			lib.Add(&mlib.MusicEntry{strconv.Itoa(id), tokens[2], tokens[3], tokens[4], tokens[5]})
		} else {
			fmt.Println("USAGE: lib add <name> <artist> <source> <type>")
		}
	case "remove":
		if len(tokens) == 3 {
			lib.RemoveByName(tokens[2])
		} else {
			fmt.Println("USAGE: lib remove <name>")
		}
	default:
		fmt.Println("Unrecognized lib command:", tokens[1])
	}
}

func main() {

	fmt.Println(` 
Enter following commands to control the player: 

lib list -- View the existing music lib 

lib add <name><artist><source><type> -- Add a music to the music lib 

lib remove <name> -- Remove the specified music from the lib 
`)

	lib := mlib.NewMusicManager()
	r := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Enter command->")
		rawLine, _, _ := r.ReadLine()
		line := string(rawLine)
		if line == "q" || line == "e" {
			break
		}
		tokens := strings.Split(line, " ")
		if tokens[0] == "lib" {
			handleLibCommands(lib, tokens)
		} else if tokens[0] == "play" {
			handlePlayCommand(lib, tokens)
		} else {
			fmt.Println("Unrecognized command:", tokens[0])
		}

	}

}
