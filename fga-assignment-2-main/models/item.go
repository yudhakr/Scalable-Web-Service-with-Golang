package models

type Item struct {
	ItemID      uint   `gorm:"primary_key"`
	ItemCode    string `gorm:"not null;unique;type:varchar(191)"`
	Description string `gorm:"not null;type:varchar(191)"`
	Quantity    int    `gorm:"not null;type:int"`
	OrderID     uint
}
