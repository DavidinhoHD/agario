package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"math/rand"

	"math"
	"strconv"
)

type Player struct {
	radius      float32
	position    rl.Vector2
	speed       float32
	lastDrawPos rl.Vector2
	//direction rl.Vector2
}

type Food struct {
	radius     float32
	posX, posY int32
}

func (p *Player) Move(mousPos rl.Vector2, speed float32) {
	//var movementPrecision float32 = 8

	// get the direction vector from player position - mouse position
	direction := subVector2(mousPos, p.position)

	// calculate length of the direction vector
	length := float32(math.Sqrt(float64(direction.X*direction.X + direction.Y*direction.Y)))

	// get the unit vector
	if length != 0 {
		direction.X = direction.X / length
		direction.Y = direction.Y / length

		// add the unit vector to the current player position times the sped of the player
		p.position.X += direction.X * speed
		p.position.Y += direction.Y * speed

		//p.position.X = float32(math.Round(float64(newX*10))) / 10
		//p.position.Y = float32(math.Round(float64(newY*10))) / 10

	}
}

func spawnFood() Food {
	var f Food
	x := rand.Int31n(1920)
	y := rand.Int31n(1080)

	f.posX = x
	f.posY = y
	f.radius = 10.0

	return f
}

func main() {
	rl.InitWindow(1920, 1080, "window title")
	defer rl.CloseWindow()

	// rl.KeyNull = 0
	rl.SetExitKey(rl.KeyNull)
	rl.SetTargetFPS(60)

	player := Player{
		radius:   30.0,
		position: rl.Vector2{X: 1920 / 2, Y: 1080 / 2},
		speed:    15,
	}

	foods := make([]Food, 0)
	foods = append(foods, spawnFood())
	foods = append(foods, spawnFood())
	foods = append(foods, spawnFood())
	foods = append(foods, spawnFood())
	foods = append(foods, spawnFood())
	foods = append(foods, spawnFood())

	player.lastDrawPos = player.position

	// gameLoop
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		cursor := rl.IsCursorOnScreen()

		for _, v := range foods {
			rl.DrawCircle(v.posX, v.posY, v.radius, rl.Green)
		}

		const drawSmoothing = 0.1
		newDrawX := player.position.X*drawSmoothing + player.lastDrawPos.X*(1-drawSmoothing)
		newDrawY := player.position.Y*drawSmoothing + player.lastDrawPos.Y*(1-drawSmoothing)

		drawX := int32(math.Round(float64(newDrawX)))
		drawY := int32(math.Round(float64(newDrawY)))
		rl.DrawCircle(drawX, drawY, player.radius, rl.Pink)

		player.lastDrawPos = rl.Vector2{X: newDrawX, Y: newDrawY}

		//rl.DrawCircle(int32(player.position.X), int32(player.position.Y), player.radius, rl.Pink)

		if cursor {
			mousePosition := rl.GetMousePosition()
			player.Move(mousePosition, player.speed)
		}

		fps := rl.GetFPS()
		s := strconv.FormatFloat(float64(fps), 'f', 0, 32)

		rl.DrawText(s, 5, 5, 30, rl.Green)
		rl.EndDrawing()
	}
}
