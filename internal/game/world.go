package game

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	//"fmt"
)

const (
	MinSpeed = 5.0
	GrowthRate = 1.0
	SpeedDecay = 0.99
	BackgroundImagePath = "assets/background.png"

)

var Points int = 0



var WorldDimensions = rl.Vector2{
		X: 1920 * 3,
		Y: 1080 * 3,
}
var WindowDimensions = rl.Vector2{
		X: 1920,
		Y: 1080,
}


type GameState struct {
	FoodPositions []Food		`json:"FoodsPositions"`
	PlayerPositions []Player	`json:"PlayerPositions"`
}

func (g *GameState) SpawnFood() {
	for _,v := range g.FoodPositions{
		rl.DrawCircle(int32(v.Position.X), int32(v.Position.Y), v.Radius, rl.Green)
		//fmt.Println(v.Radius)
	}
}

func (g *GameState) SpawnPlayers(p Player) {
	for _,v := range g.PlayerPositions{
		if v.ID != p.ID {
			rl.DrawCircle(int32(v.Position.X), int32(v.Position.Y), v.Radius, rl.Blue)
		}
		//fmt.Println(v.Radius)
	}
}



func DrawWorld(b rl.Texture2D) {
	tilesX := WorldDimensions.X / 320
	tilesY := WorldDimensions.Y / 180

	for i := int32(0); i < int32(tilesX); i++ {
		x := int32(320 * i)
		for j:= int32(0); j < int32(tilesY); j++ {
			y := int32(180 * j)
			rl.DrawTexture(b, x, y, rl.White)
		}
	}
}


func PlayerFoodCollision(f Food, p Player) bool {
	if rl.CheckCollisionCircles(f.Position, f.Radius, p.Position, p.Radius) {
		Points ++
		return true
	}
	return false
}

func PlayerPlayerCollision(p1, p2 Player) bool {
	return rl.CheckCollisionCircles(p1.Position, p1.Radius, p2.Position, p2.Radius)
}

func SetupCamera(p Player) rl.Camera2D {
	c := rl.Camera2D{
		Offset: rl.Vector2{X: WindowDimensions.X / 2, Y: WindowDimensions.Y / 2},
		Target: p.Position,
		Rotation: 0.0,
		Zoom: 1.0,
	}
	return c
}
