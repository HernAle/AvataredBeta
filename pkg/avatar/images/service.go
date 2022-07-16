package images

//estructura principal del avatar
type Avatar struct {
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
func (avatar *Avatar) getColor(encodedInformation []byte) [3]byte {
	rgb := [3]byte{}
	copy(rgb[:], encodedInformation[0:3])
	avatar.color = rgb
	return avatar.color
}

//metodo para crear un avatar simetrico, las dos primeras columna se reflejan en las dos ultimas
func (avatar *Avatar) buildGrid(encodedInformation []byte) []byte {
	grid := []byte{}
	for i := 0; i < len(encodedInformation) && i+3 <= len(encodedInformation)-1; i += 3 {
		box := make([]byte, 5)
		copy(box[:], encodedInformation[i:i+3])
		box[3] = box[1]
		box[4] = box[2]
		grid = append(grid, box...)
	}
	avatar.grid = grid
	return avatar.grid
}

//metodo para encontrar los valores pares del hash a rellenar
func (avatar *Avatar) findOddBoxes([]GridPoint) []GridPoint {
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
