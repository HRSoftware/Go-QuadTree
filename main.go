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

	rootQuadTree := types.MakeQuadTree(rl.Rectangle{
		X:      0,
		Y:      0,
		Width:  float32(WIDTH),
		Height: float32(HEIGHT),
	})

	var rectPositions []rl.Vector2

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
			rectPositions = append(rectPositions, rl.Vector2{X: newRec.X, Y: newRec.Y})
			rootQuadTree.Insert(rl.Vector2{X: newRec.X, Y: newRec.Y})
		}

		// render the quadtree
		for _, val := range rootQuadTree.Visualise() {
			rl.DrawRectangleV(rl.Vector2{X: val.Rect.X, Y: val.Rect.Y}, rl.Vector2{X: val.Rect.Width, Y: val.Rect.Height}, val.Col)
		}

		// draw where the mouse is
		rl.DrawRectangleV(mousePos, rl.Vector2{X: 20., Y: 20.}, rl.Red)

		queryResults := rootQuadTree.Query(rl.Rectangle{X: mousePos.X, Y: mousePos.Y, Width: RECT_SIZE.X, Height: RECT_SIZE.Y})

		for _, rect := range rectPositions {
			_, err := utilities.LinearSearch(queryResults[:], rect)
			if err != nil {
				rl.DrawRectangleV(rect, RECT_SIZE, rl.White)
			} else {
				rl.DrawRectangleV(rect, RECT_SIZE, rl.Green)
			}
		}

		rl.DrawFPS(10, 10)
		rl.EndDrawing()
	}
}
