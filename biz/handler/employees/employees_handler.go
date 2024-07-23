package employees

import (
	"ST/biz/dal"
	"ST/biz/types/employees"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// GetAffiliationInfo @routers /employees/affiliation [GET]
func GetAffiliationInfo(c *gin.Context) {
	employeeIds := c.QueryArray("employeeIds")
	if len(employeeIds) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "employeeIds parameter is required"})
		return
	}

	var intEmployeeIds []int32
	for _, id := range employeeIds {
		intID, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid employee ID: " + id})
			return
		}
		intEmployeeIds = append(intEmployeeIds, int32(intID))
	}

	var employee_affiliations []employeesType.EmployeeAffiliation
	err := dal.DB.Where("employee_id in ?", intEmployeeIds).Find(&employee_affiliations).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var responseData struct {
		Employees []struct {
			ID           int32 `json:"id"`
			EnterpriseID int32 `json:"enterprise_id"`
			FarmID       int32 `json:"farm_id"`
		} `json:"employees"`
	}

	for _, affiliation := range employee_affiliations {
		responseData.Employees = append(responseData.Employees, struct {
			ID           int32 `json:"id"`
			EnterpriseID int32 `json:"enterprise_id"`
			FarmID       int32 `json:"farm_id"`
		}{
			ID:           affiliation.EmployeeID,
			EnterpriseID: affiliation.EnterpriseID,
			FarmID:       affiliation.FarmID,
		})
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "success", "data": responseData})
}
