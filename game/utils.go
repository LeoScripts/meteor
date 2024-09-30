package game

const (
	screamWith = 800
	screamHeith = 600
)

type Vector struct {
	X float64
	Y float64
}

type Rect struct {
	X float64
	Y float64
	Width float64
	Heidth float64
}

func NewRect(x, y, width, heidth float64) Rect {
	return Rect{
		X: x,
		Y: y,
		Width: width,
		Heidth: heidth,
	}
} 

// verifica se um retangulo estao em cima do outro
func (r Rect) Intersects(other Rect) bool {
	return r.X <= other.maxX() &&
	other.X <= r.maxX() &&
	r.Y <= other.maxY() &&
	other.maxY() <= r.maxY()

}
// largura das imagens
func (r Rect) maxX() float64 {
	return r.X + r.Width
}
// altura das imagens
func (r Rect) maxY() float64 {
	return r.Y + r.Heidth
}