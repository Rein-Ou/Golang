package rectangle

type rectangle struct {
	length int
	width  int
}

func Area(r rectangle) int {
	return r.width * r.length
}

func Primeter(r rectangle) int {
	return (r.length + r.width) * 2
}
func NewRectangleEmpty() *rectangle {
	return new(rectangle)
}
func NewRectangle(l int, w int) *rectangle {
	return &rectangle{l, w}
}

func (r *rectangle) GetWidth() int {
	return r.width
}

func (r *rectangle) GetLength() int {
	return r.length
}

func (r *rectangle) SetWidth(w int) {
	r.width = w
}

func (r *rectangle) SetLength(l int) {
	r.length = l
}
