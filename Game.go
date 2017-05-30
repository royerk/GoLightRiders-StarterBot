package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Game class, based on the python bot
type Game struct {
	initialTimebank, lastTimebank int
	timePerMove                   int
	playerNames                   [2]string
	myBot                         string
	myBotID, otherBotID           int
	fieldWidth, fieldHeight       int
	field                         Field
	round                         int
	players                       [2]Player
}

func (g *Game) initGame() {
	g.timePerMove = 10
	g.playerNames[0] = "I am not set"
	g.playerNames[1] = "Other ain t set"
	g.myBot = "I am still not set"
	g.myBotID = -1
	g.otherBotID = -1
	g.field = Field{}
	g.players[0] = Player{}
	g.players[1] = Player{}
}
func (g Game) myPlayer() Player {
	return g.players[g.myBotID]
}
func (g Game) otherPlayer() Player {
	return g.players[g.otherBotID]
}
func (g *Game) update(message []string) {
	if strings.Compare(message[0], "settings") == 0 {
		if strings.Compare(message[1], "timebank") == 0 {
			g.lastTimebank, _ = strconv.Atoi(message[2])
		} else if strings.Compare(message[1], "time_per_move") == 0 {
			g.timePerMove, _ = strconv.Atoi(message[2])
		} else if strings.Compare(message[1], "player_names") == 0 {
			names := strings.Split(message[2], ",")
			g.playerNames[0] = names[0]
			g.playerNames[1] = names[1]
		} else if strings.Compare(message[1], "your_bot") == 0 {
			g.myBot = message[2]
		} else if strings.Compare(message[1], "your_botid") == 0 {
			g.myBotID, _ = strconv.Atoi(message[2])
			g.otherBotID = 1 - g.myBotID
		} else if strings.Compare(message[1], "field_width") == 0 {
			g.fieldWidth, _ = strconv.Atoi(message[2])
		} else if strings.Compare(message[1], "field_height") == 0 {
			g.fieldHeight, _ = strconv.Atoi(message[2])
		} else if strings.Compare(message[1], "timebank") == 0 {
			g.lastTimebank, _ = strconv.Atoi(message[2])
		} else {
			fmt.Fprintln(os.Stderr, "Can't read settings:", message, "in game.update.")
		}
	} else if strings.Compare(message[0], "update") == 0 {
		if strings.Compare(message[1], "game") == 0 {
			if strings.Compare(message[2], "round") == 0 {
				g.round, _ = strconv.Atoi(message[3])
			} else if strings.Compare(message[2], "field") == 0 {
				g.field.parse(message[3])
			}
		}
	} else if strings.Compare(message[0], "action") == 0 && strings.Compare(message[1], "move") == 0 {
		g.lastTimebank, _ = strconv.Atoi(message[2])
	} /*else {
		fmt.Fprintln(os.Stderr, "Can't read:", message, "in game.update.")
	}*/
}
func (g *Game) run(bot Bot) {
	reader := bufio.NewReader(os.Stdin)
	fieldIsSet := false
	for true {
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		message := strings.Split(text, " ")
		g.update(message)
		if !fieldIsSet && g.fieldHeight != 0 && g.fieldWidth != 0 {
			g.field.initField(g.fieldHeight, g.fieldWidth)
			fieldIsSet = true
		}
		if strings.Compare(message[0], "action") == 0 && strings.Compare(message[1], "move") == 0 {
			bot.play(*g)
		}
	}
}
