package carriage

import (
	"fmt"
	"TruckMonitor-Backend/model"
)

func convertEmployeeName(employee *model.Employee)  (result string) {
	result = fmt.Sprintf("%s %s", employee.Surname, employee.Name)
	if len(employee.Patronymic) > 0 {
		result = fmt.Sprintf("%s %s", result, employee.Patronymic)
	}
	return
}
