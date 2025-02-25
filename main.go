package Go_QuadTree

import (
	"fmt"
	_ "fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Vec2f = [2]float32

type QuadTree struct {
	boundary  rl.Rectangle
	points    []Vec2f
	northWest *QuadTree
	northEast *QuadTree
	southWest *QuadTree
	southEast *QuadTree
	// allocator mem.Allocator
}

func makeQuadTree(aabb rl.Rectangle) *QuadTree {
	return &QuadTree{
		boundary:  aabb,
		points:    nil,
		northWest: nil,
		northEast: nil,
		southWest: nil,
		southEast: nil,
	}
}

func destroyQuadTree(tree *QuadTree) {
	if tree.northWest == nil || tree.northEast == nil || tree.southWest == nil || tree.southEast == nil {
		return
	}

	destroyQuadTree(tree.northWest)
	destroyQuadTree(tree.northEast)
	destroyQuadTree(tree.southWest)
	destroyQuadTree(tree.southEast)
}

func insertQuadTree(tree *QuadTree, point Vec2f) bool {
	if !rl.CheckCollisionPointRec(
		rl.Vector2{X: point[0], Y: point[1]},
		tree.boundary,
	) {
		return false
	}

	if len(tree.points) < 4 || tree.northWest == nil {

		if tree.points == nil {
			tree.points = make([]Vec2f, 0)
		}

		tree.points = append(tree.points, point)
		fmt.Println("Appending point")
		return true
	}

	if tree.northEast == nil {
		fmt.Println("Subdividing")
		subdivideQuadTree(tree)
	}

	if insertQuadTree(tree.northWest, point) {
		return true
	}
	if insertQuadTree(tree.northEast, point) {
		return true
	}
	if insertQuadTree(tree.southWest, point) {
		return true
	}
	if insertQuadTree(tree.southEast, point) {
		return true
	}

	panic("Unable to subdivide any further, and unable to insert point into quad tree. Consider adjusting qtree parameters")
}

func subdivideQuadTree(tree *QuadTree) {
	halfBounds := Vec2f{tree.boundary.Width / 2, tree.boundary.Height / 2}

	tree.northWest = makeQuadTree(rl.Rectangle{X: tree.boundary.X, Y: tree.boundary.Y, Width: halfBounds[0], Height: halfBounds[1]})
	tree.northEast = makeQuadTree(rl.Rectangle{X: tree.boundary.X + halfBounds[0], Y: tree.boundary.Y, Width: halfBounds[0], Height: halfBounds[1]})
	tree.southWest = makeQuadTree(rl.Rectangle{X: tree.boundary.X, Y: tree.boundary.Y + halfBounds[1], Width: halfBounds[0], Height: halfBounds[1]})
	tree.southEast = makeQuadTree(rl.Rectangle{X: tree.boundary.X + halfBounds[0], Y: tree.boundary.Y + halfBounds[1], Width: halfBounds[0], Height: halfBounds[1]})

}

func queryQuadTree(tree *QuadTree, rec rl.Rectangle) []Vec2f {
	results := make([]Vec2f, 0)

	if !rl.CheckCollisionRecs(tree.boundary, rec) {
		return results
	}

	for _, point := range tree.points {
		if rl.CheckCollisionPointRec(rl.Vector2{X: point[0], Y: point[1]}, rec) {
			_ = append(
				results,
				point,
			)
		}
	}

	if tree.northWest == nil {
		return results
	}

	if childResults := queryQuadTree(tree.northWest, rec); len(childResults) > 0 {
		for _, point := range childResults {
			_ = append(
				results,
				point,
			)
		}
	}

	if childResults := queryQuadTree(tree.northEast, rec); len(childResults) > 0 {
		for _, point := range childResults {
			_ = append(
				results,
				point,
			)
		}
	}

	if childResults := queryQuadTree(tree.southEast, rec); len(childResults) > 0 {
		for _, point := range childResults {
			_ = append(
				results,
				point,
			)
		}
	}

	if childResults := queryQuadTree(tree.southWest, rec); len(childResults) > 0 {
		for _, point := range childResults {
			_ = append(
				results,
				point,
			)
		}
	}

	return results
}
