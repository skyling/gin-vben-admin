package models

// CodeIndex 各种code 生成
type CodeIndex struct {
	Base
	Type  string `gorm:"uniqueIndex:uni_type_date;size:32;not null;default:'';comment:类型"`
	Date  string `gorm:"uniqueIndex:uni_type_date;size:20;not null;default:'';comment:日期"`
	Index int64  `gorm:"type:int(11);not null;default:0;comment:序号"`
}
