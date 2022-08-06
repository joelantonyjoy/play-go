package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

// all types

type Game struct {
	plateCount int
	towerCount int

	towers []Tower
	plates []Plate
}

type Plate struct {
	size int
}

type Tower struct {
	slots []Plate
}

func NewGame(plateCount, towerCount int) *Game {
	game := &Game{
		plateCount: plateCount,
		towerCount: towerCount,
	}
	game.initGame()
	return game
}

func (g *Game) initGame() {
	g.towers = make([]Tower, g.towerCount)
	g.plates = make([]Plate, g.plateCount)

	for i := 0; i < g.plateCount; i++ {
		g.plates[i] = Plate{size: i + 1}
	}

	for i := 0; i < g.towerCount; i++ {
		g.towers[i] = Tower{}
		g.towers[i].slots = make([]Plate, g.plateCount)
		for j := g.plateCount - 1; j >= 0; j-- {
			if i == 0 {
				g.towers[i].push(g.plates[j])
			} else {
				g.towers[i].push(Plate{size: 0})
			}
		}
	}

}

func (g *Game) move(from, to int) bool {
	if !(from >= 1 && from <= g.towerCount) {
		return false
	}
	if !(to >= 1 && to <= g.towerCount) {
		return false
	}
	fromTower := g.towers[from-1]
	toTower := g.towers[to-1]

	fromTowerTopPlate := g.findTopPlate(fromTower)
	toTowerTopPlate := g.findTopPlate(toTower)

	if fromTowerTopPlate.size == 0 {
		return false
	}
	if (fromTowerTopPlate.size < toTowerTopPlate.size) || toTowerTopPlate.size == 0 {
		plate := fromTower.pop()
		toTower.push(plate)
		return true
	}

	return false
}

func (g *Game) findTopPlate(tower Tower) Plate {
	for _, slot := range tower.slots {
		if slot.size != 0 {
			return slot
		}
	}
	return Plate{}
}

func (g *Game) over() bool {
	lastTower := g.towers[g.towerCount-1]

	for _, plate := range lastTower.slots {
		if plate.size == 0 {
			return false
		}
	}

	return true
}

func (g *Game) print() {
	// gap between tower
	gap := pattern(' ', g.plateCount)

	// print without plates
	for i := 0; i < 2; i++ {
		for j := 0; j < g.towerCount; j++ {
			empty := pattern(' ', g.plateCount)

			p(gap)
			p(empty)
			p("|")
			p(empty)
		}
		pln("")
	}

	// print with plates
	for i := 0; i < g.plateCount; i++ {
		for j := 0; j < g.towerCount; j++ {
			p(gap)
			plateView := g.towers[j].slots[i].view(g.plateCount)
			p(plateView)
		}
		pln("")
	}

	for j := 0; j < g.towerCount; j++ {
		empty := pattern(' ', g.plateCount-1)

		p(gap)
		p(empty)
		pf("(%d)", j+1)
		p(empty)
	}
	pln("")
}

func (g *Game) debug() {
	for i := 0; i < g.plateCount; i++ {
		for j := 0; j < g.towerCount; j++ {
			pf("%d", g.towers[j].slots[i].size)
		}
		pln("")
	}
}

func (t *Tower) push(plate Plate) bool {
	for i := len(t.slots) - 1; i >= 0; i-- {
		if t.slots[i].size == 0 {
			t.slots[i] = plate
			return true
		}
	}
	return false
}

func (t *Tower) pop() Plate {
	for i := 0; i < len(t.slots); i++ {
		if t.slots[i].size != 0 {
			plate := t.slots[i]
			t.slots[i] = Plate{size: 0}
			return plate
		}
	}

	for _, plate := range t.slots {
		if plate.size != 0 {
			return plate
		}
	}
	return Plate{}
}

func (p *Plate) view(plateCount int) string {
	empty := pattern(' ', plateCount-p.size)
	print := pattern('-', p.size)

	parts := make([]string, 0)
	parts = append(parts, empty, print, "|", print, empty) // ***|***
	return strings.Join(parts, "")
}

func main() {
	var err error
	totalMove := 0
	lastMoveStatus := "make your first move"
	console := bufio.NewReader(os.Stdin)

	towerCount := 0
	platCount := 0

	clearScreen()

	fmt.Print(`Enter number of tower (3 default):`)
	input, _, _ := console.ReadLine()
	inputStr := string(input)
	if len(inputStr) == 0 {
		towerCount = 3
	} else {
		towerCount, err = strconv.Atoi(string(inputStr))
		if err != nil {
			panic("invalid input")
		}
		if towerCount < 2 {
			panic("invalid input, at least 2 towers needed")
		}
	}

	fmt.Print(`Enter number of plates (3 default):`)
	input, _, _ = console.ReadLine()
	inputStr = string(input)
	if len(inputStr) == 0 {
		platCount = 3
	} else {
		platCount, err = strconv.Atoi(string(inputStr))
		if err != nil {
			panic("invalid input")
		}
		if platCount < 1 {
			panic("invalid input, at least 1 plate needed")
		}
	}

	game := NewGame(platCount, towerCount)

	for {
		clearScreen()

		game.print()
		// game.debug()

		pln("")
		if lastMoveStatus != "" {
			pln(lastMoveStatus, ",enter 0 to exit")
		}

		gameOver := game.over()
		if gameOver {
			pf("Game over, Total moves %d !!!\n", totalMove)
			break
		}

		fmt.Print(`Enter "from" tower no:`)
		input, _, _ := console.ReadLine()
		fromChoice, err := strconv.Atoi(string(input))
		if err != nil {
			lastMoveStatus = "invalid input"
			continue
		}
		if fromChoice == 0 {
			pf("Game exit, Total moves %d !!!\n", totalMove)
			break
		}

		fmt.Print(`Enter "to" tower no:`)
		input, _, _ = console.ReadLine()
		toChoice, err := strconv.Atoi(string(input))
		if err != nil {
			lastMoveStatus = "invalid input"
			continue
		}
		if toChoice == 0 {
			pf("Game exit, Total moves %d !!!\n", totalMove)
			break
		}

		pf("from:%d, to:%d =>", fromChoice, toChoice)

		moved := game.move(fromChoice, toChoice)
		if moved {
			lastMoveStatus = fmt.Sprintf(`from:%d, to:%d => moved`, fromChoice, toChoice)
		} else {
			lastMoveStatus = fmt.Sprintf(`from:%d, to:%d => move not allowed`, fromChoice, toChoice)
		}

		totalMove++
	}
}

func clearScreen() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}

func pattern(ch byte, times int) string {
	p := ""
	for i := 1; i <= times; i++ {
		p += string(ch)
	}
	return p
}

func p(str string) {
	fmt.Print(str)
}

func pln(str ...any) {
	fmt.Println(str...)
}

func pf(format string, a ...any) {
	fmt.Printf(format, a...)
}
