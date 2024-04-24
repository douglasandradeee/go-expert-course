package tax

import (
	"testing"
)

func TestCalculateTax(t *testing.T) {
	amount := 500.00
	expectedTax := 5.0

	result := CalculateTax(amount)
	if result != expectedTax { // if the result is not equal to the expected tax
		t.Error("For", amount, "expected", expectedTax, "got", result)
	}
}

func TestCalculateTaxBatch(t *testing.T) {
	type calctax struct {
		amount, expectedTax float64
	}

	table := []calctax{
		{500.00, 5.0},
		{1000.00, 10.0},
		{1500.00, 10.0}, // this will fail
		{0.0, 0.0},
	}

	for _, item := range table {
		result := CalculateTax(item.amount)
		if result != item.expectedTax {
			t.Errorf("For %v expected %v got %v", item.amount, item.expectedTax, result)
		}
	}
}

func BenchmarkCalculateTax(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateTax(500.00)
	}
}

func BenchmarkCalculateTax2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateTax2(500.00)
	}
}

func FuzzCalculateTax(f *testing.F) {
	seed := []float64{-1, -2, -2.5, 500.00, 1000.00, 1500.00, 0.0, 1501.00}

	for _, amount := range seed {
		f.Add(amount)
	}

	f.Fuzz(func(t *testing.T, amount float64) {
		t.Error("amount:", amount)
		result := CalculateTax(amount)
		if amount <= 0 && result != 0 {

			t.Errorf("Received %f but expected 0", result)
		}

		if amount > 20000 && result != 20 { // this will fail
			t.Errorf("Received %f but expected 20", result)
		}
	})
}
