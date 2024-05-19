package payroll

import "github.com/douglasandradeee/go-expert-course/solid/ocp/employee"

func Payroll(e employee.Employee) float64 {
	return e.CalculateSalary()
}
