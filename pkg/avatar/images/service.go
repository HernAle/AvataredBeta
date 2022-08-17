package images

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

//estructura principal del avatar
type Avatar struct {
	hash       []byte
	color      [3]byte
	grid       []byte
	gridPoints []GridPoint
	pixelMap   []DrawingPoints
}

//estructura de casillas de la tabla a rellenar de color
type GridPoint struct {
	value byte
	index int
}

//estructura embebida de coordenadas de las casillas
type DrawingPoints struct {
	topLeft     Point
	bottomRight Point
}
type Point struct {
	x, y int
}

//metodo para la obtencion del color del avatar mediante los 3 primeros valores del hash
func (avatar *Avatar) getColor() [3]byte {
	rgb := [3]byte{}
	copy(rgb[:], avatar.hash[0:3])
	avatar.color = rgb
	return avatar.color
}

//metodo para crear un avatar simetrico, las dos primeras columna se reflejan en las dos ultimas
func (avatar *Avatar) buildGrid() []byte {
	grid := []byte{}
	for i := 0; i < len(avatar.hash) && i+3 <= len(avatar.hash)-1; i += 3 {
		box := make([]byte, 5)
		copy(box[:], avatar.hash[i:i+3])
		box[3] = box[1]
		box[4] = box[0]
		grid = append(grid, box...)
	}
	avatar.grid = grid
	return avatar.grid
}

//metodo para encontrar los valores pares del hash a rellenar
func (avatar *Avatar) findOddBoxes() []GridPoint {
	grid := []GridPoint{}
	for i, code := range avatar.grid {
		if code%2 == 0 {
			point := GridPoint{
				value: code,
				index: i,
			}
			grid = append(grid, point)
		}
	}
	avatar.gridPoints = grid
	return avatar.gridPoints
}

//metodo para obtener las coordenadas de cada valor par
func (avatar *Avatar) buildPixelMap() []DrawingPoints {
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
	for _, gridPoint := range avatar.gridPoints {
		drawingPoints = append(drawingPoints, pixelFunc(gridPoint))
	}
	avatar.pixelMap = drawingPoints
	return avatar.pixelMap
}

//metodo para construccion de avatar
func (avatar *Avatar) buildImage() error {
	img := image.NewRGBA(image.Rect(0, 0, 250, 250))
	col := color.RGBA{avatar.color[0], avatar.color[1], avatar.color[2], 255}
	for _, pixel := range avatar.pixelMap {
		for x := pixel.topLeft.x; x < pixel.bottomRight.x; x++ {
			for y := pixel.topLeft.y; y < pixel.bottomRight.y; y++ {
				img.Set(x, y, col)
			}
		}
	}
	f, _ := os.Create("your-avatar.png")
	return png.Encode(f, img)
}

func (avatar *Avatar) BuildAndSaveImage(encodedInformation []byte) error {
	avatar.hash = encodedInformation
	avatar.getColor()
	avatar.buildGrid()
	avatar.findOddBoxes()
	avatar.buildPixelMap()
	avatar.buildImage()
	return nil
}

func NewAvatar() *Avatar {
	return &Avatar{}
}
