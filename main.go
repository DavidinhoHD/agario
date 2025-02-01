package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"

	"math"
	"strconv"
)

type Player struct {
	radius    float32
	position  rl.Vector2
	speed     float32
	direction rl.Vector2
}

func (p *Player) Move(mousPos rl.Vector2, speed float32) {
	// get the direction vector from player position - mouse position
	p.direction = subVector2(mousPos, p.position)

	// calculate length of the direction vector
	length := float32(math.Sqrt(float64(p.direction.X*p.direction.X + p.direction.Y*p.direction.Y)))

	// get the unit vector
	p.direction.X = p.direction.X / length
	p.direction.Y = p.direction.Y / length

	// add the unit vector to the current player position times the sped of the player
	p.position.X += p.direction.X * speed
	p.position.Y += p.direction.Y * speed
}

type Food struct {
	radius     float32
	posX, posY int32
}

func main() {
	rl.InitWindow(1920, 1080, "window title")
	defer rl.CloseWindow()

	// rl.KeyNull = 0
	rl.SetExitKey(rl.KeyNull)
	rl.SetTargetFPS(60)

	player := Player{
		radius:    30.0,
		position:  rl.Vector2{X: 1920 / 2, Y: 1080 / 2},
		speed:     15,
		direction: rl.Vector2{X: 0, Y: 0},
	}
	/*
		food := Food{
			radius: 10.0,
			posX:   300,
			posY:   300,
		}
	*/

	foods := make([]Food, 0)
	foods = append(foods, spawnFood())
	foods = append(foods, spawnFood())
	foods = append(foods, spawnFood())
	foods = append(foods, spawnFood())
	foods = append(foods, spawnFood())
	foods = append(foods, spawnFood())

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		cursor := rl.IsCursorOnScreen()

		for _, v := range foods {
			rl.DrawCircle(v.posX, v.posY, v.radius, rl.Green)
		}

		rl.DrawCircle(int32(player.position.X), int32(player.position.Y), player.radius, rl.Pink)

		if cursor {
			mousePosition := rl.GetMousePosition()
			player.Move(mousePosition, player.speed)
		}

		fps := rl.GetFPS()
		s := strconv.FormatFloat(float64(fps), 'f', 0, 32)

		rl.DrawText(s, 5, 5, 30, rl.Green)
		rl.ClearBackground(rl.RayWhite)
		rl.EndDrawing()
	}
}
