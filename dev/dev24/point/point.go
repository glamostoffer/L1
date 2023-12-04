package point

type Point struct {
	x, y float64
}

func NewPoint(x, y float64) Point {
	return Point{
		x: x,
		y: y,
	}
}

func (p *Point) GetCoordinates() (float64, float64) {
	return p.x, p.y
}
