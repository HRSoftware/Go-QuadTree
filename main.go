package main

import (
	_ "fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
	"goQuadTree/types"
	"goQuadTree/utilities"
)

func main() {
	var WIDTH int32 = 640
	var HEIGHT int32 = 480

	RECT_SIZE := rl.Vector2{X: 10., Y: 10.}

	rl.InitWindow(WIDTH, HEIGHT, "Quadtree Test")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	quadTree := types.MakeQuadTree(rl.Rectangle{X: 0, Y: 0, Width: float32(WIDTH), Height: float32(HEIGHT)})

	rectPositions := map[rl.Vector2]rl.Color{}

	var randColours []rl.Color

	for idx, _ := range randColours {
		randColours[idx] = utilities.RandomColour()
	}

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
		mousePos := rl.GetMousePosition()

		if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
			newRec := rl.Rectangle{X: mousePos.X, Y: mousePos.Y, Width: float32(WIDTH), Height: float32(HEIGHT)}
			rectPositions[rl.Vector2{X: newRec.X, Y: newRec.Y}] = utilities.RandomColour()
			quadTree.Insert(rl.Vector2{X: newRec.X, Y: newRec.Y})
		}

		visualisation := quadTree.Visualise()

		for _, val := range visualisation {
			rl.DrawRectangleV(rl.Vector2{X: val.X, Y: val.Y}, rl.Vector2{X: val.Width, Y: val.Height}, rectPositions[rl.Vector2{X: val.X, Y: val.Y}])
		}

		rl.DrawRectangleV(mousePos, rl.Vector2{X: 20., Y: 20.}, rl.Red)

		queryResults := quadTree.Query(rl.Rectangle{X: mousePos.X, Y: mousePos.Y, Width: RECT_SIZE.X, Height: RECT_SIZE.Y})

		for key, _ := range rectPositions {
			_, err := utilities.LinearSearch(queryResults[:], key)
			if err != nil {
				rl.DrawRectangleV(key, RECT_SIZE, rl.White)
			} else {
				rl.DrawRectangleV(key, RECT_SIZE, rl.Green)
			}
		}

		rl.DrawFPS(10, 10)
		rl.EndDrawing()
	}
}
