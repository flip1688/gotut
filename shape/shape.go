package shape

var (
	Total     = 10
	number    = 15
	errorText = "error"
)

// Triangle type
type Triangle struct {
	Base   float32
	height float32
}

// Square type
type Square struct {
	Width float32
}

// Rectangle type
type Rectangle struct {
	Width  float32
	Height float32
}

const (
	shapName = "Circle"
)

func (r Rectangle) Area() float32 {
	return r.Height * r.Width
}
func (t Triangle) Area() float32 {
	return t.Base * t.height * 0.5
}
func (t *Triangle) SetHeight(newHeight float32) {
	t.height = newHeight
}

// Distance type
type Distance float32

type HelperFunc func() error

type Shape interface {
	Area() float32
}
