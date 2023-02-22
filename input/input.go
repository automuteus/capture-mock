package input

import (
	"fmt"
	"github.com/automuteus/automuteus/v8/pkg/game"
	"log"
	"strconv"
	"strings"
)

func String(prompt, def string) string {
	var input string
	output := prompt
	if def != "" {
		output += " [" + def + "]:"
	} else {
		output += ":"
	}
	log.Println(output)
	_, err := fmt.Scanln(&input)
	if err != nil {
		log.Println(err)
	}
	if input == "" {
		input = def
	}
	return input
}

func Region() game.Region {
	var input string
	output := fmt.Sprintf("Region? [%s]:\nN for North America\nE for Europe\nA for Asia", game.NA.ToString())
	log.Println(output)
	_, err := fmt.Scanln(&input)
	if err != nil {
		log.Println(err)
	}
	if input == "" {
		return game.NA
	}
	switch strings.ToUpper(string(input[0])) {
	case "N":
		return game.NA
	case "A":
		return game.AS
	case "E":
		return game.EU
	default:
		log.Println("I didn't recognize \"" + input + "\"")
		return Region()
	}
}

func PlayMap() game.PlayMap {
	var input string
	output := fmt.Sprintf("PlayMap? [%s]:\nS for Skeld\nM for Mira\nP for Polus\nD for dlekS\nA for Airship", game.MapNames[game.SKELD])
	log.Println(output)
	_, err := fmt.Scanln(&input)
	if err != nil {
		log.Println(err)
	}
	if input == "" {
		return game.SKELD
	}
	switch strings.ToUpper(string(input[0])) {
	case "S":
		return game.SKELD
	case "M":
		return game.MIRA
	case "P":
		return game.POLUS
	case "D":
		return game.DLEKS
	case "A":
		return game.AIRSHIP
	default:
		log.Println("I didn't recognize \"" + input + "\"")
		return PlayMap()
	}
}

func Color() int {
	var input string
	output := "Color? [red]:"
	for str, c := range game.ColorStrings {
		output += fmt.Sprintf("\n%v for %s", c, str)
	}
	log.Println(output)
	_, err := fmt.Scanln(&input)
	if err != nil {
		log.Println(err)
	}
	if input == "" {
		return game.Red
	}
	i, err := strconv.ParseInt(input, 10, 64)
	if err != nil {
		log.Println(err)
		return Color()
	}
	if i > game.Coral || i < 0 {
		log.Printf("I didn't recognize %v\n", i)
		return Color()
	}
	return int(i)
}

var playerActionStrings = []string{
	game.JOINED:       "JOINED",
	game.LEFT:         "LEFT",
	game.DIED:         "DIED",
	game.CHANGECOLOR:  "CHANGECOLOR",
	game.FORCEUPDATED: "FORCEUPDATED",
	game.DISCONNECTED: "DISCONNECTED",
	game.EXILED:       "EXILED",
}

func PlayerAction() game.PlayerAction {
	var input string
	output := "Player Action? [JOINED]:"
	for action, str := range playerActionStrings {
		output += fmt.Sprintf("\n%d for %s", action, str)
	}
	log.Println(output)
	_, err := fmt.Scanln(&input)
	if err != nil {
		log.Println(err)
	}
	if input == "" {
		return game.JOINED
	}
	i, err := strconv.ParseInt(input, 10, 64)
	if err != nil {
		log.Println(err)
		return PlayerAction()
	}
	if i > int64(game.EXILED) || i < 0 {
		log.Printf("I didn't recognize %v\n", i)
		return PlayerAction()
	}
	return game.PlayerAction(i)
}

func Phase() game.Phase {
	var input string
	output := "Game State? [LOBBY]:"
	for i, name := range game.PhaseNames {
		output += fmt.Sprintf("\n%d for %s", i, name)
	}
	log.Println(output)
	_, err := fmt.Scanln(&input)
	if err != nil {
		log.Println(err)
	}
	if input == "" {
		return game.LOBBY
	}
	i, err := strconv.ParseInt(input, 10, 64)
	if err != nil {
		log.Println(err)
		return Phase()
	}
	if i > int64(game.MENU) || i < 0 {
		log.Printf("I didn't recognize %v\n", i)
		return Phase()
	}
	return game.Phase(i)
}

var GameResultStrings = []string{
	game.HumansByVote:       "HumansByVote",
	game.HumansByTask:       "HumansByTask",
	game.ImpostorByVote:     "ImposterByVote",
	game.ImpostorByKill:     "ImposterByKill",
	game.ImpostorBySabotage: "ImposterBySabotage",
	game.ImpostorDisconnect: "ImposterDisconnect",
	game.HumansDisconnect:   "HumansDisconnect",
}

func GameResult() game.GameResult {
	var input string
	output := "Game Result? [HumansByVote]:"
	for i, name := range GameResultStrings {
		output += fmt.Sprintf("\n%d for %s", i, name)
	}
	log.Println(output)
	_, err := fmt.Scanln(&input)
	if err != nil {
		log.Println(err)
	}
	if input == "" {
		return game.HumansByVote
	}
	i, err := strconv.ParseInt(input, 10, 64)
	if err != nil {
		log.Println(err)
		return GameResult()
	}
	if i > int64(game.HumansDisconnect) || i < 0 {
		log.Printf("I didn't recognize %v\n", i)
		return GameResult()
	}
	return game.GameResult(i)
}

func Bool(prompt string, def bool) bool {
	var input string
	output := prompt + fmt.Sprintf(" [%v]:\nT for True\nF for False", def)
	log.Println(output)
	_, err := fmt.Scanln(&input)
	if err != nil {
		log.Println(err)
	}
	if input == "" {
		return def
	}
	switch strings.ToUpper(string(input[0])) {
	case "T":
		fallthrough
	case "Y":
		return true
	case "F":
		fallthrough
	case "N":
		return false
	default:
		log.Println("I didn't recognize \"" + input + "\"")
		return Bool(prompt, def)
	}
}
