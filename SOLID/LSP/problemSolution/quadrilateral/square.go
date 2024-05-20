package quadrilateral

type SquareInterface interface {
	Area
	GetSize() float64
	SetSize(size float64)
}

type Square struct {
	Size float64
}

func (s *Square) GetArea() float64 {
	return s.Size * s.Size
}

func (s *Square) GetSize() float64 {
	return s.Size
}

func (s *Square) SetSize(size float64) {
	s.Size = size
}
