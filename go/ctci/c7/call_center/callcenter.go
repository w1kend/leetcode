package callcenter

import (
	"fmt"
	"leetcode/go/ctci/c7/call_center/call"
	"leetcode/go/ctci/c7/call_center/employee"
	"leetcode/go/pkg/queue"
)

type CallCenter struct {
	employees map[int64]*employee.Employee
	available map[employee.Kind]queue.Queue

	callsQueue []call.Call

	priority []employee.Kind
}

func NewCallCenter() CallCenter {
	return CallCenter{
		employees:  map[int64]*employee.Employee{},
		available:  map[employee.Kind]queue.Queue{},
		callsQueue: make([]call.Call, 0, 20),
		priority:   []employee.Kind{employee.KindRespondent, employee.KindManager, employee.KindDirector},
	}
}

func (cc *CallCenter) DispatchCall(call call.Call) error {
	nextEmployee := cc.FindAvailableEmployee(call.MinKind)

	if nextEmployee == nil {
		cc.callsQueue = append(cc.callsQueue, call)
		return nil
	}

	err := call.AssingEmployee(nextEmployee)
	if err != nil {
		return fmt.Errorf("failed assining employee to a call: %w", err)
	}
	err = call.InitiateCall()
	if err != nil {
		return fmt.Errorf("failed initializing a call: %w", err)
	}

	return nil
}

func (cc *CallCenter) FindAvailableEmployee(minKind employee.Kind) *employee.Employee {
	kindIdx := 0
	for i, kind := range cc.priority {
		if kind == minKind {
			kindIdx = i
			break
		}
	}

	for _, employeeKind := range cc.priority[kindIdx:] {
		queue, ok := cc.available[employeeKind]
		if !ok || queue.IsEmpty() {
			continue
		}
		return cc.employees[int64(queue.Pop())]
	}

	return nil
}
