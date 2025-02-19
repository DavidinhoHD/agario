package game
import (
	rl "github.com/gen2brain/raylib-go/raylib"

)

type Food struct {
	ID			uint64 		`json:"id"`
	Position 	rl.Vector2	`json:"position"`	
	Radius		float32		`json:"-"`
}


func (f *Food) SpawnFood() {
	for _,v := range GS.FoodPositions{
		rl.DrawCircle(int32(v.Position.X), int32(v.Position.Y), v.Radius, rl.Green)
	}
}