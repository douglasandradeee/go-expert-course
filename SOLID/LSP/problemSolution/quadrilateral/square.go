package quadrilateral

type SquareIterface interface {
	Area
	GetSize() float64
	SetSize(size float64)
}

type Square struct {
	size float64
}

func (s *Square) GetArea() float64 {
	return s.size * s.size
}

func (s *Square) GetSize() float64 {
	return s.size
}

func (s *Square) SetSize(size float64) {
	s.size = size
}
