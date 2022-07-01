package main

import (
	"crypto/md5"
	"image"
	"image/color"
	"image/png"
	"os"
)

//estructura de puntos de la grilla a rellenar
type GridPoint struct {
	value byte
	index int
}

//estructura de coordenadas de puntos de grilla a rellenar
type Point struct {
	x, y int
}
type DrawingPoints struct {
	topLeft     Point
	bottomRight Point
}

//estructura del identicon
type Identicon struct {
	name       string
	hash       [16]byte
	color      [3]byte
	grid       []byte
	gridPoints []GridPoint
	pixelMap   []DrawingPoints
}

//creacion del hash a partir de un dato del usuario (input)
func hashInput(input []byte) Identicon {
	checkSum := md5.Sum(input)
	return Identicon{
		name: string(input),
		hash: checkSum,
	}
}

//eleccion del color del identicon mediante los primeros 3 digitos del hash
func pickColor(identicon Identicon) Identicon {
	rgb := [3]byte{}
	copy(rgb[:], identicon.hash[0:3])
	identicon.color = rgb
	return identicon
}

//identicon puede ser simetrico de modo que hacemos un 5x5
//donde las primeras dos columnas se compian a las dos ultimas
func buildGrid(identicon Identicon) Identicon {
	grid := []byte{}
	for i := 0; i < len(identicon.hash) && i+3 <= len(identicon.hash)-1; i += 3 {
		chunk := make([]byte, 5)
		copy(chunk[:], identicon.hash[i:i+3])
		chunk[3] = chunk[1]
		chunk[4] = chunk[2]
		grid = append(grid, chunk...)
	}
	identicon.grid = grid
	return identicon
}

//hallar los valores pares del hash que van a rellenarse
func filterOddSquares(identicon Identicon) Identicon {
	grid := []GridPoint{}
	for i, code := range identicon.grid {
		if code%2 == 0 {
			point := GridPoint{
				value: code,
				index: i,
			}
			grid = append(grid, point)
		}
	}
	identicon.gridPoints = grid
	return identicon
}

//obtencion de coordenadas de cada valor par
func buildPixelMap(identicon Identicon) Identicon {
	drawingPoints := []DrawingPoints{}
	pixelFunc := func(p GridPoint) DrawingPoints {
		horizontal := (p.index % 5) * 50
		vertical := (p.index / 5) * 50
		topLeft := Point{horizontal, vertical}
		bottomRight := Point{horizontal + 50, vertical + 50}

		return DrawingPoints{
			topLeft,
			bottomRight,
		}
	}
	for _, gridPoint := range identicon.gridPoints {
		drawingPoints = append(drawingPoints, pixelFunc(gridPoint))
	}
	identicon.pixelMap = drawingPoints
	return identicon
}

func DrawImage(identicon Identicon) error {
	img := image.NewRGBA(image.Rect(0, 0, 250, 250))
	col := color.RGBA{identicon.color[0], identicon.color[1], identicon.color[2], 255}
	for _, pixel := range identicon.pixelMap {
		for x := pixel.topLeft.x; x < pixel.bottomRight.x; x++ {
			for y := pixel.topLeft.y; y < pixel.bottomRight.y; y++ {
				img.Set(x, y, col)
			}
		}
	}
	f, _ := os.Create("image.png")
	return png.Encode(f, img)
}
