package main

func soma(a, b *int) int {
	*a = 50
	return *a + *b
}

func main() {

	minhavar1 := 10
	minhavar2 := 20
	soma(&minhavar1, &minhavar2)
	println(minhavar1)

}
