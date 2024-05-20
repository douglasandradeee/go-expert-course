package quadrilateral

type Rectangle struct {
	Area
	width  float64
	height float64
}

type RectangleInterface interface {
	GetWidth() float64
	GetHeight() float64

	setWidth(width float64)
	setHeight(height float64)
}

func (r *Rectangle) GetArea() float64 {
	return r.width * r.height
}

func (r *Rectangle) GetWidth() float64 {
	return r.width
}

func (r *Rectangle) GetHeight() float64 {
	return r.height
}

func (r *Rectangle) SetWidth(width float64) {
	r.width = width
}

func (r *Rectangle) SetHeight(height float64) {
	r.height = height
}
