package call

import "leetcode/go/ctci/c7/call_center/employee"

type Call struct {
	MinKind  employee.Kind
	caller   any
	employee *employee.Employee
}

func (c *Call) AssingEmployee(employee *employee.Employee) error {
	employee.MarkUnavailable()
	c.employee = employee
	return nil
}

func (c *Call) InitiateCall() error {
	return nil
}
