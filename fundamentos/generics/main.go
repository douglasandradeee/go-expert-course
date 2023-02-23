package main

type MyNumber int64

// Number - é uma constraint
type Number interface {
	~int64 | ~float64
}

// Usando generics na criação da função de soma para que a mesma seja agnóstica ao tipo de dado passado.
func Soma[T Number](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

// comparable implementa uma interface de tipos comparáveis, ou seja, se o tipos não forem comparáveis a função não funciona, vai dar error.
// O comparable compara apenas a igualdade de alguma coisa. Ele não consegue resolver comparações onde algo seja maior que ou menor que.
func Compara[T comparable](a T, b T) bool {
	if a == b {
		return true
	}
	return false
}

func main() {

	m := map[string]int64{
		"Douglas": 10000,
		"Gabi":    900000,
		"João":    70000,
		"Maria":   3000}

	m2 := map[string]float64{
		"Douglas": 1050.22,
		"Gabi":    9000.99,
		"João":    7001.33,
		"Maria":   3000.33}

	m3 := map[string]MyNumber{
		"Douglas": 10000,
		"Gabi":    900000,
		"João":    70000}

	println(Soma(m))
	println(Soma(m2))
	println(Soma(m3))
	println(Compara(10, 10))
}
