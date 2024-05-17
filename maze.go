package main

import (
	"fmt"
	"math"
	"math/rand/v2"
	"slices"
	"unicode/utf8"
)

func main() {
	maze := setupMaze(10, 5, 15)
	print(&maze, nil)

	start := generatePoint(&maze)
	end := generatePoint(&maze)
	path := shortestPath(&maze, start, end)
	print(&maze, &path)
}

type Point struct {
	x          int
	y          int
	wall       bool
	distance   int
	neighbours []*Point
	prev       *Point
}

func shortestPath(maze *[][]*Point, start *Point, end *Point) []*Point {
	// TODO
	// var q []*Point
	return nil
}

func setupMaze(sizeX int, sizeY int, walls int) [][]*Point {
	var maze [][]*Point

	for j := 0; j < sizeY; j++ {
		var row []*Point

		for i := 0; i < sizeX; i++ {
			newPoint := Point{
				x:          i,
				y:          j,
				wall:       false,
				distance:   sizeX * sizeY,
				neighbours: nil,
				prev:       nil,
			}
			row = append(row, &newPoint)
		}
		maze = append(maze, row)
	}

	setWalls(&maze, walls)
	setNeighbours(&maze)

	return maze
}

func setWalls(maze *[][]*Point, walls int) {
	var flatMaze []*Point
	for _, row := range *maze {
		flatMaze = append(flatMaze, row...)
	}

	rand.Shuffle(len(flatMaze), func(i, j int) {
		flatMaze[i], flatMaze[j] = flatMaze[j], flatMaze[i]
	})

	for i := 0; i < walls; i++ {
		flatMaze[i].wall = true
	}
}

func setNeighbours(maze *[][]*Point) {
	for j, row := range *maze {
		for i, _ := range row {
			point := (*maze)[j][i]
			point.neighbours = getPointNeighbours(point, maze)
		}
	}
}

func getPointNeighbours(p *Point, maze *[][]*Point) []*Point {
	sizeX := float64(len((*maze)[0]))
	sizeY := float64(len(*maze))
	var neighbours []*Point

	for row := math.Max(float64(p.x-1), 0); row < math.Min(float64(p.x+2), sizeX); row++ {
		for col := math.Max(float64(p.y-1), 0); col < math.Min(float64(p.y+2), sizeY); col++ {
			point := (*maze)[int(col)][int(row)]
			if !point.wall && point != p && (int(row) == p.x || int(col) == p.y) {
				neighbours = append(neighbours, point)
			}
		}
	}
	return neighbours
}

func print(maze *[][]*Point, path *[]*Point) {
	for j := 0; j < len(*maze); j++ {
		row := ""
		for i := 0; i < len((*maze)[0]); i++ {
			point := (*maze)[j][i]
			icon, _ := utf8.DecodeRune([]byte{0xE2, 0xAC, 0x9C})
			if point.wall {
				icon, _ = utf8.DecodeRune([]byte{0xE2, 0xAC, 0x9B})
			}
			if path != nil {
				if slices.Contains(*path, point) {
					icon, _ = utf8.DecodeRune([]byte{0xE2, 0xAD, 0x95})
				}
				if point == (*path)[0] {
					icon, _ = utf8.DecodeRune([]byte{0xE2, 0xAD, 0x90})
				}
				if point == (*path)[len(*path)-1] {
					icon, _ = utf8.DecodeRune([]byte{0xF0, 0x9F, 0x8E, 0xAF})
				}
			}
			row += string(icon)
		}
		fmt.Println(row)
	}
}

func generatePoint(maze *[][]*Point) *Point {
	sizeX := len((*maze)[0])
	sizeY := len(*maze)
	var point *Point
	for {
		point = (*maze)[rand.IntN(sizeY)][rand.IntN(sizeX)]
		if !point.wall {
			break
		}
	}
	return point
}
