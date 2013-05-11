package main

import (
	"fmt"
	"log"
	"os"

	qt "github.com/foolusion/quadtree"
)

type point struct {
	id int
	XY
}

func main() {
	points := ReadPointsFromFile(file)
	bounds := BoundingBox(points)
	root := qt.New(bounds)

	InsertAllPoints(root, points)

	results := NearestNeighbors(root)
}

func ReadPointsFromFile(fn string) []point {
	file, err := os.Open(fn)
	if err != nil {
		log.Fatalf(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	points := make([]*XY, 100)

	done := false

	for !done {
		line, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalf(err)
		}

		ptString := strings.Split(line, '\t')

		id, err := strconv.Atoi(ptString[0])
		if err != nil {
			log.Fatalf(err)
		}

		x, err := strconv.ParseFloat(ptString[1], 64)
		if err != nil {
			log.Fatalf(err)
		}
		y, err := strconv.ParseFloat(ptString[2], 64)
		if err != nil {
			log.Fatalf(err)
		}

		pt := &point{id, XY{x, y}}

		points = append(points, pt)
	}

	return points
}
