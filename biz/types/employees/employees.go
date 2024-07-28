package employeesType

// EmployeeAffiliation @routers /employees/affiliation [GET]
type EmployeeAffiliation struct {
	ID           int32  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	EmployeeID   int32  `gorm:"column:employee_id;not null" json:"employee_id"`
	FarmID       int32  `gorm:"column:farm_id;not null" json:"farm_id"`
	EnterpriseID int32  `gorm:"column:enterprise_id;not null" json:"enterprise_id"`
	CreateTime   string `gorm:"column:create_time;not null" json:"create_time"`
}

// CreateAffiliationJson @routers /employees/affiliation [POST]
type CreateAffiliationRequestJson struct {
	EmployeeID   int32 `gorm:"column:employee_id;not null" json:"employee_id"`
	FarmID       int32 `gorm:"column:farm_id;not null" json:"farm_id"`
	EnterpriseID int32 `gorm:"column:enterprise_id;not null" json:"enterprise_id"`
}
