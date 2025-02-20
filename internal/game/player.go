package game

import (
	//"math"

	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Player struct {
	ID          uint64     `json:"id"`
	Radius      float32    `json:"radius"`
	Position    rl.Vector2 `json:"position"`
	Speed       float32    `json:"speed"`
	LastDrawPos rl.Vector2 `json:"-"`
}

func (p *Player) DrawPlayer() {
	x := float32(1920 / 2)
	y := float32(1080 / 2)
	rl.DrawCircle(int32(x), int32(y), p.Radius, rl.Pink)
}

func (p *Player) Move(mousePos rl.Vector2, speed float32) {
	// get direction Vector
	direction := rl.Vector2Subtract(mousePos, p.Position)

	// get unit vector for constant speed regardless of direction length
	normalizedDirection := rl.Vector2Normalize(direction)
	// scale unit vector with speed
	newPos := rl.Vector2Add(p.Position, rl.Vector2Scale(normalizedDirection, speed))
	// update the player position and limit to world dimensions
	p.Position = rl.Vector2Clamp(newPos, rl.Vector2{X: 0, Y: 0}, WorldDimensions)

	// smooth stopping transition
	/*

		// Create a smooth transition for stopping
		const stopRadius = 15.0
		if length < stopRadius {
			speed *= length / stopRadius // Gradually reduce speed as we get closer to center
			if speed < 1 { // Minimum speed threshold
				return
			}
		}

	*/
}

func (p *Player) ChangePlayerProperties() {
	p.Radius += GrowthRate
	p.Speed = float32(math.Max(float64(p.Speed*SpeedDecay), float64(MinSpeed)))
}
