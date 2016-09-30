package main

import (
	"github.com/ajhager/engi"
	"time"
//	"fmt"
)

var (
	Scale = float32(1)
	Size = 250
	Swich = false
	)


type Position struct {
	x float32
	y float32
}

type Map struct {
	data [][]bool
	info [][]int
}

type Game struct {
	*engi.Game
	bot   			engi.Drawable
	batch 			*engi.Batch
	font  			*engi.Font
	Map	  			Map
	play  			bool
	pointDisplay 	Position
	touch			bool
}

func (game *Game) Preload() {
	engi.Files.Add("bot", "data/blub.png")
	engi.Files.Add("font", "data/font.png")
		
	midx := engi.Width() / 8
	midy := engi.Height() / 8
	game.touch = false
	game.pointDisplay = Position{midx, midy}


	game.Map.data = make([][]bool, Size)
	for i, _ := range game.Map.data {
		game.Map.data[i] = make([]bool, Size)
	}
	
	game.Map.info = make([][]int, Size)
	for i, _ := range game.Map.info {
		game.Map.info[i] = make([]int, Size)
	}
	game.play = false

	game.Map.data[Size / 2][Size / 2] = true
	game.Map.data[Size / 2][Size / 2 + 1] = true
	game.Map.data[Size / 2][Size / 2 - 1] = true
	game.Map.data[Size / 2 + 1][Size / 2] = true
	game.Map.data[Size / 2 + 2][Size / 2] = true

	
	go game.Fuck()

}

func (game *Game) Fuck() {
	game.Map.data[Size / 2][Size / 2] = true
	game.Map.data[Size / 2][Size / 2 + 1] = true
	game.Map.data[Size / 2][Size / 2 - 1] = true
	game.Map.data[Size / 2 + 1][Size / 2] = true
	game.Map.data[Size / 2 + 2][Size / 2] = true
	time.Sleep(1 * time.Second)
	game.Fuck()
}

func (game *Game) Update(delta float32) {
	if (game.play) {
		game.AlgoGameOfLife()
		if (Swich) {
			time.Sleep(200 * time.Millisecond)
		}
		
	}
}

func getNbLiving(data [][]bool, x, y int) int {
	nb := 0
	if y < Size - 1 && data[y + 1][x] {
		nb += 1
	}
	if y < Size - 1 && x < Size - 1 && data[y + 1][x + 1] {
		nb += 1
	}
	if x < Size - 1 && data[y][x + 1] {
		nb += 1
	}
	if y > 0 && x < Size - 1 && data[y - 1][x + 1] {
		nb += 1
	}
	if y > 0 && data[y - 1][x] {
		nb += 1
	}
	if y > 0 && x > 0 && data[y - 1][x - 1] {
		nb += 1
	}
	if x > 0 && data[y][x - 1] {
		nb += 1
	}
	if y < Size - 1 && x > 0 && data[y + 1][x - 1] {
		nb += 1
	}
	return nb
}


func (game *Game) AlgoGameOfLife() {
	for y, value := range game.Map.data {
		for x := range value {
			game.Map.info[y][x] = 0
			game.Map.info[y][x] = getNbLiving(game.Map.data, x, y)
		}
	}

	for y, value := range game.Map.data {
		for x, state := range value {
			if state {
				if !(game.Map.info[y][x] == 2 || game.Map.info[y][x] == 3 || (Swich && game.Map.info[y][x] == 4)) {
					game.Map.data[y][x] = false
				}
			} else {
				if game.Map.info[y][x] == 3 || (Swich && game.Map.info[y][x] == 2) {
					game.Map.data[y][x] = true
				}
			}
		}
	}

}

func (game *Game) Setup() {
	engi.SetBg(0x362d38)
	game.bot = engi.Files.Image("bot")
	game.font = engi.NewGridFont(engi.Files.Image("font"), 20, 20)
	game.batch = engi.NewBatch(engi.Width(), engi.Height())
}

func (game *Game) Render() {
	game.batch.Begin()

	
	
	for y, value := range game.Map.data {
		for x := range value {
			if value[x] || x == 0 || x == Size - 1 || y == 0 || y == Size - 1{
				game.batch.Draw(game.bot, game.pointDisplay.x + float32(x * 16) * Scale, game.pointDisplay.y + float32(y * 16) * Scale, 0.5, 0.5, Scale, Scale, 0, 0xffffff, 1)
			}
		}
	}
	game.font.Print(game.batch, "Game of Life", game.pointDisplay.x, game.pointDisplay.y, 0xffffff)
	game.batch.End()

}



func (game *Game) Resize(w, h float32) {
	game.batch.SetProjection(w, h)
}

func (game *Game) Mouse(x, y float32, action engi.Action) {
	if (action == engi.PRESS) {
		game.touch = true
	}
	if (action == engi.RELEASE) {
		game.touch = false
	}

	if (action == engi.PRESS || (action == engi.MOVE && game.touch)) {

		caseX := int((x - game.pointDisplay.x) / 16 / Scale)
		caseY := int((y - game.pointDisplay.y) / 16 / Scale)

		if (caseX > 0 && caseX < Size - 1 && caseY > 0 && caseY < Size - 1) {
			if (action == engi.MOVE) {
				game.Map.data[caseY][caseX] = true
			} else {
				game.Map.data[caseY][caseX] = !game.Map.data[caseY][caseX]
			}
			
		}

	}
}

func (game *Game) Key(key engi.Key, modifier engi.Modifier, action engi.Action) {
	if (action == engi.PRESS || action == engi.REPEAT) {
		if (key == engi.Space) {
			game.play = !game.play
		} else if key == engi.S {
			Swich = !Swich
		}
	}

}

func (g *Game) Scroll(amount float32) {
	Scale += amount * .01
}

func main() {

	game := Game{}

	engi.Open("Hello", 800, 600, false, &game)
}