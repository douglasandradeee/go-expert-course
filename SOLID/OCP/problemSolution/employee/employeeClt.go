package employee

type EmployeeCLT struct {
	MonthlySalary float64
}

func (e *EmployeeCLT) CalculateSalary() float64 {
	return e.MonthlySalary
}
