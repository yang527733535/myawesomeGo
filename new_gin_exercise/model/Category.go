package model

type Category struct {
	ID         uint   `json:"id" gorm:"primary_key"`
	Name       string `json:"name" gorm:"type:varchar(50);not null;unique"`
	CreateTime MyTime `gorm:"comment:'创建时间';type:timestamp;";json:"createTime"`
	//UpTime     MyTime `gorm:"comment:'创建时间';type:timestamp;";json:"createTime"`
	//UpdateTime MyTime `gorm:"comment:'修改时间';type:timestamp;";json:"updateTime"`
}
