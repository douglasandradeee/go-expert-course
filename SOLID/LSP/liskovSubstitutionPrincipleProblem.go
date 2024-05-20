package lsp

type quadrilateral interface {
	GetWidth() float64
	GetHeight() float64

	setWidth(width float64)
	setHeight(height float64)
}

type Square struct {
	width  float64
	height float64
}

type Rectangle struct {
	width  float64
	height float64
}

func (s *Square) GetArea() float64 {
	//devemos usar height ou width aqui?
	// lembrando que no quadrado todos os lados são iguais.
	return s.height * s.height
}

func (s *Square) GetWidth() float64 {
	return s.width
}

func (s *Square) GetHeight() float64 {
	return s.height
}

func (s *Square) SetWidth(width float64) {
	// os dois lados tem de se manter com o mesmo valor, portanto ao setar o width
	// estaríamos setando tbm o height
	s.width = width
	s.height = width
}

func (s *Square) SetHeight(height float64) {
	// os dois lados tem de se manter com o mesmo valor, portanto ao setar o height
	// estaríamos setando tbm o width
	s.height = height
	s.width = height
}

func (r *Rectangle) GetArea() float64 {
	return r.width * r.width
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
