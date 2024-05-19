package main

import (
	"fmt"

	"github.com/douglasandradeee/go-expert-course/solid/ocp/problemSolution/employee"
	"github.com/douglasandradeee/go-expert-course/solid/ocp/problemSolution/payroll"
)

func main() {
	pj := &employee.EmployeePJ{HoursWorked: 190, HourlyRate: 100}
	clt := &employee.EmployeeCLT{MonthlySalary: 9000}
	trainne := &employee.EmployeeTrainne{MonthlyStipend: 1500}

	fmt.Printf("Salary PJ: %.2f\n", payroll.Payroll(pj))
	fmt.Printf("Salary CLT: %.2f\n", payroll.Payroll(clt))
	fmt.Printf("Salary Trainne: %.2f\n", payroll.Payroll(trainne))
}
