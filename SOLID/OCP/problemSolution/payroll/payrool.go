package payroll

import "github.com/douglasandradeee/go-expert-course/SOLID/OCP/employee"

func Payroll(e employee.Employee) float64 {
	return e.CalculateSalary()
}
