package models

type Region struct {
	RegionID   uint   `gorm:"primaryKey;column:region_id;type:int" json:"region_id"`
	RegionName string `gorm:"column:region_name;type:varchar(25)" json:"region_name"`
}

func (Region) TableName() string { return "hr.regions" }
