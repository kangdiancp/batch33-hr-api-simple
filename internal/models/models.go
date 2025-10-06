package models

type Region struct {
	RegionID   uint   `gorm:"primaryKey;column:region_id;type:int" json:"region_id"`
	RegionName string `gorm:"column:region_name;type:varchar(25)" json:"region_name"`
}

func (Region) TableName() string { return "hr.regions" }

type Country struct {
	CountryID   string `gorm:"primaryKey;column:country_id;type:char(2)" json:"country_id"`
	CountryName string `gorm:"column:country_name;type:varchar(40)" json:"country_name"`
	RegionID    uint   `gorm:"column:region_id" json:"region_id"`
	Region      Region `gorm:"foreignKey:RegionID;references:RegionID" json:"region"`
}

// hr.countries yg akan dicreate table di db
func (Country) TableName() string { return "hr.countries" }

type Department struct {
	DepartmentID   uint   `gorm:"" json:"department_id"`
	DepartmentName string `gorm:"" json:"department_name`
}
