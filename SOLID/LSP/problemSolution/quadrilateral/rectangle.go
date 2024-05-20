package quadrilateral

type Rectangle struct {
	Area
	Width  float64
	Height float64
}

type RectangleInterface interface {
	GetWidth() float64
	GetHeight() float64

	SetWidth(width float64)
	SetHeight(height float64)
}

func (r *Rectangle) GetArea() float64 {
	return r.Width * r.Height
}

func (r *Rectangle) GetWidth() float64 {
	return r.Width
}

func (r *Rectangle) GetHeight() float64 {
	return r.Height
}

func (r *Rectangle) SetWidth(width float64) {
	r.Width = width
}

func (r *Rectangle) SetHeight(height float64) {
	r.Height = height
}
