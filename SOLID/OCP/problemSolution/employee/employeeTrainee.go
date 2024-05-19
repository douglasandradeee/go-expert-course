package employee

type EmployeeTrainne struct {
	MonthlyStipend float64
}

func (e *EmployeeTrainne) CalculateSalary() float64 {
	return e.MonthlyStipend
}
